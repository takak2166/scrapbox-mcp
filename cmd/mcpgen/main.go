package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ktr0731/go-mcp/codegen"
)

type GetPageInput struct {
	PageTitle string `json:"page_title" jsonschema:"description=Page title to retrieve,required"`
}

type SearchPagesInput struct {
	Query string `json:"query" jsonschema:"description=Search query,required"`
}

type CreatePageInput struct {
	PageTitle string  `json:"page_title" jsonschema:"description=Title of the new page,required"`
	BodyText  *string `json:"body_text" jsonschema:"description=Optional body text for the new page"`
}

func main() {
	outDir := "internal/mcp"
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		log.Fatalf("failed to create output directory: %v", err)
	}
	f, err := os.Create(filepath.Join(outDir, "mcp.gen.go"))
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	defer f.Close()

	def := &codegen.ServerDefinition{
		Capabilities: codegen.ServerCapabilities{
			Tools:   &codegen.ToolCapability{},
			Logging: &codegen.LoggingCapability{},
		},
		Implementation: codegen.Implementation{
			Name:    "Scrapbox MCP Server",
			Version: "1.0.0",
		},
		Tools: []codegen.Tool{
			{
				Name:        "get_page",
				Description: "Get page content from Scrapbox",
				InputSchema: GetPageInput{},
			},
			{
				Name:        "list_pages",
				Description: "Get a list of pages in the project (max 1000 pages)",
				InputSchema: struct{}{},
			},
			{
				Name:        "search_pages",
				Description: "Full-text search across all pages in the project (max 100 pages)",
				InputSchema: SearchPagesInput{},
			},
			{
				Name:        "create_page",
				Description: "Create a new page",
				InputSchema: CreatePageInput{},
			},
		},
	}

	if err := codegen.Generate(f, def, "scrapbox"); err != nil {
		log.Fatalf("failed to generate code: %v", err)
	}
}
