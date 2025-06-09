package main

import (
	"context"
	"log"

	mcp "github.com/ktr0731/go-mcp"
	"github.com/takak2166/scrapbox-mcp/internal/config"
	mcpServer "github.com/takak2166/scrapbox-mcp/internal/go-mcp"
	scrapbox "github.com/takak2166/scrapbox-mcp/pkg/scrapbox"
	"golang.org/x/exp/jsonrpc2"
)

func main() {
	// Load configuration from environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	client := scrapbox.NewClient(cfg.ProjectName, cfg.ScrapboxSID)
	toolHandler := mcpServer.NewToolHandler(client)
	handler := mcpServer.NewHandler(toolHandler)

	// Start the MCP server with stdio transport
	ctx, listener, binder := mcp.NewStdioTransport(context.Background(), handler, nil)
	srv, err := jsonrpc2.Serve(ctx, listener, binder)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	srv.Wait()
}
