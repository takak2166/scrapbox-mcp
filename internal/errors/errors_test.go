package errors

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestScrapboxError(t *testing.T) {
	tests := map[string][]struct {
		code     int
		message  string
		err      error
		expected string
	}{
		"ok: basic error": {
			{
				code:     ErrNotFound,
				message:  "page not found",
				err:      nil,
				expected: "scrapbox error: page not found (code: 404)",
			},
		},
		"err: error with inner error": {
			{
				code:     ErrServerError,
				message:  "server error",
				err:      errors.New("connection timeout"),
				expected: "scrapbox error: server error (code: 500): connection timeout",
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			for _, tc := range tt {
				err := NewScrapboxError(tc.code, tc.message, tc.err)
				if diff := cmp.Diff(tc.expected, err.Error()); diff != "" {
					t.Errorf("Error() mismatch (-want +got):\n%s", diff)
				}
				if !IsScrapboxError(err) {
					t.Error("IsScrapboxError() returned false, want true")
				}
			}
		})
	}
}
