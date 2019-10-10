.PHONY: dev
dev:
	go run main.go -c ./config/dev.json

APP_NAME := requestbin
APP_VERSION := $(shell git describe --tags --always || git rev-parse HEAD)
APP_PKG := $(shell echo ${PWD} | sed -e "s\#${GOPATH}/src/\#\#g")

.PHONY: build
build:
	@mkdir -p ./dist
	go build -ldflags "-X ${APP_PKG}/app.Name=${APP_NAME} -X ${APP_PKG}/app.Version=${APP_VERSION}" -o ./dist/requestbin main.go

.PHONY: image
image:
	docker build -t ${APP_NAME}:${APP_VERSION} .
