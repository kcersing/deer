package infras

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	db "kcers-order/biz/dal/db/mysql"
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/order/aggregate"
	"kcers-order/biz/infras/order/events"
	"kcers-order/biz/infras/order/repo"
	"testing"
)

func TestOrder(t *testing.T) {
	item := []common.Item{{Name: "商品1", ProductId: 1001, Quantity: 2, Price: 99.9}}
	//CreateTest(&aggregate.Order{
	//	Sn:          "SN20230001",
	//	MemberId:    1,
	//	CreatedId:   2,
	//	Items:       item,
	//	TotalAmount: 99.9,
	//})(t)
	order := aggregate.NewOrder("SN20230001", item, 99.9, 1, 2)

	evt := events.NewCreatedOrderEvent(order.AggregateID, order.Items, order.TotalAmount, order.MemberId, order.CreatedId)

	dbs := db.InItDB("root:root@tcp(127.0.0.1:3306)/orders?charset=utf8mb4&parseTime=True&loc=Local", true)

	orderRepo := repo.NewOrderRepository(dbs, context.Background())

	err := order.AddUncommittedEvent(evt)

	dispatcher := initEventHandlers()
	err = dispatcher.Dispatch(context.Background(), evt)
	klog.Info(err)
	err = orderRepo.Save(order)
	klog.Info(err)

}

//func CreateTest(order *aggregate.Order) func(t *testing.T) {
//	return func(t *testing.T) {
//
//		klog.Info(order)
//
//
//
//		//item := []Item{{ProductId: 1001, Quantity: 2, Price: 99.9}}
//		//ord := NewOrder("SN20230001", 111, item, 199.8)
//		//
//		//klog.Info(createEvent)
//	}
//}
