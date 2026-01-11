package logger

import (
	"common/consts"
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"

	kitexzap "github.com/hertz-contrib/obs-opentelemetry/logging/zap"

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
		output = zapcore.AddSync(ioWriter)
	} else {
		output = &zapcore.BufferedWriteSyncer{
			WS:            zapcore.AddSync(ioWriter),
			FlushInterval: time.Minute,
		}
	}

	_log := kitexzap.NewLogger(opts...)
	hlog.SetLogger(_log)
	hlog.SetLevel(hlog.LevelTrace)
	hlog.SetOutput(output)
	if runtime.GOOS == "linux" {
		hlog.SetOutput(output)
	}
	//klog.SetOutput(output)
}
