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
	ESAddress            = "http://101.201.55.134:9200"
	DefaultLimit         = 10
	OpenTelemetryAddress = "127.0.0.1:4317"

	NacosIpAddr   = "101.201.55.134"
	NacosPort     = 8848
	NacosLogDir   = "./file/nacos/log"
	NacosCacheDir = "./file/nacos/cache"
	NacosLogLevel = "debug"

	AdminServiceName      = "admin"
	SystemRpcServiceName  = "system"
	UserRpcServiceName    = "user"
	MessageRpcServiceName = "message"
	ProductRpcServiceName = "product"
	MemberRpcServiceName  = "member"
	OrderRpcServiceName   = "order"
	CrmRpcServiceName     = "crm"

	AdminSnowflakeNode   = 1
	SystemSnowflakeNode  = 2
	UserSnowflakeNode    = 3
	MessageSnowflakeNode = 4
	ProductSnowflakeNode = 5
	MemberSnowflakeNode  = 6
	OrderSnowflakeNode   = 7
	CrmSnowflakeNode     = 8
)
