package testhelper

import (
	"context"

	"github.com/ms-url-shortner/app/model"
)

type MockURLShortnerServie struct {
	data *model.URLDTO
	// UIDService *MockUIDService
}

func (uss *MockURLShortnerServie) ShortURL(ctx context.Context, url string) *model.URLDTO {
	return uss.data
}

func (uss *MockURLShortnerServie) SetData(shortUrl *model.URLDTO) {
	uss.data = shortUrl
}

type MockUIDService struct {
}

func (uis *MockUIDService) GetUniqueID() string {
	return "test123"
}
