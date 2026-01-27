package mw

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = ServerMiddleware

// ServerMiddleware 中间件
func ServerMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		//// 打印客户端信息
		klog.Infof("[mw server] client address: %v\n", ri.From().Address())
		klog.Info(req)
		klog.Info(resp)
		if err = next(ctx, req, resp); err != nil {
			klog.Error(err)
			return err
		}
		return nil
	}
}
