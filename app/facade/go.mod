module facade

go 1.25.3

replace (
	common => ../../common
	gen => ../../gen
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
)

require (
	common v0.0.0-00010101000000-000000000000
	gen v0.0.0-00010101000000-000000000000
	github.com/cloudwego/hertz v0.10.2
	github.com/golang-jwt/jwt/v4 v4.5.2
	github.com/hertz-contrib/jwt v1.0.4
	github.com/pkg/errors v0.9.1
)

require (
	github.com/apache/thrift v0.16.0 // indirect
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic v1.14.1 // indirect
	github.com/bytedance/sonic/loader v0.3.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/cloudwego/gopkg v0.1.6 // indirect
	github.com/cloudwego/netpoll v0.7.2 // indirect
	github.com/elastic/pkcs8 v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/nyaruka/phonenumbers v1.0.55 // indirect
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.14.0 // indirect
	golang.org/x/crypto v0.41.0 // indirect
	golang.org/x/sys v0.35.0 // indirect
	google.golang.org/protobuf v1.36.8 // indirect
)
