package client

import (
	"common/consts"
	"common/rpc/client"
	"gen/kitex_gen/member/memberservice"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
)

var MemberClient memberservice.Client
var memberOnceClient sync.Once

func InitMemberRpc() {

	memberOnceClient.Do(func() {
		nr := client.NewNacosResolver(consts.MemberRpcServiceName, consts.MemberRpcServiceName)

		r := client.Resolver{
			R:                nr,
			ServiceName:      consts.MemberRpcServiceName,
			BasicServiceName: consts.AdminServiceName,
			EndpointAddress:  consts.OpenTelemetryAddress,
		}

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
