package repo

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"kcers-order/biz/dal/db/mysql/ent"
	orderevents2 "kcers-order/biz/dal/db/mysql/ent/orderevents"
	ordersnapshots2 "kcers-order/biz/dal/db/mysql/ent/ordersnapshots"
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/order/aggregate"
	"kcers-order/biz/infras/order/events"
	"kcers-order/biz/infras/status"
)

type OrderRepository interface {
	Save(order *aggregate.Order) (err error)
	FindById(id int64) (order *aggregate.Order, err error)
}

type OrderRepo struct {
	db  *ent.Client
	ctx context.Context
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
		SetOrderSn(order.Sn).
		SetStatus(string(order.Status)).
		SetCreatedID(order.CreatedId).
		SetMemberID(order.MemberId).
		SetVersion(order.Version).
		OnConflict().
		UpdateNewValues()

	order.AggregateID, err = orderEnt.ID(o.ctx)

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
		ets[i] = tx.OrderEvents.
			Create().
			SetEventID(e.GetId()).
			SetAggregateID(order.AggregateID).
			SetEventType(e.GetType()).
			SetAggregateType(e.GetAggregateType()).
			SetEventVersion(order.Version).
			SetEventData(&common.EventData{
				Event: e,
			})
	}
	if _, err = tx.OrderEvents.CreateBulk(ets...).Save(o.ctx); err != nil {
		return errors.Wrap(err, "保存订单事件失败")
	}

	_, err = tx.OrderSnapshots.Create().
		SetAggregateVersion(order.Version).
		SetAggregateID(order.AggregateID).
		SetAggregateData(order).
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
	klog.Info("123")
	// 1. 尝试加载最新快照
	snapshot, err := o.db.OrderSnapshots.
		Query().
		Where(ordersnapshots2.AggregateID(id)).
		Order(ent.Desc(ordersnapshots2.FieldAggregateVersion)).
		First(o.ctx)
	klog.Info(err)
	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.Wrap(err, "查询快照失败")
	}
	klog.Info(snapshot)
	var lastVersion int64
	if snapshot != nil {
		lastVersion = snapshot.AggregateVersion
	}

	eventAlls, err := o.db.Debug().OrderEvents.Query().Where(
		orderevents2.AggregateID(id),
		orderevents2.EventVersionGT(lastVersion),
	).
		Order(ent.Asc(orderevents2.FieldCreatedAt)).
		All(o.ctx)
	klog.Info(eventAlls)
	if err != nil {
		return nil, errors.Wrap(err, "查询事件记录失败")
	}

	for _, eventEnt := range eventAlls {

		klog.Info(eventEnt)
		var event common.Event
		klog.Info(eventEnt.EventType)
		switch eventEnt.EventType {
		case string(status.Created):

			eventData, ok := eventEnt.EventData.Event.(events.CreatedOrderEvent)
			if ok {
				klog.Info(eventData)
			}
			klog.Info(eventData, ok)
			//event = &events.CreatedOrderEvent{
			//	EventBase: common.EventBase{
			//		EventID:     eventEnt.EventID,
			//		AggregateID: eventEnt.AggregateID,
			//		Timestamp:   eventEnt.CreatedAt,
			//		EventType:   eventEnt.EventType,
			//	},
			//	Sn:          eventEnt.EventData.Sn,
			//	TotalAmount: eventEnt.EventData.TotalAmount,
			//	Items:       eventEnt.EventData.Items,
			//	MemberId:    eventEnt.EventData.MemberId,
			//	CreatedId:   eventEnt.EventData.CreatedId,
			//}
		case string(status.Paid):
			event = &events.PaidOrderEvent{}
		case string(status.Shipped):
			event = &events.ShippedOrderEvent{}
		case string(status.Cancelled):
			event = &events.CancelledOrderEvent{}
		case string(status.Refunded):
			event = &events.RefundedOrderEvent{}
		case string(status.Completed):
			event = &events.CompletedOrderEvent{}
		default:
			klog.Warnf("unsupported event type: %s", eventEnt.EventType)
			continue
		}
		err := order.Apply(event)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}
