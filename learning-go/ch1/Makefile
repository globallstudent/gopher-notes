.DEFAULT_GOAL := build

BINARY := $(notdir $(CURDIR))

.PHONY:fmt vet build
fmt:
	go fmt ./..

vet:
	go vet ./..

build:
	go build -o $(BINARY)

.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -f $(BINARY)
	@go clean -cache -testcache -modcache