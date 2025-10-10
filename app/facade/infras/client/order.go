package client

import (
	"common/mw"
	"common/pkg/conf"
	"common/pkg/errno"
	"context"
	"gen/kitex_gen/order"
	"gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var orderClient orderservice.Client

func initOrderRpc() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := orderservice.NewClient(
		conf.OrderRpcServiceName,
		client.WithResolver(r), // resolver
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry

		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientsMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "facade"}),
	)
	if err != nil {
		panic(err)
	}
	orderClient = c
}

func GetOrderById(ctx context.Context, orderId int64) (interface{}, error) {
	resp, err := orderClient.GetOrderInfo(ctx, &order.GetOrderInfoReq{
		Id: 0,
		Sn: nil,
	})
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.Code != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.Code, resp.BaseResp.Message)
	}
	return resp.Order, nil

}
