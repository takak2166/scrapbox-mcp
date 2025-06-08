package mcpgo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/takak2166/scrapbox-mcp/internal/config"
	"github.com/takak2166/scrapbox-mcp/pkg/scrapbox"
)

// Server wraps the MCP server and Scrapbox client
type Server struct {
	mcpServer *server.MCPServer
	client    *scrapbox.Client
}

// NewServer creates a new MCP server instance
func NewServer(cfg *config.Config) *server.MCPServer {
	client := scrapbox.NewClient(cfg.ProjectName, cfg.ScrapboxSID)
	mcpSrv := server.NewMCPServer(
		"Scrapbox MCP Server",
		"1.0.0",
	)

	s := &Server{
		mcpServer: mcpSrv,
		client:    client,
	}

	s.registerTools()
	return mcpSrv
}

func (s *Server) registerTools() {
	// get_page
	getPageTool := mcp.NewTool("get_page",
		mcp.WithDescription("Get a Scrapbox page by title"),
		mcp.WithString("title", mcp.Required(), mcp.Description("Page title to retrieve")),
	)
	s.mcpServer.AddTool(getPageTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		title, err := req.RequireString("title")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		page, err := s.client.GetPage(ctx, title)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("failed to get page: %v", err)), nil
		}
		b, _ := json.Marshal(page)
		return mcp.NewToolResultText(string(b)), nil
	})

	// list_pages
	listPagesTool := mcp.NewTool("list_pages",
		mcp.WithDescription("List all Scrapbox pages"),
	)
	s.mcpServer.AddTool(listPagesTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		pages, err := s.client.ListPages(ctx)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("failed to list pages: %v", err)), nil
		}
		b, _ := json.Marshal(pages)
		return mcp.NewToolResultText(string(b)), nil
	})

	// search_pages
	searchPagesTool := mcp.NewTool("search_pages",
		mcp.WithDescription("Search Scrapbox pages"),
		mcp.WithString("query", mcp.Required(), mcp.Description("Search query")),
	)
	s.mcpServer.AddTool(searchPagesTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		query, err := req.RequireString("query")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		pages, err := s.client.SearchPages(ctx, query)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("failed to search pages: %v", err)), nil
		}
		b, _ := json.Marshal(pages)
		return mcp.NewToolResultText(string(b)), nil
	})

	// create_page_url
	createPageURLTool := mcp.NewTool("create_page_url",
		mcp.WithDescription("Generate a URL for creating a new page"),
		mcp.WithString("title", mcp.Required(), mcp.Description("Page title")),
		mcp.WithString("body_text", mcp.Description("Body text for the new page")),
	)
	s.mcpServer.AddTool(createPageURLTool, func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		title, err := req.RequireString("title")
		if err != nil {
			return mcp.NewToolResultError(err.Error()), nil
		}
		bodyText := req.GetString("body_text", "")
		url, err := s.client.CreatePageURL(ctx, title, bodyText)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("failed to generate page URL: %v", err)), nil
		}
		return mcp.NewToolResultText(url), nil
	})
}
