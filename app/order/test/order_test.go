package test

import (
	"context"
	"gen/kitex_gen/base"
	"order/biz/dal/db"
	"order/biz/infras"

	"github.com/cloudwego/kitex/pkg/klog"

	"order/biz/infras/aggregate"
	"order/biz/infras/events"
	"order/biz/infras/repo"
	"testing"
)

func TestOrderSave(t *testing.T) {
	db.InitDB()
	repo.InitOrderRepository()

	item := []*base.OrderItem{{Name: "商品1", ProductId: 1001, Quantity: 2, Price: 9990}}
	order := aggregate.NewOrder()
	evt := events.NewCreatedOrderEvent("SN20230001", item, 9990, 1, 2)

	err := order.Apply(evt)
	dispatcher := infras.InitEventHandlers()
	err = dispatcher.Dispatch(context.Background(), evt)
	klog.Info(err)
	err = repo.OrderRepoClient.Save(order)
	klog.Info(err)

}

func TestOrderPain(t *testing.T) {

	odr, err := repo.OrderRepoClient.FindById(1)
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

	odr, err := repo.OrderRepoClient.FindById(1)
	if err != nil {
		klog.Info(err)
	}
	klog.Info(odr)
}
