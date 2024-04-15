# 定义变量
ifndef GOPATH
	GOPATH := $(shell go env GOPATH)
endif

GOBIN=$(GOPATH)/bin
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) mod tidy


all: deps build ## 默认的构建目标


clean: ## 清理目标
	$(GOCLEAN)
	rm -rf target
 

deps: ## 安装依赖目标
	@export GOPROXY=https://goproxy.cn,direct
	$(GOGET) -v
 

build: ## 构建目标
	$(GOBUILD) -o target/simple-go -v ./main.go
	cp ./config/app.ini target/config/ ## 复制配置文件



start: ## 运行目标
	cd ./target
	nohup ./simple-go  > /dev/null 2>&1 &



stop: ## 停止目标
	-pkill -f simple-go



restart: stop start ## 重启项目


.DEFAULT_GOAL := all ## 默认构建目标是



help: ## show help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
