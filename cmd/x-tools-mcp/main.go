package main

import (
	"context"
	"log"

	"github.com/takak2166/scrapbox-mcp/internal/config"
	mcp "github.com/takak2166/scrapbox-mcp/internal/imported/x-tools/mcp"
	xtoolsmcpimpl "github.com/takak2166/scrapbox-mcp/internal/x-tools-mcp"
	"github.com/takak2166/scrapbox-mcp/pkg/scrapbox"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create Scrapbox client
	client := scrapbox.NewClient(cfg.ProjectName, cfg.ScrapboxSID)

	// Create MCP server with type-safe tools
	server := xtoolsmcpimpl.CreateServer(client)

	// Run server over stdin/stdout
	ctx := context.Background()
	transport := mcp.NewStdIOTransport()

	log.Printf("Starting Scrapbox MCP server (x-tools-mcp implementation) for project: %s", cfg.ProjectName)

	if err := server.Run(ctx, transport); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
