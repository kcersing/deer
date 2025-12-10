MODULE = github.com/kcersing/deer

.PHONY: init
init:
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	export GO111MODULE=on
	export GOPROXY=https://goproxy.cn
	export GOPROXY=https://mirrors.aliyun.com/goproxy/
	go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
	go install github.com/cloudwego/hertz/cmd/hz@latest
	go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
	go install github.com/cloudwego/cwgo@latest
	go install github.com/cloudwego/thriftgo@latest

# start the environment of demo
.PHONY: start
start:
	docker-compose up -d

# stop the environment of demo
.PHONY: stop
stop:
	docker-compose down

# run the facade service
.PHONY: facade
facade:
	cd app/facade
	sh app/facade/run.sh

# run the user service
.PHONY: user
user:
	cd app/user
	go run app/user/*.go

# run the member service
.PHONY: member
member:
	cd app/member
	go run app/member/*.go

# run the order service
.PHONY: order
order:
	cd app/order
	go run app/order/*.go

# run the product service
.PHONY: product
product:
	cd app/product
	go run app/product/*.go


#  cwgo server --type RPC --module order --server_name order --idl ../../../idl/rpc/order.thrift