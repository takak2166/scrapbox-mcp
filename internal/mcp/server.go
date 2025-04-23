package scrapbox

import (
	"context"
	"fmt"

	mcp "github.com/ktr0731/go-mcp"
)

// Server implements ServerToolHandler interface.
type Server struct {
	projectName string
	client      *Client
}

// NewServer creates a new Server instance.
func NewServer(projectName string, client *Client) *Server {
	return &Server{
		projectName: projectName,
		client:      client,
	}
}

// HandleToolGetPage handles get_page tool requests.
func (s *Server) HandleToolGetPage(ctx context.Context, req *ToolGetPageRequest) (*mcp.CallToolResult, error) {
	page, err := s.client.GetPage(ctx, req.PageTitle)
	if err != nil {
		return nil, fmt.Errorf("failed to get page: %w", err)
	}
	return &mcp.CallToolResult{
		Content: []mcp.CallToolContent{
			mcp.TextContent{Text: fmt.Sprintf("Title: %s\nText: %s", page.Title, page.Text)},
		},
	}, nil
}

// HandleToolListPages handles list_pages tool requests.
func (s *Server) HandleToolListPages(ctx context.Context, req *ToolListPagesRequest) (*mcp.CallToolResult, error) {
	pages, err := s.client.ListPages(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list pages: %w", err)
	}
	return &mcp.CallToolResult{
		Content: []mcp.CallToolContent{
			mcp.TextContent{Text: fmt.Sprintf("Pages: %v", pages)},
		},
	}, nil
}

// HandleToolSearchPages handles search_pages tool requests.
func (s *Server) HandleToolSearchPages(ctx context.Context, req *ToolSearchPagesRequest) (*mcp.CallToolResult, error) {
	pages, err := s.client.SearchPages(ctx, req.Query)
	if err != nil {
		return nil, fmt.Errorf("failed to search pages: %w", err)
	}
	return &mcp.CallToolResult{
		Content: []mcp.CallToolContent{
			mcp.TextContent{Text: fmt.Sprintf("Search results: %v", pages)},
		},
	}, nil
}

// HandleToolCreatePage handles create_page tool requests.
func (s *Server) HandleToolCreatePage(ctx context.Context, req *ToolCreatePageRequest) (*mcp.CallToolResult, error) {
	bodyText := ""
	if req.BodyText != nil {
		bodyText = *req.BodyText
	}
	page, err := s.client.CreatePage(ctx, req.PageTitle, bodyText)
	if err != nil {
		return nil, fmt.Errorf("failed to create page: %w", err)
	}
	return &mcp.CallToolResult{
		Content: []mcp.CallToolContent{
			mcp.TextContent{Text: fmt.Sprintf("Created page: %v", page)},
		},
	}, nil
}
