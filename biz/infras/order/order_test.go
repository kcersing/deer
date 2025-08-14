package order

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	db "kcers-order/biz/dal/db/mysql"
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
	var err error
	DB := db.InItDB("root:root@tcp(127.0.0.1:3306)/orders?charset=utf8mb4&parseTime=True&loc=Local", true)

	ctx := context.Background()
	//// 创建订单
	//items := []Item{{ProductId: 1001, Quantity: 2, Price: 99.9}}
	//order := NewOrder("SN20230001", 111, items, 199.8)
	//
	//// 发布创建事件
	//event := NewOrderCreatedEvent(1, "SN20230001", items, 199.8, 10001)
	//order.AddEvent(event)
	//// 创建仓储
	repo := NewOrderRepository(DB, ctx, 1)
	//// 2. 调用仓库保存订单（已实现幂等性）
	//err = repo.Save(order)
	//if err != nil {
	//	klog.Error(err)
	//	return
	//}
	//// 3. 发布订单创建事件（触发库存扣减等后续操作）
	dispatcher := initEventHandlers()
	err = dispatcher.Dispatch(ctx, event)
	//if err != nil {
	//	return
	//}
	order, err := repo.FindById(1)
	klog.Info(err)
	klog.Info(order)
	////// 支付订单
	//_ = order.Pay(199.8, "alipay")
	//
	////assert.NoError(t, err)
	////assert.Equal(t, Paid, order.Status)
	////assert.Equal(t, int64(1), order.Version)
	//
	//
	//
	//_ = repo.Save(order)
	//
	//_ = order.Cancel("cancel")
	//
	//_ = repo.Save(order)

}
