package consts

const (
	UserTableName    = "t_user"
	ProductTableName = "t_product"
	OrderTableName   = "t_order"
	MemberTableName  = "t_member"
	LogFilePath      = "./file/log/"
	ExportFilePath   = "./file/export"
	SecretKey        = "secret key"
	IdentityKey      = "id"

	TCP            = "tcp"
	ExportEndpoint = "4317"
	EtcdAddress    = "101.201.55.134:2379"
	ESAddress      = "http://101.201.55.134:9200"
	DefaultLimit   = 10

	UserRpcServiceName    = "user"
	SystemRpcServiceName  = "system"
	ProductRpcServiceName = "product"
	OrderRpcServiceName   = "order"
	MemberRpcServiceName  = "member"
	CrmRpcServiceName     = "crm"
	AdminServiceName      = "admin"
)
