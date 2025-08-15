package infras

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"kcers-order/biz/infras/order/aggregate"
	"testing"
)

func TestOrder(t *testing.T) {
	item := []aggregate.Item{{ProductId: 1001, Quantity: 2, Price: 99.9}}
	CreateTest(&aggregate.Order{
		Sn:          "SN20230001",
		MemberId:    1,
		CreatedId:   2,
		Items:       item,
		TotalAmount: 99.9,
	})(t)
}
func CreateTest(order *aggregate.Order) func(t *testing.T) {
	return func(t *testing.T) {

		klog.Info(order.UncommittedEvents)
		//order.Confirm()
		//klog.Info(order)
		//item := []Item{{ProductId: 1001, Quantity: 2, Price: 99.9}}
		//ord := NewOrder("SN20230001", 111, item, 199.8)
		//
		//klog.Info(createEvent)
	}
}
