package client

import (
	"common/consts"
	"common/rpc/client"
	"gen/kitex_gen/member/memberservice"
	"message/conf"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

var MemberClient memberservice.Client
var memberOnceClient sync.Once
var serviceResolver = conf.GetConf().Kitex.Resolver

func InitMemberRpc() {

	memberOnceClient.Do(func() {
		r := client.NewResolver(serviceResolver, consts.MemberRpcServiceName, consts.MessageRpcServiceName, consts.OpenTelemetryAddress)

		c, err := memberservice.NewClient(
			r.ServiceName,
			r.Options()...,
		)
		r.NewOpenTelemetryProvider()
		if err != nil {
			klog.Fatalf("ERROR: cannot init client: %v\n", err)
		}
		MemberClient = c

	})
}
