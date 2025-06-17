package mcpgolang

import (
	"context"
	"encoding/json"
	"fmt"

	mcp "github.com/metoro-io/mcp-golang"
	"github.com/takak2166/scrapbox-mcp/pkg/scrapbox"
)

// Tool argument structures for type-safe MCP tool calls

// GetPageArgs represents arguments for the get_page tool
type GetPageArgs struct {
	PageTitle string `json:"page_title" jsonschema:"required,description=Page title to retrieve"`
}

// ListPagesArgs represents arguments for the list_pages tool
type ListPagesArgs struct {
	// No arguments needed for list_pages
}

// SearchPagesArgs represents arguments for the search_pages tool
type SearchPagesArgs struct {
	Query string `json:"query" jsonschema:"required,description=Search query"`
}

// CreatePageURLArgs represents arguments for the create_page_url tool
type CreatePageURLArgs struct {
	PageTitle string  `json:"page_title" jsonschema:"required,description=Page title"`
	BodyText  *string `json:"body_text" jsonschema:"description=Body text for the new page"`
}

// RegisterTools registers all Scrapbox tools with the MCP server
func RegisterTools(server *mcp.Server, client *scrapbox.Client) error {
	// Register get_page tool
	err := server.RegisterTool("get_page", "Get a Scrapbox page by title", func(args GetPageArgs) (*mcp.ToolResponse, error) {
		page, err := client.GetPage(context.Background(), args.PageTitle)
		if err != nil {
			return nil, fmt.Errorf("Failed to get page: %w", err)
		}

		pageJSON, err := json.Marshal(page)
		if err != nil {
			return nil, fmt.Errorf("Failed to marshal page: %w", err)
		}

		return mcp.NewToolResponse(mcp.NewTextContent(string(pageJSON))), nil
	})
	if err != nil {
		return fmt.Errorf("Failed to register get_page tool: %w", err)
	}

	// Register list_pages tool
	err = server.RegisterTool("list_pages", "Get a list of pages in the project (max 1000 pages)", func(args ListPagesArgs) (*mcp.ToolResponse, error) {
		pages, err := client.ListPages(context.Background())
		if err != nil {
			return nil, fmt.Errorf("Failed to list pages: %w", err)
		}

		pagesJSON, err := json.Marshal(pages)
		if err != nil {
			return nil, fmt.Errorf("Failed to marshal pages: %w", err)
		}

		return mcp.NewToolResponse(mcp.NewTextContent(string(pagesJSON))), nil
	})
	if err != nil {
		return fmt.Errorf("Failed to register list_pages tool: %w", err)
	}

	// Register search_pages tool
	err = server.RegisterTool("search_pages", "Full-text search across all pages in the project (max 100 pages)", func(args SearchPagesArgs) (*mcp.ToolResponse, error) {
		pages, err := client.SearchPages(context.Background(), args.Query)
		if err != nil {
			return nil, fmt.Errorf("Failed to search pages: %w", err)
		}

		pagesJSON, err := json.Marshal(pages)
		if err != nil {
			return nil, fmt.Errorf("Failed to marshal pages: %w", err)
		}

		return mcp.NewToolResponse(mcp.NewTextContent(string(pagesJSON))), nil
	})
	if err != nil {
		return fmt.Errorf("Failed to register search_pages tool: %w", err)
	}

	// Register create_page_url tool
	err = server.RegisterTool("create_page_url", "Generate a URL for creating a new page", func(args CreatePageURLArgs) (*mcp.ToolResponse, error) {
		bodyText := ""
		if args.BodyText != nil {
			bodyText = *args.BodyText
		}

		pageURL, err := client.CreatePageURL(context.Background(), args.PageTitle, bodyText)
		if err != nil {
			return nil, fmt.Errorf("Failed to generate page URL: %w", err)
		}

		return mcp.NewToolResponse(mcp.NewTextContent(pageURL)), nil
	})
	if err != nil {
		return fmt.Errorf("Failed to register create_page_url tool: %w", err)
	}

	return nil
}
