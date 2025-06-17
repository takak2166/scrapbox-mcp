package scrapbox

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/takak2166/scrapbox-mcp/internal/errors"
)

// Client is a Scrapbox API client.
type Client struct {
	httpClient  *http.Client
	baseURL     string
	projectName string
	cookie      string
}

// Page represents a Scrapbox page.
type Page struct {
	Title string `json:"title"`
	Lines []Line `json:"lines"`
}

// Line represents a line of text in a Scrapbox page.
type Line struct {
	Text    string `json:"text"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
}

// PageList represents a list of Scrapbox pages.
type PageList struct {
	Pages []Page `json:"pages"`
}

// SearchPage represents a page in search results.
type SearchPage struct {
	Title string   `json:"title"`
	Lines []string `json:"lines"`
}

// SearchPageList represents a list of search results.
type SearchPageList struct {
	Pages []SearchPage `json:"pages"`
}

// NewClient creates a new Scrapbox API client.
func NewClient(projectName, cookie string) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		baseURL:     "https://scrapbox.io/api",
		projectName: projectName,
		cookie:      cookie,
	}
}

// GetPage retrieves a page by title.
func (c *Client) GetPage(ctx context.Context, title string) (*Page, error) {
	endpoint := fmt.Sprintf("%s/pages/%s/%s", c.baseURL, c.projectName, url.PathEscape(title))
	log.Printf("GET request to %s", endpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, &errors.ScrapboxError{Code: errors.ErrServerError, Message: "Failed to create request", Err: err}
	}
	req.Header.Set("Cookie", fmt.Sprintf("connect.sid=%s", c.cookie))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &errors.ScrapboxError{Code: errors.ErrServerError, Message: "Failed to send request", Err: err}
	}
	defer resp.Body.Close()

	log.Printf("Response status code: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return nil, &errors.ScrapboxError{Code: resp.StatusCode, Message: "unexpected status code", Err: nil}
	}

	var page Page
	if err := json.NewDecoder(resp.Body).Decode(&page); err != nil {
		return nil, &errors.ScrapboxError{Code: errors.ErrServerError, Message: "Failed to decode response", Err: err}
	}

	return &page, nil
}

// ListPages retrieves a list of pages.
func (c *Client) ListPages(ctx context.Context) (*PageList, error) {
	endpoint := fmt.Sprintf("%s/pages/%s", c.baseURL, c.projectName)
	log.Printf("GET request to %s", endpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, &errors.ScrapboxError{Code: errors.ErrServerError, Message: "Failed to create request", Err: err}
	}
	req.Header.Set("Cookie", fmt.Sprintf("connect.sid=%s", c.cookie))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &errors.ScrapboxError{Code: errors.ErrServerError, Message: "Failed to send request", Err: err}
	}
	defer resp.Body.Close()

	log.Printf("Response status code: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return nil, &errors.ScrapboxError{Code: resp.StatusCode, Message: "unexpected status code", Err: nil}
	}

	var pageList PageList
	if err := json.NewDecoder(resp.Body).Decode(&pageList); err != nil {
		return nil, &errors.ScrapboxError{Code: errors.ErrServerError, Message: "Failed to decode response", Err: err}
	}

	return &pageList, nil
}

// SearchPages searches pages by query.
func (c *Client) SearchPages(ctx context.Context, query string) (*SearchPageList, error) {
	endpoint := fmt.Sprintf("%s/pages/%s/search/query?q=%s", c.baseURL, c.projectName, url.QueryEscape(query))
	log.Printf("GET request to %s", endpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, &errors.ScrapboxError{Code: errors.ErrServerError, Message: "Failed to create request", Err: err}
	}
	req.Header.Set("Cookie", fmt.Sprintf("connect.sid=%s", c.cookie))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, &errors.ScrapboxError{Code: errors.ErrServerError, Message: "Failed to send request", Err: err}
	}
	defer resp.Body.Close()

	log.Printf("Response status code: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return nil, &errors.ScrapboxError{Code: resp.StatusCode, Message: "unexpected status code", Err: nil}
	}

	var pageList SearchPageList
	if err := json.NewDecoder(resp.Body).Decode(&pageList); err != nil {
		return nil, &errors.ScrapboxError{Code: errors.ErrServerError, Message: "Failed to decode response", Err: err}
	}

	return &pageList, nil
}

// CreatePageURL generates a URL for creating a new page.
func (c *Client) CreatePageURL(ctx context.Context, title, text string) (string, error) {
	baseURL := "https://scrapbox.io"
	pageURL := fmt.Sprintf("%s/%s/%s", baseURL, c.projectName, url.PathEscape(title))
	if text != "" {
		pageURL = fmt.Sprintf("%s?body=%s", pageURL, url.QueryEscape(text))
	}
	return pageURL, nil
}
