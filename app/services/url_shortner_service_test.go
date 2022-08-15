package services

import (
	"context"
	"testing"

	testhelper "github.com/ms-url-shortner/app/test-helper"
	"github.com/stretchr/testify/assert"
)

func TestNewShortenURLService(testRunner *testing.T) {
	testRunner.Run("NewShortenURLService can't be nil", func(testRunnerChild *testing.T) {
		mockURLShortService := NewShortenURLService(&testhelper.MockUIDService{})
		assert.NotNil(testRunnerChild, mockURLShortService, "NewShortenURLService can't be nil object")
	})
}

func TestShortURL(testRunner *testing.T) {
	mockURLShortService := NewShortenURLService(&testhelper.MockUIDService{})
	//Test case 1
	testRunner.Run("TestShortURL data already exists in url repo", func(testRunnerChild *testing.T) {
		testhelper.MockConfig()
		SetDataInURLRepository("https://www.infracloud.io", "https://tinyurl.com/ea95sh23")
		reponseDTO := mockURLShortService.ShortURL(context.Background(), "https://www.infracloud.io")
		assert.Equal(testRunnerChild, "https://tinyurl.com/ea95sh23", *reponseDTO.URL, "TestShortURL must return https://tinyurl.com/ea95sh23")
	})

	//Test case 2
	testRunner.Run("TestShortURL url is already shorted", func(testRunnerChild *testing.T) {
		testhelper.MockConfig()
		reponseDTO := mockURLShortService.ShortURL(context.Background(), "https://www.test.com")
		assert.Equal(testRunnerChild, "https://www.test.com", *reponseDTO.URL, "TestShortURL must return the same url https://www.test.com")
	})

	//Test case 3
	testRunner.Run("TestShortURLreturn shorted url", func(testRunnerChild *testing.T) {
		testhelper.MockConfig()
		reponseDTO := mockURLShortService.ShortURL(context.Background(), "https://www.infracloudtest.io")
		assert.Equal(testRunnerChild, "https://shorturl.com/test123", *reponseDTO.URL, "TestShortURL must return the shorted url")
	})
}

func TestNewUidService(testRunner *testing.T) {
	testRunner.Run("TestNewUidService can't be nil", func(testRunnerChild *testing.T) {
		uidService := NewUidService()
		assert.NotNil(testRunnerChild, uidService, "TestNewUidService can't be nil object")
	})
}

func TestGetUniqueID(testRunner *testing.T) {
	testRunner.Run("TestGetUniqueID must return random string", func(testRunnerChild *testing.T) {
		mockUIDService := NewUidService()
		uid := mockUIDService.GetUniqueID()
		assert.NotEmpty(testRunnerChild, uid, "TestGetUniqueID must return random string")
	})
}
