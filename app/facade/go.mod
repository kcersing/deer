module facade

go 1.27.2

replace (
	common => ../../common
	gen => ../../gen
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
)

require (
	common v0.0.0-00010101000000-000000000000
	gen v0.0.0-00010101000000-000000000000
	github.com/bytedance/sonic v1.15.5
	github.com/cloudwego/hertz v0.10.6
	github.com/cloudwego/kitex v0.16.4
	github.com/golang-jwt/jwt/v4 v4.5.2
	github.com/hertz-contrib/jwt v1.0.4
	github.com/hertz-contrib/websocket v0.2.0
	github.com/medivhzhan/weapp/v2 v2.5.0
	github.com/pkg/errors v0.9.1
	github.com/rabbitmq/amqp091-go v1.16.0
	github.com/redis/go-redis/v9 v9.23.0
)

require (
	github.com/apache/thrift v0.20.0 // indirect
	github.com/bytedance/gopkg v0.1.4 // indirect
	github.com/bytedance/sonic/loader v0.5.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/cloudwego/gopkg v0.2.1 // indirect
	github.com/cloudwego/netpoll v0.7.5 // indirect
	github.com/elastic/pkcs8 v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/klauspost/cpuid/v2 v2.4.0 // indirect
	github.com/stretchr/objx v0.5.3 // indirect
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	go.uber.org/atomic v1.12.0 // indirect
	golang.org/x/arch v0.14.0 // indirect
	golang.org/x/crypto v0.57.0 // indirect
	golang.org/x/sys v0.48.0 // indirect
	google.golang.org/protobuf v1.36.12 // indirect
)
