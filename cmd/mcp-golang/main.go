package main

import (
	"log"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
	"github.com/takak2166/scrapbox-mcp/internal/config"
	mcpServer "github.com/takak2166/scrapbox-mcp/internal/mcp-golang"
	scrapbox "github.com/takak2166/scrapbox-mcp/pkg/scrapbox"
)

func main() {
	done := make(chan struct{})

	// Load configuration from environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	client := scrapbox.NewClient(cfg.ProjectName, cfg.ScrapboxSID)

	// Create MCP server with stdio transport
	server := mcp.NewServer(stdio.NewStdioServerTransport())

	// Register tools
	if err := mcpServer.RegisterTools(server, client); err != nil {
		log.Fatalf("Failed to register tools: %v", err)
	}

	// Start the MCP server
	if err := server.Serve(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	<-done
}
