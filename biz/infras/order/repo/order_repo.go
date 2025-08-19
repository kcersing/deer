package repo

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"kcers-order/biz/dal/db/mysql/ent"
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/order/aggregate"
)

type OrderRepository interface {
	Save(order *aggregate.Order) error
	FindById(id int64) (*aggregate.Order, error)
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

func (o *OrderRepo) Save(order *aggregate.Order) error {
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
		//if err != nil {
		//	return errors.Wrap(err, "marshal event data failed")
		//}
		e.SetAggregateID(order.AggregateID)
		ets[i] = tx.OrderEvents.Create().
			SetEventID(e.GetId()).
			SetAggregateID(order.AggregateID).
			SetEventType(e.GetType()).
			SetAggregateType("order").
			SetEventVersion(order.Version).
			SetEventData(&common.EventData{
				Type:  e.GetType(),
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

func (o *OrderRepo) FindById(id int64) (*aggregate.Order, error) {
	//TODO implement me
	panic("implement me")
}
