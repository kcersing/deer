package client

import (
	"common/consts"
	"common/rpc/client"
	"gen/kitex_gen/message/messageservice"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

var MessageClient messageservice.Client
var messageOnceClient sync.Once

func InitMessageRpc() {

	messageOnceClient.Do(func() {
		r := client.NewResolver(serviceResolver, consts.MessageRpcServiceName, consts.AdminServiceName, consts.OpenTelemetryAddress)

		c, err := messageservice.NewClient(
			r.ServiceName,
			r.Options()...,
		)
		r.NewOpenTelemetryProvider()
		if err != nil {
			klog.Fatalf("ERROR: cannot init client: %v\n", err)
		}
		MessageClient = c
	})
}
