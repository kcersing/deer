package logger

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

// InitLogger to init logrus
func InitLogger() {
	// Customizable output directory.
	logFilePath := consts.HlogFilePath
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
	logger := kitexlogrus.NewLogger()
	//logger.Logger().SetReportCaller(true)
	// hlog will warp a layer of logrus, so you need to calculate the depth of the caller file separately.
	//logger.Logger().AddHook(NewCustomHook(10))
	// Provides compression and deletion
	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
	}
	if runtime.GOOS == "linux" {
		logger.SetOutput(lumberjackLogger)
		logger.SetLevel(klog.LevelDebug)
	} else {
		//logger.SetOutput(lumberjackLogger)
		logger.SetLevel(klog.LevelDebug)
	}

	klog.SetLogger(logger)

}

// CustomHook Custom Hook for processing logs
type CustomHook struct {
	CallerDepth int
}

func NewCustomHook(depth int) *CustomHook {
	return &CustomHook{
		CallerDepth: depth,
	}
}

func (hook *CustomHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *CustomHook) Fire(entry *logrus.Entry) error {
	// Get caller information and specify depth
	pc, file, line, ok := runtime.Caller(hook.CallerDepth)
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		entry.Data["caller"] = fmt.Sprintf("%s:%d %s", file, line, funcName)
	}

	return nil
}
