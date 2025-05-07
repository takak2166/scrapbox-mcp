package errors

import (
	"fmt"
)

type ScrapboxError struct {
	Code    int
	Message string
	Err     error
}

func (e *ScrapboxError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("scrapbox error: %s (code: %d): %v", e.Message, e.Code, e.Err)
	}
	return fmt.Sprintf("scrapbox error: %s (code: %d)", e.Message, e.Code)
}

const (
	ErrInvalidCredentials = 401
	ErrNotFound           = 404
	ErrRateLimit          = 429
	ErrServerError        = 500
)

func IsScrapboxError(err error) bool {
	_, ok := err.(*ScrapboxError)
	return ok
}
