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
	MemberId       int64
	OrderSn        string
	Nature         string
	Items          []Item
	TotalAmount    float64
	Status         Status
	CompletionAt   time.Time
	CloseAt        time.Time
	RefundAt       time.Time
	RefundedAmount float64

	Version      int64        // 乐观锁版本号
	Events       []Event      // 未提交事件
	mu           sync.RWMutex // 并发控制锁
	stateMachine *StateMachine
}

// AddEvent 事件管理方法
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
	case *CreatedEvent:
		o.Id = e.AggregateID  // 修正：使用事件中的AggregateID而非e.Id
		o.OrderSn = e.OrderSn // 补充：设置订单编号
		o.Items = e.Items
		o.TotalAmount = e.TotalAmount
		o.Status = Created
	case *PaidEvent:
		o.Status = Paid
	case *CancelledEvent:
		o.Status = Cancelled
	case *ShippedEvent:
		o.Status = Shipped
	case *CompletedEvent:
		o.Status = Completed
	case *RefundedEvent:
		o.Status = Refunded
		o.RefundedAmount = e.RefundedAmount
	default:
		// 添加未知事件日志
		klog.Errorf("unsupported event type: %T", e)
	}
}

// NewOrder 创建订单
func NewOrder(sn string, memberId int64, items []Item, amount float64) *Order {
	order := &Order{
		OrderSn:     sn,
		MemberId:    memberId,
		Items:       items,
		TotalAmount: amount,
		Status:      Created,
	}
	order.stateMachine = NewStateMachine(order)
	return order
}

// Item 订单项值对象
type Item struct {
	ProductId int64
	Quantity  int64
	Price     float64
	Name      string
}

// 领域行为：支付订单
func (o *Order) Pay(amount float64, method string) error {
	event := &PaidEvent{
		BaseEvent: BaseEvent{
			EventID:     uuid.New().String(),
			AggregateID: o.Id,
			Timestamp:   time.Now(),
		},
		PayAmount: amount,
		PayMethod: method,
	}

	return o.stateMachine.Transition(Paid, event)
}

// 领域行为：取消订单
func (o *Order) Cancel(reason string) error {
	event := &CancelledEvent{
		BaseEvent: BaseEvent{
			EventID:     uuid.New().String(),
			AggregateID: o.Id,
			Timestamp:   time.Now(),
		},
		Reason: reason,
	}

	return o.stateMachine.Transition(Cancelled, event)
}

func (o *Order) Shipped() {

}

func (o *Order) Completed() {

}

func (o *Order) Refunded(amount float64) {

}
