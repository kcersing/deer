package repo

import (
	"context"
	"order/biz/dal/db"
	"order/biz/infras"
	"order/biz/infras/events"

	"order/biz/dal/db/ent"
	entOrder "order/biz/dal/db/ent/order"
	orderevents2 "order/biz/dal/db/ent/orderevents"
	ordersnapshots2 "order/biz/dal/db/ent/ordersnapshots"
	"order/biz/infras/aggregate"
	"order/biz/infras/common"

	"common/eventbus"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

var OrderRepoClient *OrderRepo

type OrderRepository interface {
	Save(order *aggregate.Order) (err error)
	FindById(id int64) (order *aggregate.Order, err error)
	FindAll()
	Delete()
}

type OrderRepo struct {
	db              *ent.Client
	ctx             context.Context
	subscriptionSvc infras.SubscriptionService
	publisher       *eventbus.EventPublisher
}

func (o *OrderRepo) FindAll() {
	//TODO implement me
	panic("implement me")
}

func (o *OrderRepo) Delete() {
	//TODO implement me
	panic("implement me")
}

func InitOrderRepository() {
	OrderRepoClient = &OrderRepo{
		db:  db.Client,
		ctx: context.Background(),
	}
}

// NewOrderRepositoryWithDeps 创建带可选依赖的仓储（订阅服务、事件发布器）
func NewOrderRepositoryWithDeps(db *ent.Client, ctx context.Context, sub infras.SubscriptionService, pub *eventbus.EventPublisher) OrderRepository {
	return &OrderRepo{
		db:              db,
		ctx:             ctx,
		subscriptionSvc: sub,
		publisher:       pub,
	}
}

var _ OrderRepository = &OrderRepo{}

func (o *OrderRepo) Save(order *aggregate.Order) (err error) {
	es := order.GetUncommittedEvents()
	if len(es) == 0 {
		return nil
	}
	tx, err := o.db.Tx(o.ctx)
	if err != nil {
		return errors.Wrap(err, "创建事务失败")
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			klog.Errorf("transaction panicked: %v", r)
			panic(r)
		} else if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				klog.Errorf("rollback failed: %v", rbErr)
			}
		}
	}()

	orderEnt := tx.Order.Create().
		SetSn(order.Sn).
		SetStatus(entOrder.Status(order.Status)).
		SetCreatedID(order.CreatedId).
		SetMemberID(order.MemberId).
		SetVersion(order.Version)
	//OnConflict().
	//UpdateNewValues()

	order.Id, err = orderEnt.ID(o.ctx)
	order.AggregateID = order.Id
	if err != nil {
		klog.Errorf("save order failed: %v", err)
		return errors.Wrap(err, "保存订单失败")
	}

	items := make([]*ent.OrderItemCreate, len(order.Items))
	for i, item := range order.Items {
		items[i] = tx.OrderItem.
			Create().
			SetName(item.Name).
			SetProductID(item.ProductId).
			SetQuantity(item.Quantity).
			SetUnitPrice(item.Price).
			SetOrderID(order.AggregateID).
			SetCreatedID(order.CreatedId)
	}
	if _, err = tx.OrderItem.CreateBulk(items...).Save(o.ctx); err != nil {
		return errors.Wrap(err, "保存订单项失败")
	}
	ets := make([]*ent.OrderEventsCreate, len(es))
	for i, e := range es {
		e.SetAggregateID(order.AggregateID)
		eventData, _ := sonic.Marshal(e)
		ets[i] = tx.OrderEvents.
			Create().
			SetEventID(e.GetId()).
			SetAggregateID(order.AggregateID).
			SetEventType(e.GetType()).
			SetAggregateType(e.GetAggregateType()).
			SetEventVersion(order.Version).
			SetEventData(eventData)
	}
	if _, err = tx.OrderEvents.CreateBulk(ets...).Save(o.ctx); err != nil {
		return errors.Wrap(err, "保存订单事件失败")
	}
	order.Id = order.AggregateID
	orderData, _ := sonic.Marshal(order)
	_, err = tx.OrderSnapshots.Create().
		SetAggregateVersion(order.Version).
		SetAggregateID(order.AggregateID).
		SetAggregateData(orderData).
		Save(o.ctx)
	if err != nil {
		return errors.Wrap(err, "保存订单快照失败")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "提交事务失败")
	}
	if err == nil {
		// 优先使用通用事件发布器进行分发（可发布到本地和 MQ），没有则回退到旧的 subscriptionSvc
		if o.publisher != nil {
			for _, e := range es {
				// 使用 Distributed 使事件同时发往 MQ 与本地（按需要调整）
				// 将领域事件作为 payload 传递，发布器会封装为 transport event
				// 设置时间与版本信息到 transport event 在 publisher 端处理
				if perr := o.publisher.Distributed(o.ctx, e.GetType(), e); perr != nil {
					klog.Errorf("发布事件失败(event_id=%s): %v", e.GetId(), perr)
				}
				// 轻微延迟以避免短时间内大量 goroutine 累积在 publishToMQ
			}
		} else if o.subscriptionSvc != nil {
			for _, e := range es {
				if err := o.subscriptionSvc.ProcessEvent(o.ctx, e); err != nil {
					klog.Errorf("通知订阅者失败(event_id=%s): %v", e.GetId(), err)
				}
			}
		}
	}
	order.ClearUncommittedEvents() // 事务成功后清除未提交事件
	return nil
}

func (o *OrderRepo) FindById(id int64) (order *aggregate.Order, err error) {
	order = aggregate.NewOrder()
	// 1. 尝试加载最新快照
	snapshot, err := o.db.OrderSnapshots.
		Query().
		Where(ordersnapshots2.AggregateID(id)).
		Order(ent.Desc(ordersnapshots2.FieldAggregateVersion)).
		First(o.ctx)

	var lastVersion int64 = 0
	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.Wrap(err, "查询快照失败")
	}

	// 如果找到了快照 (err == nil)
	if snapshot != nil {
		err = sonic.Unmarshal(snapshot.AggregateData, &order)
		if err != nil {
			return nil, errors.Wrap(err, "快照数据反序列化失败")
		}
		lastVersion = snapshot.AggregateVersion
	}

	// 从快照版本之后开始查询事件
	eventAlls, err := o.db.OrderEvents.Query().Where(
		orderevents2.AggregateID(id),
		orderevents2.EventVersionGT(lastVersion),
	).
		// 使用 EventVersion 排序以确保事件顺序的绝对正确性
		Order(ent.Asc(orderevents2.FieldEventVersion)).
		All(o.ctx)

	if err != nil {
		return nil, errors.Wrap(err, "查询事件记录失败")
	}

	var eventAll []common.Event
	for _, eventEnt := range eventAlls {
		switch eventEnt.EventType {
		case string(common.Created):
			ev := &events.CreatedOrderEvent{}
			if err := sonic.Unmarshal(eventEnt.EventData, ev); err != nil {
				return nil, err
			}
			eventAll = append(eventAll, ev)
		case string(common.Paid):
			ev := &events.PaidOrderEvent{}
			if err := sonic.Unmarshal(eventEnt.EventData, ev); err != nil {
				return nil, err
			}
			eventAll = append(eventAll, ev)
		case string(common.Shipped):
			ev := &events.ShippedOrderEvent{}
			if err := sonic.Unmarshal(eventEnt.EventData, ev); err != nil {
				return nil, err
			}
			eventAll = append(eventAll, ev)
		case string(common.Cancelled):
			ev := &events.CancelledOrderEvent{}
			if err := sonic.Unmarshal(eventEnt.EventData, ev); err != nil {
				return nil, err
			}
			eventAll = append(eventAll, ev)
		case string(common.Refunded):
			ev := &events.RefundedOrderEvent{}
			if err := sonic.Unmarshal(eventEnt.EventData, ev); err != nil {
				return nil, err
			}
			eventAll = append(eventAll, ev)
		case string(common.Completed):
			ev := &events.CompletedOrderEvent{}
			if err := sonic.Unmarshal(eventEnt.EventData, ev); err != nil {
				return nil, err
			}
			eventAll = append(eventAll, ev)
		default:
			klog.Warnf("unsupported event type: %s", eventEnt.EventType)
			continue
		}
	}
	err = order.Load(eventAll)
	if err != nil {
		return nil, err
	}
	return order, err
}
