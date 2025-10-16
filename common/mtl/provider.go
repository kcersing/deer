package mtl

import (
	"common/consts"
	"context"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
)

/**
* Provider
* https://cloudwego.cn/zh/docs/hertz/tutorials/third-party/open-telemetry/#provider
 */

func InitProvider(serviceName string) {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	defer func(ctx context.Context, p provider.OtelProvider) {
		_ = p.Shutdown(ctx)
	}(context.Background(), p)
}
