package consts

const (
	LogFilePath    = "./file/log/"
	ExportFilePath = "./file/export"
	SecretKey      = "secret key"
	IdentityKey    = "id"
	LogLevel       = "debug"
	LogFileName    = "log/kitex.log"
	LogMaxSize     = 20
	LogMaxBackups  = 5
	LogMaxAge      = 10

	WechatFilePath = "./file/wechat/"

	TCP = "tcp"

	EtcdAddress          = "127.0.0.1:2379"
	ESAddress            = "http://127.0.0.1:9200"
	DefaultLimit         = 10
	OpenTelemetryAddress = "127.0.0.1:4317"

	NacosIpAddr   = "127.0.0.1"
	NacosPort     = 8848
	NacosLogDir   = "./file/nacos/log"
	NacosCacheDir = "./file/nacos/cache"
	NacosLogLevel = "debug"

	UserRpcServiceName    = "user"
	SystemRpcServiceName  = "system"
	ProductRpcServiceName = "product"
	OrderRpcServiceName   = "order"
	MemberRpcServiceName  = "member"
	CrmRpcServiceName     = "crm"
	AdminServiceName      = "admin"
)
