package main

import (
	"admin/biz/infras/logger"
	"admin/biz/mw"
	"admin/conf"
	"admin/rpc"
	"common/consts"
	"common/mtl"
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/monitor-prometheus"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/opensergo/sentinel/adapter"
	"github.com/hertz-contrib/pprof"
	"github.com/spf13/cobra"
)

var (
	version   string // Git提交哈希值
	buildTime string // 构建时间
	commit    string
	goVersion = fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
)

var rootCmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin HTTP service for deer platform",
	Run: func(cmd *cobra.Command, args []string) {

		// 1. 初始化 OpenTelemetry（Tracing、Metrics、Logs）
		mtl.InitProvider(consts.AdminServiceName)
		// 2. 初始化日志（必须在 Provider 之后，以关联 TraceID）
		logger.InitLog(true)
		// 3. 初始化 Prometheus 指标并注册到 Etcd
		// ⚠️ 建议：从 conf 读取端口，避免单机多服务冲突
		metricsPort := fmt.Sprintf(":%d", 9090)
		mtl.InitMetric(consts.AdminServiceName+"-metrics", metricsPort, consts.EtcdAddress)
		// 4. 初始化飞行记录器
		mtl.InitFlightRecorder()

		rpc.InitRpc()
		mw.InitJwt()

		// 2. 读取 Hertz 配置
		serviceAddress := conf.GetConf().Hertz.Address
		serviceTracer := conf.GetConf().Hertz.Tracer

		// 3. 初始化 Tracer
		tracer, cfg := tracing.NewServerTracer()

		// 4. 构建 Hertz 实例
		h := server.New(
			server.WithStreamBody(true),
			server.WithHostPorts(serviceAddress),
			server.WithHandleMethodNotAllowed(true),
			//server.WithRegistry(r, info),
			tracer,
			server.WithTracer(
				prometheus.NewServerTracer(serviceTracer, "/hertz",
					prometheus.WithEnableGoCollector(true),
				),
			),
		)
		//h.Use(gzip.Gzip(gzip.DefaultCompression))

		// 静态文件导出
		h.Static("/export", "./files/")

		// 注册 pprof 性能分析
		pprof.Register(h)

		// OTel 中间件
		h.Use(tracing.ServerMiddleware(cfg))

		// Sentinel 流控
		h.Use(adapter.SentinelServerMiddleware(
			adapter.WithServerBlockFallback(func(c context.Context, ctx *app.RequestContext) {
				ctx.JSON(http.StatusTooManyRequests, nil)
				ctx.Abort()
			}),
		))

		register(h)
		// 5. 启动 Hertz 服务
		// 使用 h.Run() 替代 h.Spin()，底层会监听 SIGINT/SIGTERM 并优雅退出
		if err := h.Run(); err != nil {
			hlog.Fatalf("Hertz server stopped with error: %v", err)
		}
	},
}

func init() {
	if version == "" {
		version = "unknown"
	}
	rootCmd.Version = version
	rootCmd.SetVersionTemplate(
		fmt.Sprintf("Git Commit Hash: %s\nUTC Build Time : %s\nVersion : %s\nGolang Version : %s\n",
			commit, buildTime, version, goVersion),
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		hlog.Error(err)
		os.Exit(1)
	}
}
