package scrapbox

import (
	"context"
	"encoding/json"
	"fmt"

	mcp "github.com/ktr0731/go-mcp"
	"github.com/takak2166/scrapbox-mcp/pkg/scrapbox"
)

// ToolHandler implements ServerToolHandler interface.
type ToolHandler struct {
	client *scrapbox.Client
}

// NewToolHandler creates a new ToolHandler instance.
func NewToolHandler(client *scrapbox.Client) *ToolHandler {
	return &ToolHandler{
		client: client,
	}
}

// HandleToolGetPage handles get_page tool requests.
func (h *ToolHandler) HandleToolGetPage(ctx context.Context, req *ToolGetPageRequest) (*mcp.CallToolResult, error) {
	page, err := h.client.GetPage(ctx, req.PageTitle)
	if err != nil {
		return nil, fmt.Errorf("Failed to get page: %w", err)
	}
	b, err := json.Marshal(page)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal page: %w", err)
	}
	return &mcp.CallToolResult{
		Content: []mcp.CallToolContent{
			mcp.TextContent{Text: string(b)},
		},
	}, nil
}

// HandleToolListPages handles list_pages tool requests.
func (h *ToolHandler) HandleToolListPages(ctx context.Context, req *ToolListPagesRequest) (*mcp.CallToolResult, error) {
	pages, err := h.client.ListPages(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to list pages: %w", err)
	}
	b, err := json.Marshal(pages)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal pages: %w", err)
	}
	return &mcp.CallToolResult{
		Content: []mcp.CallToolContent{
			mcp.TextContent{Text: string(b)},
		},
	}, nil
}

// HandleToolSearchPages handles search_pages tool requests.
func (h *ToolHandler) HandleToolSearchPages(ctx context.Context, req *ToolSearchPagesRequest) (*mcp.CallToolResult, error) {
	pages, err := h.client.SearchPages(ctx, req.Query)
	if err != nil {
		return nil, fmt.Errorf("Failed to search pages: %w", err)
	}
	b, err := json.Marshal(pages)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal pages: %w", err)
	}
	return &mcp.CallToolResult{
		Content: []mcp.CallToolContent{
			mcp.TextContent{Text: string(b)},
		},
	}, nil
}

// HandleToolCreatePageUrl handles create_page_url tool requests.
func (h *ToolHandler) HandleToolCreatePageUrl(ctx context.Context, req *ToolCreatePageUrlRequest) (*mcp.CallToolResult, error) {
	bodyText := ""
	if req.BodyText != nil {
		bodyText = *req.BodyText
	}
	pageURL, err := h.client.CreatePageURL(ctx, req.PageTitle, bodyText)
	if err != nil {
		return nil, fmt.Errorf("Failed to generate page URL: %w", err)
	}
	return &mcp.CallToolResult{
		Content: []mcp.CallToolContent{
			mcp.TextContent{Text: pageURL},
		},
	}, nil
}
