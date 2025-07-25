package order

import (
	"fmt"
	"sync"
)

type Order struct {
	Id          int64
	Items       []OrderItem
	TotalAmount float64
	Status      OrderStatus
	Version     int          // 乐观锁版本号
	Events      []Event      // 事件列表
	mu          sync.RWMutex // 读写锁
}

// OrderItem 订单商品项
type OrderItem struct {
	ProductId int64
	Quantity  int
	Price     float64
}

func CreateOrder(Id int64, items []OrderItem) (order *Order, err error) {
	return
}

func (o *Order) applyEvent(event Event) {

	o.Version++
	o.Events = append(o.Events, event)
	switch e := event.(type) {
	case *OrderCreatedEvent:
		o.Id = e.Id
		o.Items = e.Items
		o.TotalAmount = e.TotalAmount
		o.Status = OrderCreated
	case *OrderPaidEvent:
		o.Status = OrderPaid
	case *OrderCancelledEvent:
		o.Status = OrderCancelled
	case *OrderShippedEvent:
		o.Status = OrderShipped
	case *OrderCompletedEvent:
		o.Status = OrderCompleted
	case *OrderRefundedEvent:
		o.Status = OrderRefunded
		o.TotalAmount -= e.RefundedAmount
	default:

	}
}

// GetUncommittedEvents 获取未提交的事件
func (o *Order) GetUncommittedEvents() []Event {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return append([]Event{}, o.Events...)
}

// ClearUncommittedEvents 清除未提交的事件
func (o *Order) ClearUncommittedEvents() {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.Events = make([]Event, 0)
}
func (o *Order) GetId() int64 {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.Id
}
func (o *Order) GetStatus() OrderStatus {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.Status
}

func (o *Order) GetVersion() int {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.Version
}
func (o *Order) GetEvents() []Event {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.Events
}
func (o *Order) GetTotalAmount() float64 {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.TotalAmount
}

func (o *Order) Pay(payMethod string) (err error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.Status != OrderCreated {
		return fmt.Errorf("订单状态错误，当前状态：%s ", o.Id)
	}

	if !o.Status.CanTransitionTo(OrderCancelled) {
		return fmt.Errorf("订单状态不允许从 %s 转换到 %s ", o.Status, OrderCancelled)
	}
	o.Status = OrderPaid

	return
}
