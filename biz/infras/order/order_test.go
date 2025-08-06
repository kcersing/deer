package order

import (
	"context"
	"kcers-order/biz/dal/db/mysql/ent"
	"testing"
)

func initEventHandlers() *EventDispatcher {
	dispatcher := NewEventDispatcher()
	inventoryHandler := &InventoryHandler{}
	dispatcher.RegisterHandler("created", inventoryHandler)
	dispatcher.RegisterHandler("cancelled", inventoryHandler)
	return dispatcher
}

func TestOrder(t *testing.T) {

	//// 初始化
	client, _ := ent.Open("mysql", "dsn")
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
	//order := NewOrder("SN20230001", 111, items, 199.8)
	//
	//// 发布创建事件
	//event := NewOrderCreatedEvent(1, "SN20230001", items, 199.8, 10001)
	//order.AddEvent(event)
	//
	//// 保存订单
	//_ = repo.Save(order)
	//
	//// 分发事件
	//_ = dispatcher.Dispatch(context.Background(), event)
	//
	//// 支付订单
	//_ = order.Pay(199.8, "alipay")
	//_ = repo.Save(order)
	//
	//_ = order.Cancel("cancel")
	//_ = repo.Save(order)

	// 创建订单
	items := []OrderItem{{ProductId: 1001, Quantity: 2, Price: 99.9}}
	order := NewOrder("SN20230001", 111, items, 199.8)

	repo := NewOrderRepository(client, 10)
	// 2. 调用仓库保存订单（已实现幂等性）
	repo.Save(order)

	// 3. 发布订单创建事件（触发库存扣减等后续操作）
	dispatcher := initEventHandlers()
	err := dispatcher.Dispatch(context.Background(), &CreatedEvent{
		BaseEvent:   BaseEvent{},
		OrderSn:     order.OrderSn,
		Items:       items,
		TotalAmount: 199.8,
		CreatedId:   100000,
	})
	if err != nil {
		return
	}

}
