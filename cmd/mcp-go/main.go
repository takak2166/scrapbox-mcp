// Package main implements the Scrapbox MCP server using mark3labs/mcp-go.
// It provides a command-line interface for interacting with Scrapbox through MCP.
package main

import (
	"log"

	"github.com/mark3labs/mcp-go/server"
	"github.com/takak2166/scrapbox-mcp/internal/config"
	mcpServer "github.com/takak2166/scrapbox-mcp/internal/mcp-go"
	scrapbox "github.com/takak2166/scrapbox-mcp/pkg/scrapbox"
)

func main() {
	// Load configuration from environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	client := scrapbox.NewClient(cfg.ProjectName, cfg.ScrapboxSID)

	// Create MCP server
	mcpServer := mcpServer.NewServer(client)

	// Start the MCP server with stdio transport
	if err := server.ServeStdio(mcpServer); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
