module hardware

go 1.25.3

replace (
	common => ../../common
	gen => ../../gen
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
)

require (
	common v0.0.0-00010101000000-000000000000
	gen v0.0.0-00010101000000-000000000000
	github.com/bytedance/sonic v1.14.1
	github.com/cloudwego/hertz v0.9.4-0.20241021100040-3477b0309b81
	github.com/cloudwego/kitex v0.15.1
	github.com/hertz-contrib/websocket v0.2.0
	github.com/kr/pretty v0.3.1
	github.com/rabbitmq/amqp091-go v1.10.0
	github.com/redis/go-redis/v9 v9.14.0
	gopkg.in/validator.v2 v2.0.1
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/bytedance/go-tagexpr/v2 v2.9.2 // indirect
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic/loader v0.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/cloudwego/gopkg v0.1.6 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/henrylee2cn/ameda v1.4.10 // indirect
	github.com/henrylee2cn/goutil v0.0.0-20210127050712-89660552f6f8 // indirect
	github.com/klauspost/cpuid/v2 v2.2.11 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/nyaruka/phonenumbers v1.0.55 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.14.0 // indirect
	golang.org/x/sys v0.37.0 // indirect
	google.golang.org/protobuf v1.36.8 // indirect
)
