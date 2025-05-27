# Learning Go - Chapter 1

Welcome to the first chapter of learning Go! This guide introduces the Go programming language, demonstrates a simple "Hello, World!" program, and provides a basic Makefile for building your Go code.

## About Go

Go (or Golang) is an open-source programming language developed by Google. It is known for its simplicity, efficiency, and strong support for concurrency.

Key features:
- Statically typed and compiled
- Fast compilation and execution
- Built-in concurrency with goroutines
- Simple and readable syntax

## Hello World in Go

Below is a basic "Hello, World!" program in Go:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### How to Run

1. Save the code above in a file named `main.go`.
2. Run the program using:

   ```sh
   go run main.go
   ```

## Useful Go Commands

- `go run <file.go>`: Compile and run a Go program.
- `go build`: Compile the packages and dependencies.
- `go fmt ./...`: Format all Go source files in the current module.
- `go vet ./...`: Analyze code and report suspicious constructs.
- `go test`: Run tests in the current package.
- `go mod init <module>`: Initialize a new Go module.
- `go get <package>`: Add or update dependencies.

## Formatting and Vetting

Go provides built-in tools to help maintain code quality:

- **go fmt**: Automatically formats your Go source code according to the standard style.
- **go vet**: Examines Go source code and reports suspicious constructs, such as mistakes that could lead to bugs.

## Basic Makefile

A Makefile can help automate building, formatting, vetting, and cleaning your Go programs. Here is an improved example:

```makefile
.DEFAULT_GOAL := build
.PHONY: clean fmt vet build

fmt:
    go fmt ./...

vet: fmt
    go vet ./...

build: vet
    go build -o hello hello.go

clean:
    rm -rf hello
    go clean --cache
```

- `make fmt` formats all Go files.
- `make vet` runs formatting and then static analysis.
- `make build` formats, vets, and builds the program.
- `make clean` removes the binary and cleans the build cache.

---