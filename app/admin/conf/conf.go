package conf

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/kr/pretty"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Env        string
	Hertz      Hertz      `yaml:"Hertz"`
	MySQL      MySQL      `yaml:"MySQL"`
	PostgreSQL PostgreSQL `yaml:"PostgreSQL"`
	Redis      Redis      `yaml:"Redis"`
}

type MySQL struct {
	DSN string `yaml:"DSN"`
}
type PostgreSQL struct {
	DSN string `yaml:"DSN"`
}
type Redis struct {
	Address  string `yaml:"Address"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	DB       int    `yaml:"DB"`
}

type Hertz struct {
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
		hlog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(conf); err != nil {
		hlog.Error("validate config error - %v", err)
		panic(err)
	}
	pretty.Printf("%+v\n", conf)
}
func LogLevel() hlog.Level {
	level := GetConf().Hertz.LogLevel
	switch level {
	case "trace":
		return hlog.LevelTrace
	case "debug":
		return hlog.LevelDebug
	case "info":
		return hlog.LevelInfo
	case "notice":
		return hlog.LevelNotice
	case "warn":
		return hlog.LevelWarn
	case "error":
		return hlog.LevelError
	case "fatal":
		return hlog.LevelFatal
	default:
		return hlog.LevelInfo
	}
}
