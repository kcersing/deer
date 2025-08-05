package order

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/uuid"
	"sync"
	"time"
)

// Order 订单聚合根
type Order struct {
	Id             int64
	OrderSn        string
	Nature         string
	Items          []OrderItem
	TotalAmount    float64
	Status         OrderStatus
	CompletionAt   time.Time
	CloseAt        time.Time
	RefundAt       time.Time
	RefundedAmount float64

	Version      int64        // 乐观锁版本号
	Events       []Event      // 未提交事件
	mu           sync.RWMutex // 并发控制锁
	stateMachine *OrderStateMachine
}

// OrderItem 订单项值对象
type OrderItem struct {
	ProductId int64
	Quantity  int64
	Price     float64
}

// NewOrder 创建订单
func NewOrder(orderID int64, sn string, items []OrderItem, amount float64) *Order {
	order := &Order{
		Id:          orderID,
		OrderSn:     sn,
		Items:       items,
		TotalAmount: amount,
		Status:      OrderCreated,
	}
	order.stateMachine = NewOrderStateMachine(order)
	return order
}

// 领域行为：支付订单
func (o *Order) Pay(amount float64, method string) error {
	event := &OrderPaidEvent{
		BaseEvent: BaseEvent{
			EventID:     uuid.New().String(),
			AggregateID: o.Id,
			Timestamp:   time.Now(),
		},
		PayAmount: amount,
		PayMethod: method,
		PaidAt:    time.Now(),
	}

	return o.stateMachine.Transition(OrderPaid, event)
}

// 领域行为：取消订单
func (o *Order) Cancel(reason string) error {
	event := &OrderCancelledEvent{
		BaseEvent: BaseEvent{
			EventID:     uuid.New().String(),
			AggregateID: o.Id,
			Timestamp:   time.Now(),
		},
		Reason:      reason,
		CancelledAt: time.Now(),
	}

	return o.stateMachine.Transition(OrderCancelled, event)
}

// 事件管理方法
func (o *Order) AddEvent(event Event) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.Events = append(o.Events, event)
}

func (o *Order) GetUncommittedEvents() []Event {
	o.mu.RLock()
	defer o.mu.RUnlock()
	events := make([]Event, len(o.Events))
	copy(events, o.Events)
	return events
}

func (o *Order) ClearUncommittedEvents() {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.Events = []Event{}
}

// applyEvent 根据事件更新订单状态
func (o *Order) applyEvent(event Event) {
	o.mu.Lock() // 添加写锁保护状态变更
	defer o.mu.Unlock()

	o.Version++
	o.Events = append(o.Events, event)
	switch e := event.(type) {
	case *OrderCreatedEvent:
		o.Id = e.AggregateID  // 修正：使用事件中的AggregateID而非e.Id
		o.OrderSn = e.OrderSn // 补充：设置订单编号
		o.Items = e.Items
		o.TotalAmount = e.TotalAmount
		o.Status = OrderCreated
	case *OrderPaidEvent:
		o.Status = OrderPaid
	case *OrderCancelledEvent:
		o.Status = OrderCancelled
		o.CloseAt = e.CancelledAt
	case *OrderShippedEvent:
		o.Status = OrderShipped
	case *OrderCompletedEvent:
		o.Status = OrderCompleted
		o.CompletionAt = e.CompletionAt
	case *OrderRefundedEvent:
		o.Status = OrderRefunded
		o.RefundAt = e.RefundAt
		o.RefundedAmount = e.RefundedAmount
	default:
		// 添加未知事件日志
		klog.Errorf("unsupported event type: %T", e)
	}
}

//// 初始化
//client, _ := ent.Open("mysql", "dsn")
//dispatcher := NewEventDispatcher()
//inventoryHandler := &InventoryHandler{}
//dispatcher.RegisterHandler("OrderCreated", inventoryHandler)
//dispatcher.RegisterHandler("OrderCancelled", inventoryHandler)
//
//// 创建仓储
//repo := NewOrderRepository(client, 10) // 每10个事件创建一次快照
//
//// 创建订单
//items := []OrderItem{{ProductId: 1001, Quantity: 2, Price: 99.9}}
//order := NewOrder(1, "SN20230001", items, 199.8)
//
//// 发布创建事件
//event := NewOrderCreatedEvent(1, "SN20230001", items, 199.8, 10001)
//order.AddEvent(event)
//
//// 保存订单
//_ = repo.Save(context.Background(), order)
//
//// 分发事件
//_ = dispatcher.Dispatch(context.Background(), event)
//
//// 支付订单
//_ = order.Pay(199.8, "alipay")
//_ = repo.Save(context.Background(), order)
