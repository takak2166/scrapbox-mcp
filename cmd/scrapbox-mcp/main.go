package main

import (
	"context"
	"log"

	mcp "github.com/ktr0731/go-mcp"
	"github.com/takak2166/scrapbox-mcp/internal/config"
	scrapbox "github.com/takak2166/scrapbox-mcp/internal/mcp"
	"golang.org/x/exp/jsonrpc2"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	client := scrapbox.NewClient(cfg.ProjectName, cfg.ScrapboxSID)
	server := scrapbox.NewServer(cfg.ProjectName, client)
	handler := scrapbox.NewHandler(server)

	ctx, listener, binder := mcp.NewStdioTransport(context.Background(), handler, nil)
	srv, err := jsonrpc2.Serve(ctx, listener, binder)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	srv.Wait()
}
