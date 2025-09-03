package infras

import (
	"context"
	"deer/kitex_gen/deer/order"
	db "deer/rpc/order/biz/dal/mysql"
	"deer/rpc/order/biz/infras/aggregate"
	"deer/rpc/order/biz/infras/events"
	"deer/rpc/order/biz/infras/repo"
	"github.com/cloudwego/kitex/pkg/klog"
	"testing"
)

func TestOrderSave(t *testing.T) {
	item := []*order.Item{{Name: "商品1", ProductId: 1001, Quantity: 2, Price: 99.9}}
	order := aggregate.NewOrder()
	evt := events.NewCreatedOrderEvent("SN20230001", item, 99.9, 1, 2)
	dbs := db.InItDB("root:root@tcp(127.0.0.1:3306)/orders?charset=utf8mb4&parseTime=True&loc=Local", true)
	orderRepo := repo.NewOrderRepository(dbs, context.Background())
	err := order.Apply(evt)
	dispatcher := initEventHandlers()
	err = dispatcher.Dispatch(context.Background(), evt)
	klog.Info(err)
	err = orderRepo.Save(order)
	klog.Info(err)

}

func TestOrderPain(t *testing.T) {
	dbs := db.InItDB("root:root@tcp(127.0.0.1:3306)/orders?charset=utf8mb4&parseTime=True&loc=Local", true)
	orderRepo := repo.NewOrderRepository(dbs, context.Background())
	odr, err := orderRepo.FindById(1)
	if err != nil {
		klog.Info(err)
	}
	evt := events.NewPaidOrderEvent(odr.Id, 1)
	evt.PayedAmount = 99
	evt.PayMethod = "alipay"
	evt.Remission = 0
	evt.Reason = "测试支付"
	evt.PaySn = "SN20230001"
	evt.PrepayId = "SN20230001"
	evt.PayExtra = "{}"
	err = odr.Apply(evt)
	klog.Info(err)
	klog.Info(odr.GetAppliedEvents())

	evs := odr.GetUncommittedEvents()
	for _, ev := range evs {
		klog.Info(ev)
	}
	//err = orderRepo.Save(odr)
}

func TestOrderFindById(t *testing.T) {
	dbs := db.InItDB("root:root@tcp(127.0.0.1:3306)/orders?charset=utf8mb4&parseTime=True&loc=Local", true)
	orderRepo := repo.NewOrderRepository(dbs, context.Background())
	odr, err := orderRepo.FindById(1)
	if err != nil {
		klog.Info(err)
	}
	klog.Info(odr)
}
