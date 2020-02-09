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

lint:
	@goimports -w ${GO_SRC_DIRS}
	@golangci-lint run
	@#gofmt -w ${GO_SRC_DIRS}

build: 
	@go build -v -o ${APP} ${SOURCE}

run:
	@go run ${SOURCE} --config=configs/config-dev.yml

test:
	@go test -v $(GO_TEST_DIRS)

release:
	rm -rf ${RELEASE_DIR}${APP}*
	GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o ${RELEASE_DIR}/${APP}.exe main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${RELEASE_DIR}/${APP} main.go

.PHONY: build run release lint test