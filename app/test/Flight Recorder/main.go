package main

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"log"
	"os"
	"os/signal"
	"runtime/trace"
	"syscall"
	"time"
)

var fr *trace.FlightRecorder

func main() {
	cfg := trace.FlightRecorderConfig{
		MinAge:   5 * time.Second, // 至少保留 5 秒
		MaxBytes: 3 * 1024 * 1024, // 最大 3 MB
	}

	fr = trace.NewFlightRecorder(cfg)
	if err := fr.Start(); err != nil {
		log.Fatalf("failed to start FlightRecorder: %v", err)
	}
	defer fr.Stop()
	sigChan := make(chan os.Signal)        // 未缓冲通道
	signal.Notify(sigChan, syscall.SIGINT) // 可能导致信号丢失
	// 注册信号处理，收到SIGINT时导出trace
	go func() {
		<-sigChan
		f, err := os.Create("trace.out")
		if err != nil {
			log.Println("failed to create trace file:", err)
			return
		}
		defer f.Close()
		// 导出最近5秒的trace
		if _, err := fr.WriteTo(f); err != nil {
			log.Println("failed to write trace data:", err)
		}

	}()
	for {
		hlog.Info(1)
		time.Sleep(time.Minute * 1)
		panic("123")
	}
}

// 分析 go tool trace trace.out
