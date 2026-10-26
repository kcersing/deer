package mtl

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/trace"
	"sync"
	"time"

	"github.com/cloudwego/kitex/server"
)

var (
	fr         *trace.FlightRecorder
	exportOnce sync.Once
)

// InitFlightRecorder 初始化飞行记录器（FlightRecorder）。
// 飞行记录器会在内存中持续保留最近一段时间的运行时 trace 数据，
// 当收到 SIGINT（Ctrl+C）信号时，将这段"案发现场"快照导出到 trace.out 文件，
// 便于事后用 go tool trace 分析偶发性的性能问题。
func InitFlightRecorder() {
	// 配置飞行记录器：
	//   MinAge：至少保留最近 10 秒的 trace 数据（官方建议约为要排查异常时间窗口的 2 倍）
	//   MaxBytes：内存中最多保留 5 MB 的 trace 数据，超出后丢弃最旧的部分
	cfg := trace.FlightRecorderConfig{
		MinAge:   10 * time.Second, // 至少保留 5 秒
		MaxBytes: 5 * 1024 * 1024,  // 最大 5 MB
	}

	// 创建并启动飞行记录器
	fr = trace.NewFlightRecorder(cfg)
	if err := fr.Start(); err != nil {
		log.Fatalf("failed to start FlightRecorder: %v", err)
	}
	//defer fr.Stop()
	// 注意：这里不能写 defer fr.Stop()。
	// 因为 defer 会在本函数返回时立即执行，导致记录器刚启动就被停止，
	// 后面的 goroutine 就再也拿不到有效的 trace 数据了。
	// 记录器的生命周期应与程序一致，停止操作放到信号处理完成之后。

	export := func(reason string) {
		exportOnce.Do(func() {
			name := fmt.Sprintf("trace-%d.out ", time.Now().Unix())
			f, err := os.Create(name)
			if err != nil {
				log.Printf("[flight] create %s: %v", name, err)
				return
			}
			defer func() {
				if err := f.Close(); err != nil {
					log.Printf("[flight] close: %v", err)
				}
			}()
			// 将内存中保留的最近一段时间的 trace 数据写入文件
			if _, err := fr.WriteTo(f); err != nil {
				log.Printf("[flight] write: %v", err)
			} else {
				log.Printf("[flight] trace exported to %s (%s)", name, reason)
			}
			// 导出完成后停止记录器，释放其占用的资源
			fr.Stop()

		})

	}

	if sig := getSnapshotSignal(); sig != nil {
		// 创建带缓冲的信号通道，缓冲区大小为 1。
		// signal.Notify 向通道投递信号时采用非阻塞发送：
		// 如果通道无缓冲且接收方尚未就绪，信号会被直接丢弃。
		// 因此这里必须使用缓冲通道，避免 Ctrl+C 信号丢失。
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, sig)
		// 启动 goroutine 等待信号，收到后导出 trace 数据
		go func() {
			<-sigChan
			export("signal:" + sig.String())
		}()
	}

	server.RegisterShutdownHook(func() {
		export("shutdown")
	})

}
