export GOPATH = $(PWD)/.go
export GOBIN = $(GOPATH)/bin

deps:
	which glide || go get -u github.com/Masterminds/glide
	mkdir -p $(GOPATH)/src && ln -sf $(PWD) .go/src/ #trick
	$(GOBIN)/glide install
