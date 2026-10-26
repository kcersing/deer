package conf

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env      string
	Kitex    Kitex    `yaml:"Kitex"`
	Database Database `yaml:"Database"`
	Redis    Redis    `yaml:"Redis"`
}
type Database struct {
	LogLevel string `yaml:"LogLevel"` // Add this line
	Type     string `yaml:"Type"`
	DSN      string `yaml:"DSN"`
}

type Redis struct {
	Address  string `yaml:"Address"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"DB"`
}

type Kitex struct {
	Service         string `yaml:"Service"`
	Address         string `yaml:"Address"`
	Node            int64  `yaml:"Node"`
	MetricsPort     string `yaml:"MetricsPort"`
	EnablePprof     bool   `yaml:"EnablePprof"`
	EnableGzip      bool   `yaml:"EnableGzip"`
	EnableAccessLog bool   `yaml:"EnableAccessLog"`

	LogLevel      string `yaml:"LogLevel"`
	LogFileName   string `yaml:"LogFileName"`
	LogMaxSize    int    `yaml:"LogMaxSize"`
	LogMaxBackups int    `yaml:"LogMaxBackups"`
	LogMaxAge     int    `yaml:"LogMaxAge"`

	Registry string `yaml:"Registry"`

	Resolver string `yaml:"Resolver"`
	Tracer   string `yaml:"Tracer"`
}

// GetConf gets configuration instance
func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	prefix := "conf"
	confFileRelPath := filepath.Join(prefix, "conf.yaml")
	content, err := os.ReadFile(confFileRelPath)
	if err != nil {
		panic(err)
	}
	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		klog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		klog.Error("validate config error - %v", err)
		panic(err)
	}

	//conf2 := new(Config)
	//v := viper.New()
	//v.AddConfigPath(confFileRelPath)
	//
	//if err := v.ReadInConfig(); err != nil {
	//	if err := v.Unmarshal(&conf2); err != nil {
	//		hlog.Fatalf("unmarshal err failed: %s", err.Error())
	//	}
	//}
	//v.WatchConfig()
	//v.OnConfigChange(func(e fsnotify.Event) {
	//	if err := v.Unmarshal(&conf2); err != nil {
	//		hlog.Fatalf("unmarshal err failed: %s", err.Error())
	//	}
	//
	//})
	//
	//pretty.Printf("%+v\n", conf2)
}

func LogLevel() klog.Level {
	level := GetConf().Kitex.LogLevel
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}
