export GOPATH = $(PWD)/.go
export GOBIN = $(GOPATH)/bin
export PATH := $(GOBIN):${PATH}
export PROJECT = $(GOPATH)/src/lag-api

dep:
	@which dep || go get -u github.com/golang/dep/cmd/dep
	cd ${PROJECT} && dep ensure && dep status

build: dep
	cd $(PROJECT) && go build

run-static:
	cd $(PROJECT) && go run main.go -static-rules static-example.yml -lag 300ms

run-proxy:
	cd $(PROJECT) && go run main.go -proxy http://httpbin.org -lag 500ms

docker-build:
	docker build -t lag-api .
