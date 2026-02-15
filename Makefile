BINARY := homewizard-prometheus-exporter
BUILD_DIR := $(shell pwd)/build
IMAGE := quay.io/touchardv/homewizard-prometheus-exporter
GOARCH := $(shell go env GOARCH)
GOOS := $(shell go env GOOS)
SOURCES := $(shell find . -name '*.go')
TAG := latest
TARGET ?= $(shell uname -m)

ifeq ($(GOARCH), arm64)
 DOCKER_BUILDX_PLATFORM := linux/arm64/v8
else ifeq ($(GOARCH), amd64)
 DOCKER_BUILDX_PLATFORM := linux/amd64
endif

.DEFAULT_GOAL := build
.PHONY: build
build: $(BUILD_DIR)/$(BINARY)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

build-image: $(BINARY)-linux-$(GOARCH)
	docker buildx build --progress plain \
	--platform $(DOCKER_BUILDX_PLATFORM) \
	--tag $(IMAGE):$(TAG) --load -f deployment/Dockerfile .

$(BUILD_DIR)/$(BINARY): $(SOURCES)
	go mod tidy
	go build -o $(BUILD_DIR)/$(BINARY) .

$(BINARY)-linux-$(GOARCH): $(BUILD_DIR) $(SOURCES)
	go mod tidy
	GOOS=linux GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(BINARY)-linux-$(GOARCH) .

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)
	go clean

run: $(BUILD_DIR)/$(BINARY)
	source local.env && $(BUILD_DIR)/$(BINARY) export-metrics

run-image: #build-image
	docker run -it --rm -e HOMEWIZARD_PROMETHEUS_EXPORTER_URL=http://foo.bar --entrypoint=/bin/sh $(IMAGE):$(TAG)

test:
	go test -v -cover -timeout 10s ./...