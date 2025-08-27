package infras

import (
	"context"
	db "deer/biz/dal/db/mysql"
	"deer/biz/infras/order/repo"
	"github.com/cloudwego/kitex/pkg/klog"
	"testing"
)

func TestOrder(t *testing.T) {
	//item := []common.Item{{Name: "商品1", ProductId: 1001, Quantity: 2, Price: 99.9}}
	//order := aggregate.NewOrder()
	//evt := events.NewCreatedOrderEvent("SN20230001", item, 99.9, 1, 2)
	dbs := db.InItDB("root:root@tcp(127.0.0.1:3306)/orders?charset=utf8mb4&parseTime=True&loc=Local", true)
	orderRepo := repo.NewOrderRepository(dbs, context.Background())
	//err := order.Apply(evt)
	//dispatcher := initEventHandlers()
	//err = dispatcher.Dispatch(context.Background(), evt)
	//klog.Info(err)
	//err = orderRepo.Save(order)
	//klog.Info(err)

	id, err := orderRepo.FindById(1)
	if err != nil {
		klog.Info(err)
	}
	klog.Info(id)

}
