package mtl

import (
	"common/consts"
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLog 配置 Kitex 全局日志。
// ⚠️ 必须在 InitTracing / InitProvider 之后调用，
// 否则日志无法自动关联 trace_id/span_id。
func InitLog(level klog.Level) {

	logFilePath := consts.LogFilePath
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	fileName := path.Join(logFilePath, time.Now().Format(time.DateOnly)+".log")

	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return
		}
	}

	ioWriter := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    consts.LogMaxSize,
		MaxBackups: consts.LogMaxBackups,
		MaxAge:     consts.LogMaxAge,
		Compress:   true,
	}

	buffered := &zapcore.BufferedWriteSyncer{
		WS:            zapcore.AddSync(ioWriter),
		FlushInterval: 5 * time.Second,
	}

	server.RegisterShutdownHook(func() {
		_ = buffered.Stop() // Stop 会先 Sync 再停 goroutine
	})

	//opts = append(opts, kitexzap.WithCoreEnc(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())))
	zapCfg := zap.NewProductionEncoderConfig()
	zapCfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	opts := []kitexzap.Option{
		kitexzap.WithCoreEnc(zapcore.NewJSONEncoder(zapCfg)),
		kitexzap.WithZapOptions(zap.AddCallerSkip(3), zap.AddCaller()),
	}

	logger := kitexzap.NewLogger(opts...)
	logger.SetLevel(level)

	logger.SetOutput(io.MultiWriter(buffered, os.Stdout))
	klog.SetLogger(logger)

}
