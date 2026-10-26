package mtl

import (
	"common/consts"
	"context"
	"log"
	"time"

	"github.com/hertz-contrib/obs-opentelemetry/provider"
)

/**
* Provider
* https://cloudwego.cn/zh/docs/hertz/tutorials/third-party/open-telemetry/#provider
 */
// main.go
//func main() {
//_, shutdown := mtl.InitProvider("my-service")
//defer shutdown()
// ... 启动 Hertz server
//    h.Spin() // 阻塞
//}

// InitProvider 的作用是初始化并返回一个开箱即用的 OpenTelemetry Provider，
// 让你的 Hertz 服务具备可观测性能力（链路追踪、指标采集、日志关联等）。
// 它本身不是一个独立的“服务”，而是一个在应用启动时运行的初始化函数，
// 属于 hertz-contrib/obs-opentelemetry 这个社区扩展库的一部分
// InitProvider是给 Hertz HTTP 框架用的
func InitProvider(serviceName string) (provider.OtelProvider, func()) {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		//配置导出地址
		provider.WithExportEndpoint(consts.OpenTelemetryAddress),
		//使用非安全连接
		provider.WithInsecure(),
	)
	shutdown := func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := p.Shutdown(ctx); err != nil {
			log.Println("opentelemetry: provider shutdown error:", err)
		}
	}
	return p, shutdown
}
