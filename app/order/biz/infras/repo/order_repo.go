package repo

import (
	"context"
	"deer/app/order/biz/dal/mysql/ent"
	entOrder "deer/app/order/biz/dal/mysql/ent/order"
	orderevents2 "deer/app/order/biz/dal/mysql/ent/orderevents"
	ordersnapshots2 "deer/app/order/biz/dal/mysql/ent/ordersnapshots"
	"deer/app/order/biz/infras/aggregate"
	"deer/app/order/biz/infras/common"
	"deer/app/order/biz/infras/events"
	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

type OrderRepository interface {
	Save(order *aggregate.Order) (err error)
	FindById(id int64) (order *aggregate.Order, err error)
	FindAll()
	Delete()
}

type OrderRepo struct {
	db  *ent.Client
	ctx context.Context
}

func (o *OrderRepo) FindAll() {
	//TODO implement me
	panic("implement me")
}

func (o *OrderRepo) Delete() {
	//TODO implement me
	panic("implement me")
}

func NewOrderRepository(db *ent.Client, ctx context.Context) OrderRepository {
	return &OrderRepo{
		db:  db,
		ctx: ctx,
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
		SetVersion(order.Version).
		OnConflict().
		UpdateNewValues()

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
	//if err == nil {
	//	for _, e := range es {
	//		if err := o.subscriptionSvc.ProcessEvent(o.ctx, e); err != nil {
	//			klog.Errorf("通知订阅者失败(event_id=%s): %v", e.GetID(), err)
	//		}
	//	}
	//}
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

	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.Wrap(err, "查询快照失败")
	}
	err = sonic.Unmarshal(snapshot.AggregateData, &order)
	if err != nil {
		return nil, err
	}

	//var lastVersion int64
	//if snapshot != nil {
	//	lastVersion = snapshot.AggregateVersion
	//}

	eventAlls, err := o.db.Debug().OrderEvents.Query().Where(
		orderevents2.AggregateID(id),
		//orderevents2.EventVersionGT(lastVersion),
	).
		Order(ent.Asc(orderevents2.FieldCreatedAt)).
		All(o.ctx)

	if err != nil {
		return nil, errors.Wrap(err, "查询事件记录失败")
	}

	var eventAll []common.Event
	var out common.Event
	for _, eventEnt := range eventAlls {
		switch eventEnt.EventType {
		case string(common.Created):
			out = &events.CreatedOrderEvent{}
		case string(common.Paid):
			out = &events.PaidOrderEvent{}
		case string(common.Shipped):
			out = &events.ShippedOrderEvent{}
		case string(common.Cancelled):
			out = &events.CancelledOrderEvent{}
		case string(common.Refunded):
			out = &events.RefundedOrderEvent{}
		case string(common.Completed):
			out = &events.CompletedOrderEvent{}
		default:
			klog.Warnf("unsupported event type: %s", eventEnt.EventType)
			continue
		}
		err := sonic.Unmarshal(eventEnt.EventData, &out)
		if err != nil {
			return nil, err
		}
		eventAll = append(eventAll, out)
	}
	err = order.Load(eventAll)
	if err != nil {
		return nil, err
	}
	return order, err
}
