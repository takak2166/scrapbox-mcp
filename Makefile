.PHONY: help vendor build build-all run-mcp-go run-go-mcp run-mcp-golang run-x-tools-mcp clean

# Default target
help:
	@echo "Available targets:"
	@echo "  vendor          - Vendor golang.org/x/tools/internal/mcp to internal/x-tools-mcp"
	@echo "  build           - Build all MCP server implementations"
	@echo "  build-mcp-go    - Build mark3labs/mcp-go implementation"
	@echo "  build-go-mcp    - Build ktr0731/go-mcp implementation"
	@echo "  build-mcp-golang - Build metoro-io/mcp-golang implementation"
	@echo "  build-x-tools-mcp - Build golang.org/x/tools/internal/mcp implementation"
	@echo "  run-mcp-go      - Run mark3labs/mcp-go implementation"
	@echo "  run-go-mcp      - Run ktr0731/go-mcp implementation"
	@echo "  run-mcp-golang  - Run metoro-io/mcp-golang implementation"
	@echo "  run-x-tools-mcp - Run golang.org/x/tools/internal/mcp implementation"
	@echo "  clean           - Clean build artifacts"

# Vendor golang.org/x/tools/internal/mcp
vendor:
	@echo "Vendoring golang.org/x/tools/internal/mcp..."
	@./scripts/vendor-mcp.sh

# Build targets
build: build-mcp-go build-go-mcp build-mcp-golang build-x-tools-mcp

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

build-x-tools-mcp:
	@echo "Building golang.org/x/tools/internal/mcp implementation..."
	@mkdir -p bin
	@go build -o bin/scrapbox-mcp-x-tools cmd/x-tools-mcp/main.go

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

run-x-tools-mcp: build-x-tools-mcp
	@echo "Starting golang.org/x/tools/internal/mcp implementation..."
	@./bin/scrapbox-mcp-x-tools

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -rf internal/x-tools-mcp/ 
