package serversuite

import (
	"common/consts"
	"common/mw"
	"common/pkg/utils"
	"common/rpc/registry"
	"net"
	"strings"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

func Option(rr, serviceName, address string, flakeNode int64) (opts []server.Option) {
	// address
	localIp := utils.MustGetLocalIPv4()

	if strings.HasPrefix(address, ":") {
		address = localIp + address
	} else {
		address = localIp + ":" + address
	}

	addr, err := net.ResolveTCPAddr(consts.TCP, address)
	if err != nil {
		panic(err)
	}

	r := registry.Registry(rr)
	info := GetInfo(serviceName, address, flakeNode)

	if r != nil {
		opts = append(opts, server.WithRegistry(r))
	}

	opts = append(opts,
		server.WithServiceAddr(addr),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		//server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),

		server.WithRegistryInfo(info),
	)

	return opts
}
