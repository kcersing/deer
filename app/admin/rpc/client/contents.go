package client

import (
	"common/consts"
	"common/rpc/client"
	"sync"

	"gen/kitex_gen/contents/contentsservice"

	"github.com/cloudwego/kitex/pkg/klog"
)

var ContentsClient contentsservice.Client
var ContentsOnceClient sync.Once

func InitContentsRpc() {

	ContentsOnceClient.Do(func() {
		r := client.NewResolver(serviceResolver, consts.ContentsRpcServiceName, consts.AdminServiceName, consts.OpenTelemetryAddress)

		c, err := contentsservice.NewClient(
			r.ServiceName,
			r.Options()...,
		)
		r.NewOpenTelemetryProvider()
		if err != nil {
			klog.Fatalf("ERROR: cannot init client: %v\n", err)
		}
		ContentsClient = c

	})
}
