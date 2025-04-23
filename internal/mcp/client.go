package scrapbox

import (
	"bytes"
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
	Text  string `json:"text"`
}

// PageList represents a list of Scrapbox pages.
type PageList struct {
	Pages []Page `json:"pages"`
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
	log.Printf("GET %s", endpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, errors.NewScrapboxError(errors.ErrServerError, "failed to create request", err)
	}
	req.Header.Set("Cookie", fmt.Sprintf("connect.sid=%s", c.cookie))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.NewScrapboxError(errors.ErrServerError, "failed to send request", err)
	}
	defer resp.Body.Close()

	log.Printf("Response status: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewScrapboxError(resp.StatusCode, "unexpected status code", nil)
	}

	var page Page
	if err := json.NewDecoder(resp.Body).Decode(&page); err != nil {
		return nil, errors.NewScrapboxError(errors.ErrServerError, "failed to decode response", err)
	}

	return &page, nil
}

// ListPages retrieves a list of pages.
func (c *Client) ListPages(ctx context.Context) (*PageList, error) {
	endpoint := fmt.Sprintf("%s/pages/%s", c.baseURL, c.projectName)
	log.Printf("GET %s", endpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, errors.NewScrapboxError(errors.ErrServerError, "failed to create request", err)
	}
	req.Header.Set("Cookie", fmt.Sprintf("connect.sid=%s", c.cookie))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.NewScrapboxError(errors.ErrServerError, "failed to send request", err)
	}
	defer resp.Body.Close()

	log.Printf("Response status: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewScrapboxError(resp.StatusCode, "unexpected status code", nil)
	}

	var pageList PageList
	if err := json.NewDecoder(resp.Body).Decode(&pageList); err != nil {
		return nil, errors.NewScrapboxError(errors.ErrServerError, "failed to decode response", err)
	}

	return &pageList, nil
}

// SearchPages searches pages by query.
func (c *Client) SearchPages(ctx context.Context, query string) (*PageList, error) {
	endpoint := fmt.Sprintf("%s/pages/%s/search?q=%s", c.baseURL, c.projectName, url.QueryEscape(query))
	log.Printf("GET %s", endpoint)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, errors.NewScrapboxError(errors.ErrServerError, "failed to create request", err)
	}
	req.Header.Set("Cookie", fmt.Sprintf("connect.sid=%s", c.cookie))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.NewScrapboxError(errors.ErrServerError, "failed to send request", err)
	}
	defer resp.Body.Close()

	log.Printf("Response status: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewScrapboxError(resp.StatusCode, "unexpected status code", nil)
	}

	var pageList PageList
	if err := json.NewDecoder(resp.Body).Decode(&pageList); err != nil {
		return nil, errors.NewScrapboxError(errors.ErrServerError, "failed to decode response", err)
	}

	return &pageList, nil
}

// CreatePage creates a new page.
func (c *Client) CreatePage(ctx context.Context, title, text string) (*Page, error) {
	endpoint := fmt.Sprintf("%s/pages/%s", c.baseURL, c.projectName)
	log.Printf("POST %s", endpoint)
	body := map[string]string{
		"title": title,
		"text":  text,
	}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, errors.NewScrapboxError(errors.ErrServerError, "failed to marshal request body", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(bodyJSON))
	if err != nil {
		return nil, errors.NewScrapboxError(errors.ErrServerError, "failed to create request", err)
	}
	req.Header.Set("Cookie", fmt.Sprintf("connect.sid=%s", c.cookie))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.NewScrapboxError(errors.ErrServerError, "failed to send request", err)
	}
	defer resp.Body.Close()

	log.Printf("Response status: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusCreated {
		return nil, errors.NewScrapboxError(resp.StatusCode, "unexpected status code", nil)
	}

	var page Page
	if err := json.NewDecoder(resp.Body).Decode(&page); err != nil {
		return nil, errors.NewScrapboxError(errors.ErrServerError, "failed to decode response", err)
	}

	return &page, nil
}
