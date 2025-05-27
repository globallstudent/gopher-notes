.DEFAULT_GOAL := build
.PHONY:clean fmt vet build


fmt:
	go fmt ./...


vet: fmt
	go vet ./...

build: vet
	go build -o hello hello.go

clean:
	rm -rf hello
	go clean --cache