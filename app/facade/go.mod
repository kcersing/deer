module facade

go 1.25.6

replace (
	common => ../../common
	gen => ../../gen
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
)

require (
	common v0.0.0-00010101000000-000000000000
	gen v0.0.0-00010101000000-000000000000
	github.com/bytedance/sonic v1.15.0
	github.com/cloudwego/hertz v0.10.4
	github.com/cloudwego/kitex v0.16.0
	github.com/golang-jwt/jwt/v4 v4.5.2
	github.com/hertz-contrib/jwt v1.0.4
	github.com/hertz-contrib/websocket v0.2.0
	github.com/medivhzhan/weapp/v2 v2.5.0
	github.com/pkg/errors v0.9.1
	github.com/rabbitmq/amqp091-go v1.10.0
	github.com/redis/go-redis/v9 v9.17.3
)

require (
	github.com/apache/thrift v0.0.0-00010101000000-000000000000 // indirect
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic/loader v0.5.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/cloudwego/gopkg v0.1.8 // indirect
	github.com/cloudwego/netpoll v0.7.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/elastic/pkcs8 v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.14.0 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	google.golang.org/protobuf v1.36.8 // indirect
)
