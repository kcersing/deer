package mtl

import (
	"common/consts"
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLog(env bool) {

	logFilePath := consts.LogFilePath
	if err := os.MkdirAll(logFilePath, 0o777); err != nil {
		panic(err)
	}

	// Set filename to date
	logFileName := time.Now().Format(time.DateOnly) + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return
		}
	}

	ioWriter := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    consts.LogMaxSize,    // A file can be up to 20M.
		MaxBackups: consts.LogMaxBackups, // Save up to 5 files at the same time.
		MaxAge:     consts.LogMaxAge,     // A file can exist for a maximum of 10 days.
		Compress:   true,                 // Compress with gzip.
	}

	var opts []kitexzap.Option
	var output zapcore.WriteSyncer
	if env != true {
		opts = append(opts, kitexzap.WithCoreEnc(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())))
		output = zapcore.AddSync(ioWriter)
	} else {
		opts = append(opts, kitexzap.WithCoreEnc(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())))
		// async log
		output = &zapcore.BufferedWriteSyncer{
			WS:            zapcore.AddSync(ioWriter),
			FlushInterval: time.Minute,
		}
	}
	server.RegisterShutdownHook(func() {
		output.Sync() //nolint:errcheck
	})
	_log := kitexzap.NewLogger(opts...)
	klog.SetLogger(_log)
	klog.SetLevel(klog.LevelTrace)
	klog.SetOutput(output)
	if runtime.GOOS == "linux" {
		klog.SetOutput(output)
	}
	//klog.SetOutput(output)
}
