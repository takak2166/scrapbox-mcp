package officialmcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/takak2166/scrapbox-mcp/pkg/scrapbox"
)

// Tool argument structures for type-safe MCP tool calls

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
	BodyText  *string `json:"body_text" jsonschema:"description=Body text for the new page"`
}

// Server represents the MCP server with Scrapbox tools
type Server struct {
	client *scrapbox.Client
	server *mcp.Server
}

// NewServer creates a new MCP server with Scrapbox tools
func NewServer(client *scrapbox.Client) *Server {
	server := mcp.NewServer("Scrapbox MCP Server", "1.0.0", nil)

	s := &Server{
		client: client,
		server: server,
	}

	// Register tools
	s.registerTools()

	return s
}

// GetServer returns the underlying MCP server for transport configuration
func (s *Server) GetServer() *mcp.Server {
	return s.server
}

// registerTools registers all Scrapbox tools with the MCP server
func (s *Server) registerTools() {
	// Register get_page tool
	s.server.AddTools(
		mcp.NewServerTool("get_page", "Get a Scrapbox page by title", s.handleGetPage,
			mcp.Input(
				mcp.Property("page_title", mcp.Description("Page title to retrieve")),
			),
		),
		mcp.NewServerTool("list_pages", "Get a list of pages in the project (max 1000 pages)", s.handleListPages,
			mcp.Input(),
		),
		mcp.NewServerTool("search_pages", "Full-text search across all pages in the project (max 100 pages)", s.handleSearchPages,
			mcp.Input(
				mcp.Property("query", mcp.Description("Search query")),
			),
		),
		mcp.NewServerTool("create_page_url", "Generate a URL for creating a new page", s.handleCreatePageURL,
			mcp.Input(
				mcp.Property("page_title", mcp.Description("Page title")),
				mcp.Property("body_text", mcp.Description("Body text for the new page")),
			),
		),
	)
}

// handleGetPage handles the get_page tool call
func (s *Server) handleGetPage(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[GetPageParams]) (*mcp.CallToolResultFor[any], error) {
	page, err := s.client.GetPage(ctx, params.Arguments.PageTitle)
	if err != nil {
		return nil, fmt.Errorf("Failed to get page: %w", err)
	}

	pageJSON, err := json.Marshal(page)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal page: %w", err)
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: string(pageJSON)}},
	}, nil
}

// handleListPages handles the list_pages tool call
func (s *Server) handleListPages(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[ListPagesParams]) (*mcp.CallToolResultFor[any], error) {
	pages, err := s.client.ListPages(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to list pages: %w", err)
	}

	pagesJSON, err := json.Marshal(pages)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal pages: %w", err)
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: string(pagesJSON)}},
	}, nil
}

// handleSearchPages handles the search_pages tool call
func (s *Server) handleSearchPages(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[SearchPagesParams]) (*mcp.CallToolResultFor[any], error) {
	pages, err := s.client.SearchPages(ctx, params.Arguments.Query)
	if err != nil {
		return nil, fmt.Errorf("Failed to search pages: %w", err)
	}

	pagesJSON, err := json.Marshal(pages)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal pages: %w", err)
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: string(pagesJSON)}},
	}, nil
}

// handleCreatePageURL handles the create_page_url tool call
func (s *Server) handleCreatePageURL(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[CreatePageURLParams]) (*mcp.CallToolResultFor[any], error) {
	bodyText := ""
	if params.Arguments.BodyText != nil {
		bodyText = *params.Arguments.BodyText
	}

	pageURL, err := s.client.CreatePageURL(ctx, params.Arguments.PageTitle, bodyText)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate page URL: %w", err)
	}

	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: pageURL}},
	}, nil
}
