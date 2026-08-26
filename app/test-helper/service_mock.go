package testhelper

import (
	"context"

	"github.com/ms-url-shortner/app/model"
)

type MockURLShortnerServie struct {
	data *model.URLDTO
	// UIDService *MockUIDService
}

// ShortURL returns the pre-configured mock data for testing
func (uss *MockURLShortnerServie) ShortURL(ctx context.Context, url string) *model.URLDTO {
	return uss.data
}

// SetData configures the mock response that ShortURL will return
func (uss *MockURLShortnerServie) SetData(shortUrl *model.URLDTO) {
	uss.data = shortUrl
}

type MockUIDService struct {
}

// GetUniqueID returns a fixed test ID for deterministic test results
func (uis *MockUIDService) GetUniqueID() string {
	return "test123"
}
