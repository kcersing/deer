package mw

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = ClientsMiddleware

// ClientsMiddleware 中间件
func ClientsMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		//  获取服务器信息打印
		klog.Infof("server address: %v, rpc timeout: %v, readwrite timeout: %v\n", ri.To().Address(), ri.Config().RPCTimeout(), ri.Config().ConnectTimeout())
		klog.Info(req)
		klog.Info(resp)

		if err = next(ctx, req, resp); err != nil {
			klog.Error(err)
			return err
		}
		return nil
	}
}
