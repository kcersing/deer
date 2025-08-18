package infras

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	db "kcers-order/biz/dal/db/mysql"
	"kcers-order/biz/infras/common"
	"kcers-order/biz/infras/order/aggregate"
	"kcers-order/biz/infras/order/events"
	"kcers-order/biz/infras/order/repo"
	"kcers-order/biz/infras/status"
	"testing"
)

func TestOrder(t *testing.T) {
	item := []common.Item{{ProductId: 1001, Quantity: 2, Price: 99.9}}
	CreateTest(&aggregate.Order{
		Sn:          "SN20230001",
		MemberId:    1,
		CreatedId:   2,
		Items:       item,
		TotalAmount: 99.9,
	})(t)
}
func initEventHandlers() *EventDispatcher {
	dispatcher := NewEventDispatcher()
	return dispatcher
}

func CreateTest(order *aggregate.Order) func(t *testing.T) {
	return func(t *testing.T) {

		klog.Info(order)

		evt := events.NewCreatedOrderEvent(order.Id, order.Items, order.TotalAmount, order.MemberId, order.CreatedId)

		evt.Data = common.EventData{
			Type:  string(status.Created),
			Event: evt,
		}

		dbs := db.InItDB("root:root@tcp(127.0.0.1:3306)/orders?charset=utf8mb4&parseTime=True&loc=Local", true)
		klog.Info(dbs)
		orderRepo := repo.NewOrderRepository(dbs, context.Background())

		err := order.AddEvent(evt)
		klog.Info(err)
		dispatcher := initEventHandlers()

		err = dispatcher.Dispatch(context.Background(), evt)
		err = orderRepo.Save(order)
		klog.Info(err)

		//item := []Item{{ProductId: 1001, Quantity: 2, Price: 99.9}}
		//ord := NewOrder("SN20230001", 111, item, 199.8)
		//
		//klog.Info(createEvent)
	}
}
