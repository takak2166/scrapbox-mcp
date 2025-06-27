.PHONY: help build build-all run-mcp-go run-go-mcp run-mcp-golang run-official-mcp clean

# Default target
help:
	@echo "Available targets:"
	@echo "  build           - Build all MCP server implementations"
	@echo "  build-mcp-go    - Build mark3labs/mcp-go implementation"
	@echo "  build-go-mcp    - Build ktr0731/go-mcp implementation"
	@echo "  build-mcp-golang - Build metoro-io/mcp-golang implementation"
	@echo "  build-official-mcp - Build official Go SDK implementation"
	@echo "  run-mcp-go      - Run mark3labs/mcp-go implementation"
	@echo "  run-go-mcp      - Run ktr0731/go-mcp implementation"
	@echo "  run-mcp-golang  - Run metoro-io/mcp-golang implementation"
	@echo "  run-official-mcp - Run official Go SDK implementation"
	@echo "  clean           - Clean build artifacts"

# Build targets
build: build-mcp-go build-go-mcp build-mcp-golang build-official-mcp

build-mcp-go:
	@echo "Building mark3labs/mcp-go implementation..."
	@mkdir -p bin
	@go build -o bin/scrapbox-mcp cmd/mcp-go/main.go

build-go-mcp:
	@echo "Building ktr0731/go-mcp implementation..."
	@mkdir -p bin
	@go build -o bin/scrapbox-mcp-go cmd/go-mcp/main.go

build-mcp-golang:
	@echo "Building metoro-io/mcp-golang implementation..."
	@mkdir -p bin
	@go build -o bin/scrapbox-mcp-golang cmd/mcp-golang/main.go

build-official-mcp:
	@echo "Building official Go SDK implementation..."
	@mkdir -p bin
	@go build -o bin/scrapbox-mcp-official cmd/official-mcp/main.go

# Run targets
run-mcp-go: build-mcp-go
	@echo "Starting mark3labs/mcp-go implementation..."
	@./bin/scrapbox-mcp

run-go-mcp: build-go-mcp
	@echo "Starting ktr0731/go-mcp implementation..."
	@./bin/scrapbox-mcp-go

run-mcp-golang: build-mcp-golang
	@echo "Starting metoro-io/mcp-golang implementation..."
	@./bin/scrapbox-mcp-golang

run-official-mcp: build-official-mcp
	@echo "Starting official Go SDK implementation..."
	@./bin/scrapbox-mcp-official

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/ 
 