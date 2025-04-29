package scrapbox

import (
	"context"
	"fmt"
	"log"

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
	log.Printf("Handling get_page request for title: %s", req.PageTitle)
	page, err := s.client.GetPage(ctx, req.PageTitle)
	if err != nil {
		log.Printf("Failed to get page: %v", err)
		return nil, fmt.Errorf("failed to get page: %w", err)
	}
	log.Printf("Successfully retrieved page: %s, Lines: %v", page.Title, page.Lines)
	return &mcp.CallToolResult{
		Content: []mcp.CallToolContent{
			mcp.TextContent{Text: fmt.Sprintf("Title: %s\nLines: %v", page.Title, page.Lines)},
		},
	}, nil
}

// HandleToolListPages handles list_pages tool requests.
func (s *Server) HandleToolListPages(ctx context.Context, req *ToolListPagesRequest) (*mcp.CallToolResult, error) {
	log.Printf("Handling list_pages request")
	pages, err := s.client.ListPages(ctx)
	if err != nil {
		log.Printf("Failed to list pages: %v", err)
		return nil, fmt.Errorf("failed to list pages: %w", err)
	}
	log.Printf("Successfully retrieved %d pages", len(pages.Pages))
	return &mcp.CallToolResult{
		Content: []mcp.CallToolContent{
			mcp.TextContent{Text: fmt.Sprintf("Pages: %v", pages)},
		},
	}, nil
}

// HandleToolSearchPages handles search_pages tool requests.
func (s *Server) HandleToolSearchPages(ctx context.Context, req *ToolSearchPagesRequest) (*mcp.CallToolResult, error) {
	log.Printf("Handling search_pages request with query: %s", req.Query)
	pages, err := s.client.SearchPages(ctx, req.Query)
	if err != nil {
		log.Printf("Failed to search pages: %v", err)
		return nil, fmt.Errorf("failed to search pages: %w", err)
	}
	log.Printf("Successfully found %d pages matching query", len(pages.Pages))
	return &mcp.CallToolResult{
		Content: []mcp.CallToolContent{
			mcp.TextContent{Text: fmt.Sprintf("Search results: %v", pages)},
		},
	}, nil
}

// HandleToolCreatePage handles create_page tool requests.
func (s *Server) HandleToolCreatePage(ctx context.Context, req *ToolCreatePageRequest) (*mcp.CallToolResult, error) {
	log.Printf("Handling create_page request for title: %s", req.PageTitle)
	bodyText := ""
	if req.BodyText != nil {
		bodyText = *req.BodyText
	}
	page, err := s.client.CreatePage(ctx, req.PageTitle, bodyText)
	if err != nil {
		log.Printf("Failed to create page: %v", err)
		return nil, fmt.Errorf("failed to create page: %w", err)
	}
	log.Printf("Successfully created page: %s", page.Title)
	return &mcp.CallToolResult{
		Content: []mcp.CallToolContent{
			mcp.TextContent{Text: fmt.Sprintf("Created page: %v", page)},
		},
	}, nil
}
