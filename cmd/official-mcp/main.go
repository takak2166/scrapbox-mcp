package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/takak2166/scrapbox-mcp/internal/config"
	mcpServer "github.com/takak2166/scrapbox-mcp/internal/official-mcp"
	"github.com/takak2166/scrapbox-mcp/pkg/scrapbox"
)

func main() {
	// Load configuration from environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	client := scrapbox.NewClient(cfg.ProjectName, cfg.ScrapboxSID)
	server := mcpServer.NewServer(client)

	// Start the MCP server with stdio transport
	if err := server.GetServer().Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
