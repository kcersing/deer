# 版本信息
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# 包列表
PACKAGES=$(shell go list ./... | grep -v /vendor/)
GO_FILES=$(shell find . -name "*.go" -type f -not -path "./vendor/*")


# 初始化构建环境
.PHONY: init
init:
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct
	export GO111MODULE=on
	export GOPROXY=https://goproxy.cn
	export GOPROXY=https://mirrors.aliyun.com/goproxy/
	go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
	go install github.com/cloudwego/hertz/cmd/hz@latest
	go install github.com/cloudwego/cwgo@latest
	go install github.com/cloudwego/thriftgo@latest
	go install github.com/cloudwego/thrift-gen-validator@latest

# start the environment of demo
.PHONY: start
start:
	docker-compose up -d

# stop the environment of demo
.PHONY: stop
stop:
	docker-compose down

#  cwgo server --type RPC --module order --server_name order --idl ../../../idl/rpc/order.thrift

VERSION?=1.0.0
BUILD_DATE=$(shell date +%FT%T%z)
# 模块
# MODULES=admin crm facade hardware member message order product system user
MODULES=admin member message order product system user
# 构建参数
LDFLAGS=-ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.buildTime=${BUILD_DATE}"
# 设置目标平台
# PLATFORMS=linux-amd64 windows-amd64 darwin-amd64
PLATFORMS=linux-amd64 windows-amd64
# 跨平台构建规则
.PHONY: build-run
build-run: $(PLATFORMS)
$(PLATFORMS):
	GOOS=$(shell echo $@ | cut -d- -f1) ; \
	GOARCH=$(shell echo $@ | cut -d- -f2) ; \

	@for module in $(MODULES); do \
        echo "app/$$module/"; \
		echo "go build ${LDFLAGS} ."; \
		cd "app/$$module/"; \
		go build ${LDFLAGS} . ; \
		cd ../..; \
    done

# 更新依赖库
up-get:
	@for module in $(MODULES); do \
		cd "app/$$module/"; \
		go get -u ./... ; \
		cd ../..; \
    done


# 显示当前版本信息
.PHONY: version
version:
	@echo "Version: ${VERSION}"
	@echo "Commit: ${COMMIT}" 
	@echo "Branch: ${BRANCH}" 