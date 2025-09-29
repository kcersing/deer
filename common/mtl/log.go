package mtl

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap"
	"time"

	"go.uber.org/zap/zapcore"

	"io"
)

func InitLog(ioWriter io.Writer) {
	var opts []kitexzap.Option
	var output zapcore.WriteSyncer
	klog.Info(opts, output)

	//opts = append(opts, kitexzap.WithCoreEnc(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())))
	//output = zapcore.AddSync(ioWriter)

	opts = append(opts, kitexzap.WithCoreEnc(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())))
	output = &zapcore.BufferedWriteSyncer{
		WS:            zapcore.AddSync(ioWriter),
		FlushInterval: time.Minute,
	}
	server.RegisterShutdownHook(func() {
		output.Sync()
	})
	log := kitexzap.NewLogger(opts...)
	klog.SetLogger(log)
	klog.SetLevel(klog.LevelTrace)
	klog.SetOutput(output)

}
