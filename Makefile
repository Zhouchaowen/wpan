SHELL := /bin/bash

# ==============================================================================
# $(shell git rev-parse --short HEAD)
VERSION := 1.0
WEB_FILE_PATH := /wpan/data
CONF_PATH := $(shell pwd)

all: build-wpan run-mysql run-wpan

build-wpan:
	docker build \
		-f build/Dockerfile \
		-t wpan-api-amd64:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` .

run-mysql:
	docker run \
		-itd --restart=always \
		--network host \
		--name wpan-mysql \
		-e MYSQL_ROOT_PASSWORD='123456' \
		-e MYSQL_DATABASE='file_store' \
		mysql:5.7.41

run-wpan:
	docker run \
		-itd --restart=always \
		--network host \
		--name wpan-api \
		-v /etc/localtime:/etc/localtime:ro \
		-v $(CONF_PATH)/conf:/conf \
		-v /data/wpan:$(WEB_FILE_PATH) \
		wpan-api-amd64:$(VERSION)


update-wpan: build-wpan clear-wpan run-wpan

clear-wpan:
	docker rm -f wpan
	docker rm -f wpan-mysql
	rm -rf /data/bak && mkdir /data/bak/ && mv /data/wpan* /data/bak/

# ==============================================================================
# Running tests within the local computer

test:
	go test ./... -count=1
	staticcheck -checks=all ./...

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor


