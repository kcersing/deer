module message

go 1.25.3

replace (
	common => ../../common
	gen => ../../gen
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
)

require (
	common v0.0.0-00010101000000-000000000000
	entgo.io/ent v0.14.5
	gen v0.0.0-00010101000000-000000000000
	github.com/alibabacloud-go/darabonba-openapi/v2 v2.1.13
	github.com/alibabacloud-go/dysmsapi-20170525/v4 v4.1.3
	github.com/alibabacloud-go/tea v1.3.14
	github.com/alibabacloud-go/tea-utils/v2 v2.0.9
	github.com/bytedance/sonic v1.14.1
	github.com/cloudwego/hertz v0.9.0
	github.com/cloudwego/kitex v0.15.1
	github.com/go-sql-driver/mysql v1.9.3
	github.com/jackc/pgx/v5 v5.7.6
	github.com/kr/pretty v0.3.1
	github.com/lib/pq v1.10.9
	github.com/redis/go-redis/v9 v9.14.0
	gopkg.in/validator.v2 v2.0.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/alibabacloud-go/alibabacloud-gateway-spi v0.0.5 // indirect
	github.com/alibabacloud-go/debug v1.0.1 // indirect
	github.com/alibabacloud-go/endpoint-util v1.1.0 // indirect
	github.com/alibabacloud-go/openapi-util v0.1.1 // indirect
	github.com/aliyun/credentials-go v1.4.5 // indirect
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic/loader v0.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/clbanning/mxj/v2 v2.7.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/cloudwego/gopkg v0.1.6 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/tjfoc/gmsm v1.4.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.14.0 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
)
