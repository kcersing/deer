package consts

const (
	LogFilePath    = "./file/log/"
	ExportFilePath = "./file/export"
	SecretKey      = "secret key"
	IdentityKey    = "id"

	WechatFilePath = "./file/wechat/"

	TCP                  = "tcp"
	ExportEndpoint       = "4317"
	EtcdAddress          = "101.201.55.134:2379"
	ESAddress            = "http://101.201.55.134:9200"
	DefaultLimit         = 10
	OpenTelemetryAddress = "101.201.55.134:4317"

	NacosIpAddr   = "101.201.55.134"
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
