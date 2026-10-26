package mtl

import (
	"errors"
	"log"
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry

/**
* 初始化 Prometheus 指标收集
* Metric
* Metric 包括了追踪样本以及自动将指标与产生它们的追踪样本联系起来。
# https://cloudwego.cn/zh/docs/hertz/tutorials/third-party/open-telemetry/#metric
*/
//创建一个 Prometheus 指标注册表
//注册 Go 运行时和进程相关的默认指标
//把 metrics 服务注册到 Consul 服务发现中心
//启动一个 HTTP 服务，暴露 /metrics 端点供 Prometheus 抓取

// InitMetric 初始化 Prometheus 指标，并将 metrics 端口注册到 Etcd
// metricsAddr: 例如 ":9090"
// registryAddr: Etcd 的地址，例如 "127.0.0.1:2379"
func InitMetric(serviceName string, metricsAddr string, registryAddr string) {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	//r, _ := consul.NewConsulRegister(registryAddr)
	r, err := etcd.NewEtcdRegistry([]string{registryAddr})
	if err != nil {
		log.Fatalf("failed to create etcd registry: %v", err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", metricsAddr)
	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}

	// 注册到 Etcd
	if err = r.Register(registryInfo); err != nil {
		log.Printf("failed to register prometheus service to etcd: %v", err)
	} else {
		log.Printf("successfully registered prometheus metrics to etcd: %s", metricsAddr)
	}

	// 优雅退出时注销
	server.RegisterShutdownHook(func() {
		if err := r.Deregister(registryInfo); err != nil {
			log.Printf("failed to deregister prometheus service: %v", err)
		}
	})
	// 启动 HTTP 服务暴露 /metrics 端点
	// 启动 HTTP 服务暴露 /metrics 端点
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go func() {
		if err := http.ListenAndServe(metricsAddr, mux); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("metrics server error: %v", err)
		}
	}()

}
