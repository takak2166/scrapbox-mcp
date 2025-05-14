package scrapbox

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/takak2166/scrapbox-mcp/internal/errors"
)

func TestClient_GetPage(t *testing.T) {
	tests := map[string]struct {
		statusCode int
		response   any
		title      string
		expectPage *Page
		expectErr  error
	}{
		"ok: success": {
			statusCode: http.StatusOK,
			response: Page{
				Title: "TestTitle",
				Lines: []Line{{Text: "line1"}},
			},
			title: "TestTitle",
			expectPage: &Page{
				Title: "TestTitle",
				Lines: []Line{{Text: "line1"}},
			},
			expectErr: nil,
		},
		"ng: not found": {
			statusCode: http.StatusNotFound,
			response:   map[string]string{"error": "not found"},
			title:      "NotFound",
			expectPage: nil,
			expectErr:  &errors.ScrapboxError{Code: http.StatusNotFound, Message: "unexpected status code", Err: nil},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.statusCode)
				if tc.response != nil {
					_ = json.NewEncoder(w).Encode(tc.response)
				}
			}))
			t.Cleanup(ts.Close)

			client := &Client{
				httpClient:  ts.Client(),
				baseURL:     ts.URL,
				projectName: "testproject",
				cookie:      "dummy",
			}

			page, err := client.GetPage(context.Background(), tc.title)
			if diff := cmp.Diff(tc.expectPage, page); diff != "" {
				t.Errorf("GetPage() page mismatch (-want +got):\n%s", diff)
			}
			if tc.expectErr == nil && err != nil {
				t.Errorf("GetPage() unexpected error: %v", err)
				return
			}
			if tc.expectErr != nil && err == nil {
				t.Errorf("GetPage() expected error but got nil")
				return
			}
			if tc.expectErr != nil {
				e1, ok1 := err.(*errors.ScrapboxError)
				e2, ok2 := tc.expectErr.(*errors.ScrapboxError)
				if !ok1 || !ok2 || !reflect.DeepEqual(e1.Code, e2.Code) || !reflect.DeepEqual(e1.Message, e2.Message) {
					t.Errorf("GetPage() error mismatch: got=%v, want=%v", err, tc.expectErr)
				}
			}
		})
	}
}

func TestClient_ListPages(t *testing.T) {
	tests := map[string]struct {
		statusCode int
		response   any
		expectList *PageList
		expectErr  error
	}{
		"ok: success": {
			statusCode: http.StatusOK,
			response: PageList{
				Pages: []Page{{Title: "A"}, {Title: "B"}},
			},
			expectList: &PageList{Pages: []Page{{Title: "A"}, {Title: "B"}}},
			expectErr:  nil,
		},
		"ng: unexpected status": {
			statusCode: http.StatusNotFound,
			response:   map[string]string{"error": "not found"},
			expectList: nil,
			expectErr:  &errors.ScrapboxError{Code: http.StatusNotFound, Message: "unexpected status code", Err: nil},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.statusCode)
				if tc.response != nil {
					_ = json.NewEncoder(w).Encode(tc.response)
				}
			}))
			t.Cleanup(ts.Close)
			client := &Client{
				httpClient:  ts.Client(),
				baseURL:     ts.URL,
				projectName: "testproject",
				cookie:      "dummy",
			}
			list, err := client.ListPages(context.Background())
			if diff := cmp.Diff(tc.expectList, list); diff != "" {
				t.Errorf("ListPages() mismatch (-want +got):\n%s", diff)
			}
			if tc.expectErr == nil && err != nil {
				t.Errorf("ListPages() unexpected error: %v", err)
				return
			}
			if tc.expectErr != nil && err == nil {
				t.Errorf("ListPages() expected error but got nil")
				return
			}
			if tc.expectErr != nil {
				e1, ok1 := err.(*errors.ScrapboxError)
				e2, ok2 := tc.expectErr.(*errors.ScrapboxError)
				if !ok1 || !ok2 || !reflect.DeepEqual(e1.Code, e2.Code) || !reflect.DeepEqual(e1.Message, e2.Message) {
					t.Errorf("ListPages() error mismatch: got=%v, want=%v", err, tc.expectErr)
				}
			}
		})
	}
}

func TestClient_SearchPages(t *testing.T) {
	tests := map[string]struct {
		statusCode int
		response   any
		query      string
		expectList *PageList
		expectErr  error
	}{
		"ok: success": {
			statusCode: http.StatusOK,
			response: PageList{
				Pages: []Page{{Title: "Q1"}},
			},
			query:      "Q1",
			expectList: &PageList{Pages: []Page{{Title: "Q1"}}},
			expectErr:  nil,
		},
		"ng: unexpected status": {
			statusCode: http.StatusNotFound,
			response:   map[string]string{"error": "not found"},
			query:      "none",
			expectList: nil,
			expectErr:  &errors.ScrapboxError{Code: http.StatusNotFound, Message: "unexpected status code", Err: nil},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.statusCode)
				if tc.response != nil {
					_ = json.NewEncoder(w).Encode(tc.response)
				}
			}))
			t.Cleanup(ts.Close)
			client := &Client{
				httpClient:  ts.Client(),
				baseURL:     ts.URL,
				projectName: "testproject",
				cookie:      "dummy",
			}
			list, err := client.SearchPages(context.Background(), tc.query)
			if diff := cmp.Diff(tc.expectList, list); diff != "" {
				t.Errorf("SearchPages() mismatch (-want +got):\n%s", diff)
			}
			if tc.expectErr == nil && err != nil {
				t.Errorf("SearchPages() unexpected error: %v", err)
				return
			}
			if tc.expectErr != nil && err == nil {
				t.Errorf("SearchPages() expected error but got nil")
				return
			}
			if tc.expectErr != nil {
				e1, ok1 := err.(*errors.ScrapboxError)
				e2, ok2 := tc.expectErr.(*errors.ScrapboxError)
				if !ok1 || !ok2 || !reflect.DeepEqual(e1.Code, e2.Code) || !reflect.DeepEqual(e1.Message, e2.Message) {
					t.Errorf("SearchPages() error mismatch: got=%v, want=%v", err, tc.expectErr)
				}
			}
		})
	}
}

func TestClient_CreatePage(t *testing.T) {
	tests := map[string]struct {
		statusCode int
		response   any
		title      string
		text       string
		expectPage *Page
		expectErr  error
	}{
		"ok: created": {
			statusCode: http.StatusCreated,
			response: Page{
				Title: "NewPage",
				Lines: []Line{{Text: "new line"}},
			},
			title: "NewPage",
			text:  "new line",
			expectPage: &Page{
				Title: "NewPage",
				Lines: []Line{{Text: "new line"}},
			},
			expectErr: nil,
		},
		"ng: unexpected status": {
			statusCode: http.StatusBadRequest,
			response:   map[string]string{"error": "bad request"},
			title:      "Bad",
			text:       "bad",
			expectPage: nil,
			expectErr:  &errors.ScrapboxError{Code: http.StatusBadRequest, Message: "unexpected status code", Err: nil},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.statusCode)
				if tc.response != nil {
					_ = json.NewEncoder(w).Encode(tc.response)
				}
			}))
			t.Cleanup(ts.Close)
			client := &Client{
				httpClient:  ts.Client(),
				baseURL:     ts.URL,
				projectName: "testproject",
				cookie:      "dummy",
			}
			page, err := client.CreatePage(context.Background(), tc.title, tc.text)
			if diff := cmp.Diff(tc.expectPage, page); diff != "" {
				t.Errorf("CreatePage() mismatch (-want +got):\n%s", diff)
			}
			if tc.expectErr == nil && err != nil {
				t.Errorf("CreatePage() unexpected error: %v", err)
				return
			}
			if tc.expectErr != nil && err == nil {
				t.Errorf("CreatePage() expected error but got nil")
				return
			}
			if tc.expectErr != nil {
				e1, ok1 := err.(*errors.ScrapboxError)
				e2, ok2 := tc.expectErr.(*errors.ScrapboxError)
				if !ok1 || !ok2 || !reflect.DeepEqual(e1.Code, e2.Code) || !reflect.DeepEqual(e1.Message, e2.Message) {
					t.Errorf("CreatePage() error mismatch: got=%v, want=%v", err, tc.expectErr)
				}
			}
		})
	}
}
