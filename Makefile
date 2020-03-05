SOURCE=./cmd/calendar
APP=calendar
GOBASE=$(shell pwd)
RELEASE_DIR=$(GOBASE)/bin

.DEFAULT_GOAL = build 

GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)
GO_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)	

build: 
	@go build -v -o ${APP} ${SOURCE}

lint:
	@goimports -w ${GO_SRC_DIRS}	
	@gofmt -s -w ${GO_SRC_DIRS}
	@golangci-lint run

run:
	@go run ${SOURCE} grpc_server --config=configs/config-dev.yml

run-client:
	@go run ${SOURCE} grpc_client --server=":50051"

run-web:
	@go run ${SOURCE} web_server --config=configs/config-dev.yml

test:
	@go test -v $(GO_TEST_DIRS)

gen:
	protoc -I api/proto --go_out=plugins=grpc:internal/grpcserver api/proto/calendar.proto

mod:
	go mod tidy

release:
	rm -rf ${RELEASE_DIR}${APP}*
	GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o ${RELEASE_DIR}/${APP}.exe main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${RELEASE_DIR}/${APP} main.go

.PHONY: build run release lint test gen mod run-web