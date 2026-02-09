.PHONY: test bench lint fmt vet clean help build install

# Default target
all: test

# Build the CLI tool
build:
	go build -o bin/porter ./cmd/porter

# Install the CLI tool to $GOPATH/bin
install:
	go install ./cmd/porter

# Run tests
test:
	go test -v ./...

# Run tests with coverage
coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run benchmarks
bench:
	go test -bench=. -benchmem ./...

# Format code
fmt:
	gofmt -s -w .

# Run go vet
vet:
	go vet ./...

# Run linter (requires golangci-lint)
lint:
	@which golangci-lint > /dev/null || (echo "golangci-lint not installed. Install: https://golangci-lint.run/usage/install/" && exit 1)
	golangci-lint run

# Clean build artifacts
clean:
	go clean
	rm -f coverage.out coverage.html
	rm -rf bin/

# Show help
help:
	@echo "Available targets:"
	@echo "  make build     - Build the CLI tool to bin/porter"
	@echo "  make install   - Install CLI tool to \$$GOPATH/bin"
	@echo "  make test      - Run tests"
	@echo "  make coverage  - Run tests with coverage report"
	@echo "  make bench     - Run benchmarks"
	@echo "  make fmt       - Format code with gofmt"
	@echo "  make vet       - Run go vet"
	@echo "  make lint      - Run golangci-lint"
	@echo "  make clean     - Clean build artifacts"
