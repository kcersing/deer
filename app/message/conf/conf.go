package conf

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
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
	Kitex      Kitex      `yaml:"Kitex"`
	MySQL      MySQL      `yaml:"MySQL"`
	PostgreSQL PostgreSQL `yaml:"PostgreSQL"`
	Redis      Redis      `yaml:"Redis"`
	Aliyun     Aliyun     `yaml:"Aliyun"`
	RabbitMq   RabbitMq   `yaml:"RabbitMq"`
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
}

type Aliyun struct {
	Access Access `mapstructure:"Access" yaml:"Access"`
	Sms    Sms    `mapstructure:"Sms" yaml:"Sms"`
}
type Access struct {
	AccessKeyId     string `mapstructure:"AccessKeyId" yaml:"AccessKeyId"`
	AccessKeySecret string `mapstructure:"AccessKeySecret" yaml:"AccessKeySecret"`
}
type Sms struct {
	Captcha SmsTemplate `mapstructure:"Captcha" yaml:"Captcha"`
}
type SmsTemplate struct {
	SignName     string `mapstructure:"SignName" yaml:"SignName"`
	TemplateCode string `mapstructure:"TemplateCode" yaml:"TemplateCode"`
}

type RabbitMq struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     int    `mapstructure:"port" yaml:"port"`
	Exchange string `mapstructure:"exchange" yaml:"exchange"`
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
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
	pretty.Printf("%+v\n", conf)
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
