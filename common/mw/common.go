package mw

import (
	"context"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = CommonMiddleware

// CommonMiddleware 中间件
func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// 打印请求
		klog.Infof("real request: %+v\n", req)
		// 打印远程服务信息
		klog.Infof("remote service name: %s, remote method: %s\n", ri.To().ServiceName(), ri.To().Method())
		klog.Info(req)
		klog.Info(resp)
		if err = next(ctx, req, resp); err != nil {
			klog.Error(err)

			return err
		}
		// 打印请求返回信息
		klog.Infof("real response: %+v\n", resp)
		return nil
	}
}
