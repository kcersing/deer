package repo

import (
	"common/eventbus"
	"common/pkg/utils"
	"context"
	"order/biz/dal/db"
	"order/biz/dal/db/ent"
	entOrder "order/biz/dal/db/ent/order"
	orderevents2 "order/biz/dal/db/ent/orderevents"
	ordersnapshots2 "order/biz/dal/db/ent/ordersnapshots"
	"order/biz/infras"
	"order/biz/infras/aggregate"
	"order/biz/infras/common"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

const (
	// SnapshotFrequency 定义了创建快照的频率（每N个版本）
	SnapshotFrequency = 10
)

var (
	ErrConcurrency = errors.New("并发冲突：订单已被修改，请重试")
)

type OrderRepository interface {
	Save(ctx context.Context, order *aggregate.Order) (err error)
	FindById(ctx context.Context, id int64) (order *aggregate.Order, err error)
}

type OrderRepo struct {
	db        *ent.Client
	publisher *eventbus.EventPublisher
}

// NewOrderRepo 创建带可选依赖的仓储（订阅服务、事件发布器）
func NewOrderRepo() *OrderRepo {
	return &OrderRepo{
		db:        db.Client,
		publisher: infras.GetManager().Publisher,
	}
}

var _ OrderRepository = &OrderRepo{}

func (o *OrderRepo) Save(ctx context.Context, order *aggregate.Order) (err error) {
	es := order.GetUncommittedEvents()
	if len(es) == 0 {
		return nil
	}

	// 使用事务来保存聚合根
	err = infras.WithTx(func(tx *ent.Tx) error {
		return o.saveAggregateWithinTx(ctx, tx, order, es)
	})
	if err != nil {
		return err // 错误已在 withTx 中包装
	}
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
	// 对于非创建操作，需要保存订单项（如果它们可以被修改）
	// if order.GetVersion() > int64(len(es)) {
	//   if err := o.saveOrderItems(ctx, tx, order); err != nil {
	// 	    return err
	//   }
	// }

	if err := o.saveEvents(ctx, tx, order, es); err != nil {
		return err
	}

	// 优化：基于频率的快照策略
	if order.GetVersion()%SnapshotFrequency == 0 {
		if err := o.saveSnapshot(ctx, tx, order); err != nil {
			return err
		}
	}
	return nil
}

func (o *OrderRepo) saveOrderEntity(ctx context.Context, tx *ent.Tx, order *aggregate.Order) error {
	// 计算出此操作前的版本
	originalVersion := order.GetVersion() - int64(len(order.GetUncommittedEvents()))

	// 如果 originalVersion 是 0，说明是新订单
	if originalVersion == 0 {
		orderCreate := tx.Order.Create().
			SetSn(order.Sn).
			SetStatus(entOrder.Status(order.Status)).
			SetCreatedID(order.CreatedId).
			SetMemberID(order.MemberId).
			SetTotalAmount(order.TotalAmount).
			SetNature(order.Nature).
			SetVersion(order.GetVersion()) // 直接使用新版本

		orderEnt, err := orderCreate.Save(ctx)
		if err != nil {
			return errors.Wrap(err, "创建订单实体失败")
		}
		order.SetAggregateID(orderEnt.ID)

		// 创建订单时需要保存订单项
		return o.saveOrderItems(ctx, tx, order)
	}

	// 否则，是更新操作，使用乐观锁
	update := tx.Order.UpdateOneID(order.GetAggregateID()).
		SetStatus(entOrder.Status(order.Status)).
		SetVersion(order.GetVersion()).
		SetActual(order.Actual).
		SetRemission(order.Remission).
		Where(entOrder.Version(originalVersion)) // 乐观锁检查

	if order.CompletionAt != "" {
		completionAt, err := utils.GetStringDateTime(order.CompletionAt)
		if err == nil {
			update.SetCompletionAt(completionAt)
		}
	}
	if order.CloseAt != "" {
		closeAt, err := utils.GetStringDateTime(order.CloseAt)
		if err == nil {
			update.SetCloseAt(closeAt)
		}
	}
	if order.CloseNature != "" {
		update.SetCloseNature(order.CloseNature)
	}

	_, err := update.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return ErrConcurrency
		}
		return errors.Wrap(err, "更新订单实体失败")
	}

	return nil
}

func (o *OrderRepo) saveOrderItems(ctx context.Context, tx *ent.Tx, order *aggregate.Order) error {

	items := make([]*ent.OrderItemCreate, len(order.Items))
	for i, item := range order.Items {
		items[i] = tx.OrderItem.Create().
			SetName(item.Name).
			SetProductID(item.ProductId).
			SetQuantity(item.Quantity).
			SetUnitPrice(item.Price * 100).
			SetOrderID(order.GetAggregateID()).
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

		if e.GetAggregateID() == 0 {
			e.SetAggregateID(order.GetAggregateID())
		}
		eventData, err := sonic.Marshal(e)
		if err != nil {
			return errors.Wrap(err, "序列化事件数据失败")
		}
		ets[i] = tx.OrderEvents.Create().
			SetEventID(e.GetId()).
			SetAggregateID(e.GetAggregateID()).
			SetEventType(e.GetType()).
			SetAggregateType(e.GetAggregateType()).
			SetEventVersion(e.GetVersion()).
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
		SetAggregateVersion(order.GetVersion()).
		SetAggregateID(order.GetAggregateID()).
		SetAggregateData(orderData).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "保存订单快照失败")
	}
	return nil
}

// publishEvents 负责发布领域事件。
func (o *OrderRepo) publishEvents(ctx context.Context, events []common.Event) {
	if o.publisher != nil {
		for _, e := range events {
			if err := o.publisher.Local(ctx, e.GetType(), e); err != nil {
				klog.Errorf("发布事件失败(event_id=%s): %v", e.GetId(), err)
			}
		}
	}
}

func (o *OrderRepo) FindById(ctx context.Context, id int64) (*aggregate.Order, error) {
	order := aggregate.NewOrder()
	order.SetAggregateID(id) // 预设ID

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

	if snapshot != nil {
		if err := sonic.Unmarshal(snapshot.AggregateData, order); err != nil {
			return nil, errors.Wrap(err, "快照数据反序列化失败")
		}
		lastVersion = snapshot.AggregateVersion
	}

	// 从快照版本之后开始查询事件
	eventEnts, err := o.db.OrderEvents.Query().Where(
		orderevents2.AggregateID(id),
		orderevents2.EventVersionGT(lastVersion),
	).
		Order(ent.Asc(orderevents2.FieldEventVersion)).
		All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "查询事件记录失败")
	}

	if snapshot == nil && len(eventEnts) == 0 {
		return nil, common.ErrorNotFound
	}

	var eventStream []common.Event
	for _, eventEnt := range eventEnts {
		ev, err := NewEventWithData(eventEnt.EventType, eventEnt.EventData)
		if err != nil {
			klog.Warnf("跳过无法处理的事件: %v", err)
			continue
		}
		eventStream = append(eventStream, ev)
	}

	if err := order.Load(eventStream); err != nil {
		return nil, err
	}
	return order, nil
}
