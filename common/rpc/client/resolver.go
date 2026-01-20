package client

import (
	"github.com/cloudwego/kitex/pkg/discovery"
)

type Resolver struct {
	R                discovery.Resolver
	Resolver         string
	ServiceName      string
	BasicServiceName string
	EndpointAddress  string
}

func NewResolver(resolver, serviceName, basicServiceName, endpointAddress string) (r Resolver) {
	r = Resolver{
		Resolver:         resolver,
		ServiceName:      serviceName,
		BasicServiceName: basicServiceName,
		EndpointAddress:  endpointAddress,
	}

	r.R = GetResolver(r.Resolver, r.ServiceName)
	return r
}
