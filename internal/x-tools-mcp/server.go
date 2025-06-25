package xtoolsmcpimpl

import (
	"context"
	"encoding/json"
	"fmt"

	mcp "github.com/takak2166/scrapbox-mcp/internal/imported/x-tools/mcp"
	"github.com/takak2166/scrapbox-mcp/pkg/scrapbox"
)

// Tool parameter structures for type-safe MCP tool calls

// GetPageParams represents arguments for the get_page tool
type GetPageParams struct {
	PageTitle string `json:"page_title" jsonschema:"required,description=Page title to retrieve"`
}

// ListPagesParams represents arguments for the list_pages tool
type ListPagesParams struct {
	// No arguments needed for list_pages
}

// SearchPagesParams represents arguments for the search_pages tool
type SearchPagesParams struct {
	Query string `json:"query" jsonschema:"required,description=Search query"`
}

// CreatePageURLParams represents arguments for the create_page_url tool
type CreatePageURLParams struct {
	PageTitle string  `json:"page_title" jsonschema:"required,description=Page title"`
	BodyText  *string `json:"body_text,omitempty" jsonschema:"description=Body text for the new page"`
}

// Tool handlers using x-tools-mcp's type-safe approach

// GetPageHandler handles the get_page tool call
func GetPageHandler(client *scrapbox.Client) mcp.ToolHandlerFor[GetPageParams, any] {
	return func(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[GetPageParams]) (*mcp.CallToolResultFor[any], error) {
		page, err := client.GetPage(ctx, params.Arguments.PageTitle)
		if err != nil {
			return &mcp.CallToolResultFor[any]{
				Content: []*mcp.Content{mcp.NewTextContent(fmt.Sprintf("Failed to get page: %v", err))},
				IsError: true,
			}, nil
		}

		pageJSON, err := json.Marshal(page)
		if err != nil {
			return &mcp.CallToolResultFor[any]{
				Content: []*mcp.Content{mcp.NewTextContent(fmt.Sprintf("Failed to marshal page: %v", err))},
				IsError: true,
			}, nil
		}

		return &mcp.CallToolResultFor[any]{
			Content: []*mcp.Content{mcp.NewTextContent(string(pageJSON))},
		}, nil
	}
}

// ListPagesHandler handles the list_pages tool call
func ListPagesHandler(client *scrapbox.Client) mcp.ToolHandlerFor[ListPagesParams, any] {
	return func(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[ListPagesParams]) (*mcp.CallToolResultFor[any], error) {
		pages, err := client.ListPages(ctx)
		if err != nil {
			return &mcp.CallToolResultFor[any]{
				Content: []*mcp.Content{mcp.NewTextContent(fmt.Sprintf("Failed to list pages: %v", err))},
				IsError: true,
			}, nil
		}

		pagesJSON, err := json.Marshal(pages)
		if err != nil {
			return &mcp.CallToolResultFor[any]{
				Content: []*mcp.Content{mcp.NewTextContent(fmt.Sprintf("Failed to marshal pages: %v", err))},
				IsError: true,
			}, nil
		}

		return &mcp.CallToolResultFor[any]{
			Content: []*mcp.Content{mcp.NewTextContent(string(pagesJSON))},
		}, nil
	}
}

// SearchPagesHandler handles the search_pages tool call
func SearchPagesHandler(client *scrapbox.Client) mcp.ToolHandlerFor[SearchPagesParams, any] {
	return func(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[SearchPagesParams]) (*mcp.CallToolResultFor[any], error) {
		pages, err := client.SearchPages(ctx, params.Arguments.Query)
		if err != nil {
			return &mcp.CallToolResultFor[any]{
				Content: []*mcp.Content{mcp.NewTextContent(fmt.Sprintf("Failed to search pages: %v", err))},
				IsError: true,
			}, nil
		}

		pagesJSON, err := json.Marshal(pages)
		if err != nil {
			return &mcp.CallToolResultFor[any]{
				Content: []*mcp.Content{mcp.NewTextContent(fmt.Sprintf("Failed to marshal pages: %v", err))},
				IsError: true,
			}, nil
		}

		return &mcp.CallToolResultFor[any]{
			Content: []*mcp.Content{mcp.NewTextContent(string(pagesJSON))},
		}, nil
	}
}

// CreatePageURLHandler handles the create_page_url tool call
func CreatePageURLHandler(client *scrapbox.Client) mcp.ToolHandlerFor[CreatePageURLParams, any] {
	return func(ctx context.Context, ss *mcp.ServerSession, params *mcp.CallToolParamsFor[CreatePageURLParams]) (*mcp.CallToolResultFor[any], error) {
		bodyText := ""
		if params.Arguments.BodyText != nil {
			bodyText = *params.Arguments.BodyText
		}

		pageURL, err := client.CreatePageURL(ctx, params.Arguments.PageTitle, bodyText)
		if err != nil {
			return &mcp.CallToolResultFor[any]{
				Content: []*mcp.Content{mcp.NewTextContent(fmt.Sprintf("Failed to generate page URL: %v", err))},
				IsError: true,
			}, nil
		}

		return &mcp.CallToolResultFor[any]{
			Content: []*mcp.Content{mcp.NewTextContent(pageURL)},
		}, nil
	}
}

// CreateServer creates and configures the MCP server with all Scrapbox tools
func CreateServer(client *scrapbox.Client) *mcp.Server {
	server := mcp.NewServer("scrapbox-mcp-x-tools", "v1.0.0", &mcp.ServerOptions{
		Instructions: "Scrapbox MCP server implementation using golang.org/x/tools/internal/mcp",
	})

	// Register all tools with type-safe handlers
	server.AddTools(
		mcp.NewTool("get_page", "Get a Scrapbox page by title", GetPageHandler(client)),
		mcp.NewTool("list_pages", "Get a list of pages in the project (max 1000 pages)", ListPagesHandler(client)),
		mcp.NewTool("search_pages", "Full-text search across all pages in the project (max 100 pages)", SearchPagesHandler(client)),
		mcp.NewTool("create_page_url", "Generate a URL for creating a new page", CreatePageURLHandler(client)),
	)

	return server
}
