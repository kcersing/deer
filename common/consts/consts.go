package consts

const (
	UserTableName    = "t_user"
	ProductTableName = "t_product"
	OrderTableName   = "t_order"
	MemberTableName  = "t_member"

	SecretKey   = "secret key"
	IdentityKey = "id"

	LoginName     = "admin"
	LoginPassword = "123456"

	MySQLDefaultDSN = "root:root@tcp(localhost:3306)/deer?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress     = "127.0.0.1:2379"
	ESAddress       = "http://localhost:9200"
	RedisAddress    = "127.0.0.1:6379"

	RedisConnPoolSize = 20

	RedisKey_User   = "user-"
	RedisKey_Member = "member-"

	ProductESIndex = "product"

	UserRpcServiceName    = "cwg.deer.user"
	MemberRpcServiceName  = "cwg.deer.member"
	OrderRpcServiceName   = "cwg.deer.order"
	ProductRpcServiceName = "cwg.deer.product"

	UserServiceAddress    = "127.0.0.1:9030"
	MemberServiceAddress  = "127.0.0.1:9031"
	OrderServiceAddress   = "127.0.0.1:9032"
	ProductServiceAddress = "127.0.0.1:9033"
	FacadeServiceAddress  = "127.0.0.1:8080"
)
