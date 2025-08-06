package order

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"kcers-order/biz/dal/db/mysql/ent"
	order2 "kcers-order/biz/dal/db/mysql/ent/order"
	"kcers-order/biz/dal/db/mysql/ent/orderevents"
	"kcers-order/biz/dal/db/mysql/ent/ordersnapshots"
	"sync"
	"time"
)

type OrderRepository interface {
	Save(order *Order) error
	FindById(id int64) (*Order, error)
}

// OrderRepositoryImpl 订单仓储实现
type OrderRepositoryImpl struct {
	db              *ent.Client
	ctx             context.Context
	snapshotFreq    int                 // 快照频率
	subscriptionSvc SubscriptionService // 新增：订阅服务
}

// NewOrderRepository 创建订单仓储（更新构造函数）
func NewOrderRepository(client *ent.Client, ctx context.Context, snapshotFreq int) OrderRepository {
	return &OrderRepositoryImpl{
		db:              client,
		ctx:             ctx,
		snapshotFreq:    snapshotFreq,
		subscriptionSvc: NewSubscriptionService(client), // 初始化订阅服务
	}
}

// Save 保存订单并提交事件
func (o OrderRepositoryImpl) Save(order *Order) error {
	events := order.GetUncommittedEvents()
	if len(events) == 0 {
		return nil // 无事件无需保存
	}

	// 开启事务
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

	// 1. 保存订单主体
	orderEnt := tx.Order.Create().
		SetOrderSn(order.OrderSn).
		SetStatus(string(order.Status)).
		//SetTotalAmount(order.TotalAmount).
		SetVersion(order.Version).
		SetCreatedID(events[0].(*CreatedEvent).CreatedId). // 从创建事件获取创建人
		// 支持幂等性创建
		OnConflict().
		UpdateNewValues()

	if err = orderEnt.Exec(o.ctx); err != nil {
		return errors.Wrap(err, "failed to save order")
	}

	// 2. 批量保存订单项
	items := make([]*ent.OrderItemCreate, len(order.Items))
	for i, item := range order.Items {
		items[i] = tx.OrderItem.Create().
			SetOrderID(order.Id).
			SetProductID(item.ProductId).
			SetQuantity(item.Quantity).
			SetUnitPrice(item.Price)
	}
	if _, err = tx.OrderItem.CreateBulk(items...).Save(o.ctx); err != nil {
		return errors.Wrap(err, "failed to save order items")
	}

	// 3. 保存事件记录
	eventEntities := make([]*ent.OrderEventsCreate, len(events))
	for i, event := range events {
		eventData, err := json.Marshal(event)
		if err != nil {
			return errors.Wrapf(err, "failed to marshal event %s", event.GetID())
		}
		eventEntities[i] = tx.OrderEvents.Create().
			SetEventID(event.GetID()).
			SetAggregateID(event.GetAggregateID()).
			SetEventType(event.GetType()).
			SetEventData(string(eventData))
	}
	if _, err = tx.OrderEvents.CreateBulk(eventEntities...).Save(o.ctx); err != nil {
		return errors.Wrap(err, "failed to save events")
	}

	// 4. 按频率创建快照
	if o.snapshotFreq > 0 && len(order.Events)%o.snapshotFreq == 0 {
		snapshotData, err := json.Marshal(order)
		if err != nil {
			return errors.Wrap(err, "failed to marshal snapshot")
		}
		_, err = tx.OrderSnapshots.Create().
			SetAggregateID(order.Id).
			SetAggregateVersion(order.Version).
			SetAggregateData(string(snapshotData)).
			SetCreatedAt(time.Now()).
			Save(o.ctx)
		if err != nil {
			return errors.Wrap(err, "failed to save snapshot")
		}
	}

	// 提交事务并清理事件
	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	// 新增：提交事务后通知订阅者（确保事件已持久化）
	if err == nil { // 事务成功提交后处理订阅
		for _, event := range events {
			if err := o.subscriptionSvc.ProcessEvent(o.ctx, event); err != nil {
				// 记录订阅通知失败日志，但不回滚订单事务（事件已持久化）
				klog.Errorf("通知订阅者失败(event_id=%s): %v", event.GetID(), err)
			}
		}
	}

	order.ClearUncommittedEvents() // 事务成功后清除未提交事件
	return nil
}

// FindById 从仓储获取订单（事件溯源模式）
func (o OrderRepositoryImpl) FindById(id int64) (*Order, error) {
	// 1. 尝试加载最新快照
	snapshot, err := o.db.OrderSnapshots.
		Query().
		Where(ordersnapshots.AggregateID(id)).
		Order(ent.Desc(ordersnapshots.FieldAggregateVersion)).
		First(o.ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, errors.Wrap(err, "failed to query snapshot")
	}

	var order *Order
	var lastVersion int64

	if snapshot != nil {
		// 从快照恢复
		if err := json.Unmarshal([]byte(snapshot.AggregateData), &order); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal snapshot")
		}
		lastVersion = snapshot.AggregateVersion
	} else {
		// 无快照时从初始事件重建
		order = &Order{mu: sync.RWMutex{}}
		lastVersion = 0
	}

	// 2. 加载快照后发生的事件
	events, err := o.db.OrderEvents.
		Query().
		Where(
			orderevents.AggregateID(id),
			orderevents.EventVersionGT(lastVersion), // 假设事件表有version字段
		).
		Order(ent.Asc(orderevents.FieldCreatedAt)).
		All(o.ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query events")
	}

	// 3. 回放事件重建状态
	for _, eventEnt := range events {
		var event Event
		switch eventEnt.EventType {
		case "OrderCreated":
			event = &CreatedEvent{}
		case "OrderPaid":
			event = &PaidEvent{}
		case "OrderCancelled":
			event = &CancelledEvent{}
		// 其他事件类型...
		default:
			klog.Warnf("unsupported event type: %s", eventEnt.EventType)
			continue
		}
		if err := json.Unmarshal([]byte(eventEnt.EventData), event); err != nil {
			klog.Errorf("failed to unmarshal event %s: %v", eventEnt.EventID, err)
			continue
		}
		order.applyEvent(event)
	}

	return order, nil
}

// Update 更新订单（乐观锁实现）
func (o OrderRepositoryImpl) Update(order *Order) error {
	// 使用乐观锁条件：仅当数据库版本与当前版本一致时更新
	_, err := o.db.Order.
		UpdateOneID(order.Id).
		SetStatus(string(order.Status)).
		//SetTotalAmount(order.TotalAmount).
		SetVersion(order.Version).
		Where(order2.Version(order.Version - 1)). // 条件：当前版本=数据库版本
		Save(o.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return errors.New("订单未找到")
		}
		if ent.IsConstraintError(err) {
			return errors.New("订单被另一笔交易更新，请重试")
		}
		return errors.Wrap(err, "无法更新订单")
	}
	return nil
}
