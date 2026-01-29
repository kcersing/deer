module test

go 1.25.6

replace (
	common => ../../common
	gen => ../../gen
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
)

require (
	github.com/bytedance/sonic v1.15.0
	github.com/cloudwego/hertz v0.10.4
	github.com/cloudwego/kitex v0.16.0
	github.com/rabbitmq/amqp091-go v1.10.0
	github.com/streadway/amqp v1.1.0
)

require (
	github.com/bytedance/gopkg v0.1.3 // indirect
	github.com/bytedance/sonic/loader v0.5.0 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	golang.org/x/arch v0.14.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
)
