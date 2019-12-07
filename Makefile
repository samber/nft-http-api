
SRC=main.go server.go routes.go
BIN=nft-http-api
VERSION=$(shell cat VERSION)
BUILD_ID := $(shell git rev-parse --short HEAD)

LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD_ID)"

all: build

build:
	CGO_ENABLED=0 GO111MODULE=on go build -v $(LDFLAGS) -o $(BIN) $(SRC)

release: clean deps
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 GO111MODULE=on go build $(LDFLAGS) -o $(BIN)_$(VERSION)_linux-amd64 $(SRC)
	CGO_ENABLED=0 GOOS=linux GOARCH=arm GO111MODULE=on go build $(LDFLAGS) -o $(BIN)_$(VERSION)_linux-arm $(SRC)
	CGO_ENABLED=0 GOOS=freebsd GOARCH=386 GO111MODULE=on go build $(LDFLAGS) -o $(BIN)_$(VERSION)_freebsd-386 $(SRC)

run-dev:
	GO111MODULE=on go run -v $(LDFLAGS) ${SRC} --listen 0.0.0.0:4242 --tls-cert fixtures/cert.pem --tls-key fixtures/key.pem

clean:
	rm -f $(BIN) ${BIN}_*

deps:
	GO111MODULE=on go mod download
