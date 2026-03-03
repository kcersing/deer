package repo

import (
	"context"

	"order/biz/dal/db"
	"order/biz/dal/db/ent"
	entOrder "order/biz/dal/db/ent/order"
	orderevents2 "order/biz/dal/db/ent/orderevents"
	ordersnapshots2 "order/biz/dal/db/ent/ordersnapshots"
	"order/biz/infras"
	"order/biz/infras/aggregate"
	"order/biz/infras/common"

	"common/eventbus"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type OrderRepository interface {
	Save(ctx context.Context, order *aggregate.Order) (err error)
	FindById(ctx context.Context, id int64) (order *aggregate.Order, err error)
	FindAll()
	Delete()
}

type OrderRepo struct {
	db        *ent.Client
	publisher *eventbus.EventPublisher
}

func (o *OrderRepo) FindAll() {
	panic(errors.New("not implemented"))
}

func (o *OrderRepo) Delete() {
	panic(errors.New("not implemented"))
}

// NewOrderRepo 创建带可选依赖的仓储（订阅服务、事件发布器）
func NewOrderRepo(pub *eventbus.EventPublisher) OrderRepository {
	return &OrderRepo{
		db:        db.Client,
		publisher: pub,
	}
}

var _ OrderRepository = &OrderRepo{}

func (o *OrderRepo) Save(ctx context.Context, order *aggregate.Order) (err error) {
	es := order.GetUncommittedEvents()
	klog.Info(len(es))
	if len(es) == 0 {
		return nil
	}

	// 使用事务来保存聚合根
	err = infras.WithTx(func(tx *ent.Tx) error {
		return o.saveAggregateWithinTx(ctx, tx, order, es)
	})
	klog.Info(len(es))
	if err != nil {
		return err // 错误已在 withTx 中包装
	}
	klog.Info(len(es))
	// 事务成功后，发布事件
	o.publishEvents(ctx, es)
	order.ClearUncommittedEvents() // 清除未提交事件
	return nil
}

// saveAggregateWithinTx 在一个事务中持久化订单聚合根的所有变更。
func (o *OrderRepo) saveAggregateWithinTx(ctx context.Context, tx *ent.Tx, order *aggregate.Order, es []common.Event) error {
	if err := o.saveOrderEntity(ctx, tx, order); err != nil {
		return err
	}
	if err := o.saveOrderItems(ctx, tx, order); err != nil {
		return err
	}
	if err := o.saveEvents(ctx, tx, order, es); err != nil {
		return err
	}
	if err := o.saveSnapshot(ctx, tx, order); err != nil {
		return err
	}
	return nil
}

func (o *OrderRepo) saveOrderEntity(ctx context.Context, tx *ent.Tx, order *aggregate.Order) error {
	orderEnt, err := tx.Order.Create().
		SetSn(order.Sn).
		SetStatus(entOrder.Status(order.Status)).
		SetCreatedID(order.CreatedId).
		SetMemberID(order.MemberId).
		SetVersion(order.Version).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "保存订单失败")
	}
	order.AggregateBase.SetAggregateID(orderEnt.ID)
	return nil
}

func (o *OrderRepo) saveOrderItems(ctx context.Context, tx *ent.Tx, order *aggregate.Order) error {
	items := make([]*ent.OrderItemCreate, len(order.Items))
	for i, item := range order.Items {
		items[i] = tx.OrderItem.Create().
			SetName(item.Name).
			SetProductID(item.ProductId).
			SetQuantity(item.Quantity).
			SetUnitPrice(item.Price).
			SetOrderID(order.AggregateID).
			SetCreatedID(order.CreatedId)
	}
	if _, err := tx.OrderItem.CreateBulk(items...).Save(ctx); err != nil {
		return errors.Wrap(err, "保存订单项失败")
	}
	return nil
}

func (o *OrderRepo) saveEvents(ctx context.Context, tx *ent.Tx, order *aggregate.Order, es []common.Event) error {
	ets := make([]*ent.OrderEventsCreate, len(es))
	for i, e := range es {
		e.SetAggregateID(order.AggregateID)
		eventData, err := sonic.Marshal(e)
		if err != nil {
			return errors.Wrap(err, "序列化事件数据失败")
		}
		ets[i] = tx.OrderEvents.Create().
			SetEventID(e.GetId()).
			SetAggregateID(order.AggregateID).
			SetEventType(e.GetType()).
			SetAggregateType(e.GetAggregateType()).
			SetEventVersion(order.Version).
			SetEventData(eventData)
	}
	if _, err := tx.OrderEvents.CreateBulk(ets...).Save(ctx); err != nil {
		return errors.Wrap(err, "保存订单事件失败")
	}
	return nil
}

func (o *OrderRepo) saveSnapshot(ctx context.Context, tx *ent.Tx, order *aggregate.Order) error {
	orderData, err := sonic.Marshal(order)
	if err != nil {
		return errors.Wrap(err, "序列化订单快照失败")
	}
	_, err = tx.OrderSnapshots.Create().
		SetAggregateVersion(order.Version).
		SetAggregateID(order.AggregateID).
		SetAggregateData(orderData).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "保存订单快照失败")
	}
	return nil
}

// publishEvents 负责发布领域事件。
func (o *OrderRepo) publishEvents(ctx context.Context, events []common.Event) {
	// 优先使用通用事件发布器进行分发
	if o.publisher != nil {
		for _, e := range events {
			if err := o.publisher.Distributed(ctx, e.GetType(), e); err != nil {
				klog.Errorf("发布事件失败(event_id=%s): %v", e.GetId(), err)
			}
		}
		return
	}
}

func (o *OrderRepo) FindById(ctx context.Context, id int64) (order *aggregate.Order, err error) {
	order = aggregate.NewOrder()
	// 1. 尝试加载最新快照
	snapshot, err := o.db.OrderSnapshots.
		Query().
		Where(ordersnapshots2.AggregateID(id)).
		Order(ent.Desc(ordersnapshots2.FieldAggregateVersion)).
		First(ctx)

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
		All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "查询事件记录失败")
	}

	var eventAll []common.Event
	for _, eventEnt := range eventAlls {
		ev, err := NewEventWithData(eventEnt.EventType, eventEnt.EventData)
		if err != nil {
			klog.Warnf("跳过无法处理的事件: %v", err)
			continue
		}
		eventAll = append(eventAll, ev)
	}
	err = order.Load(eventAll)
	if err != nil {
		return nil, err
	}
	return order, err
}
