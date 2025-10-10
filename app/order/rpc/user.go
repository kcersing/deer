package rpc

import (
	"common/mw"
	"context"
	"gen/kitex_gen/base"
	"gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"order/conf"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{conf.GetConf().Etcd.Address})
	if err != nil {
		panic(err)
	}
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(conf.GetConf().Kitex.Service),
		provider.WithExportEndpoint(conf.GetConf().Etcd.ExportEndpoint),
		provider.WithInsecure(),
	)
	defer func(ctx context.Context, p provider.OtelProvider) {
		_ = p.Shutdown(ctx)
	}(context.Background(), p)
	c, err := userservice.NewClient(
		conf.GetConf().Etcd.UserName, // DestService
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientsMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.GetConf().Etcd.Address}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}
func GetUserInfo(ctx context.Context, id int64) error {
	_, err := userClient.GetUserInfo(ctx, &base.IdReq{Id: id})
	return err
}
