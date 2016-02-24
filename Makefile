all : clean deps test build
.PHONY: all

export GOARCH ?= amd64
export CGO_ENABLED ?= 1
export CI_BUILD_NUMBER ?= 0

LDFLAGS += -X "main.buildDate=$(shell date -u '+%Y-%m-%d %H:%M:%S %Z')"
LDFLAGS += -X "main.build=$(CI_BUILD_NUMBER)"

savedeps:
	rm -rf vendor Godeps
	godep save -t

watch:
	go get github.com/onsi/ginkgo/ginkgo
	ginkgo watch -r -cover

clean:
	go clean -v -i ./...

deps:
	go get -t -v ./...

test:
	go test `go list ./... | grep -v /vendor/` -cover -ginkgo.failFast

build:
	go build -o db_broker -ldflags '-s -w $(LDFLAGS)'
