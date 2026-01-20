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
	Casbin     CasbinConf `yaml:"Casbin"`

	Minio  Minio  `yaml:"Minio"`
	Wechat Wechat `yaml:"Wechat"`
	Alipay AliPay `yaml:"Alipay"`
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
}
type CasbinConf struct {
	ModelText string ` json:"ModelText"`
}

type Minio struct {
	EndPoint        string `yaml:"EndPoint"`
	AccessKeyID     string `yaml:"AccessKeyID"`
	SecretAccessKey string `yaml:"SecretAccessKey"`
	UseSSL          bool   `yaml:"UseSSL"`

	VideoBucketName string `yaml:"VideoBucketName"`
	ImgBucketName   string `yaml:"ImgBucketName"`

	Url string `yaml:"Url"`
}
type Wechat struct {
	Appid              string `mapstructure:"appid" yaml:"appid"`
	AppSecret          string `mapstructure:"app_secret" yaml:"app_secret"`
	MchId              string `mapstructure:"mch_id" yaml:"mch_id"`
	ApiKey             string `mapstructure:"api_key" yaml:"api_key"`
	ApiV3Key           string `mapstructure:"api_v3_key" yaml:"api_v3_key"`
	CertFileContent    string `mapstructure:"cert_file_content" yaml:"cert_file_content"`
	KeyFileContent     string `mapstructure:"key_file_content" yaml:"key_file_content"`
	Pkcs12FileContent  string `mapstructure:"pkcs12_file_content" yaml:"pkcs12_file_content"`
	SerialNo           string `mapstructure:"serial_no" yaml:"serial_no"`
	NotifyUrl          string `mapstructure:"notify_url" yaml:"notify_url"`
	RefundNotifyUrl    string `mapstructure:"refund_notify_url" yaml:"refund_notify_url"`
	RSAPublicKeyPath   string `mapstructure:"rsa_public_key_path" yaml:"rsa_public_key_path"`
	WechatPaySerialNo  string `mapstructure:"wechat_pay_serial_no" yaml:"wechat_pay_serial_no"`
	CertificateKeyPath string `mapstructure:"certificate_key_path" yaml:"certificate_key_path"`
}

type AliPay struct {
	Appid                   string `mapstructure:"appid" yaml:"appid"`
	PrivateKey              string `mapstructure:"private_key" yaml:"private_key"`
	AppPublicCertContent    string `mapstructure:"app_public_cert_content" yaml:"app_public_cert_content"`
	AlipayRootCertContent   string `mapstructure:"alipay_root_cert_content" yaml:"alipay_root_cert_content"`
	AlipayPublicCertContent string `mapstructure:"alipay_public_cert_content" yaml:"alipay_public_cert_content"`
	NotifyUrl               string `mapstructure:"notify_url" yaml:"notify_url"`
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
