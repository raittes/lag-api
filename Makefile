export GOPATH = $(PWD)/.go
export GOBIN = $(GOPATH)/bin
export PROJECT = $(GOPATH)/src/lag-api

deps:
	which glide || go get -u github.com/Masterminds/glide
	mkdir -p $(GOPATH)/src && ln -sf $(PWD) .go/src/ #trick
	$(GOBIN)/glide install

run-static:
	cd $(PROJECT) && go run main.go -static-rules static-example.yml -lag 300ms

run-proxy:
	cd $(PROJECT) && go run main.go -proxy http://httpbin.org -lag 500ms
