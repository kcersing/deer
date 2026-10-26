package logger

import (
	"common/consts"
	"io"
	"log"
	"os"
	"path"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.uber.org/zap"

	hertzzap "github.com/hertz-contrib/logger/zap"
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
		MaxSize:    consts.LogMaxSize,
		MaxBackups: consts.LogMaxBackups,
		MaxAge:     consts.LogMaxAge,
		Compress:   true,
	}

	var opts []hertzzap.Option
	var output zapcore.WriteSyncer
	zapCfg := zap.NewProductionEncoderConfig()
	zapCfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	opts = append(opts, hertzzap.WithCoreEnc(zapcore.NewJSONEncoder(zapCfg)))
	opts = append(opts, hertzzap.WithZapOptions(zap.AddCaller(), zap.AddCallerSkip(3)))

	output = &zapcore.BufferedWriteSyncer{
		WS:            zapcore.AddSync(ioWriter),
		FlushInterval: time.Minute,
	}

	logger := hertzzap.NewLogger(opts...)

	logger.SetLevel(hlog.LevelTrace)

	fileWriter := io.MultiWriter(output, os.Stdout)
	logger.SetOutput(fileWriter)
	hlog.SetLogger(logger)
}
