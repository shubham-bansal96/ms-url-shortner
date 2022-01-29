package model

import (
	"net/http"
	"testing"

	"github.com/ms-url-shortner/app/config"
	"github.com/stretchr/testify/assert"
)

func TestNewURLDto(testRunner *testing.T) {
	testRunner.Run("NewURLDto can't be nil", func(testRunnerChild *testing.T) {
		urlDto := NewURLDto()
		assert.NotNil(testRunnerChild, urlDto, "NewURLDto can't be nil object")
	})
}

func TestValidate(testRunner *testing.T) {
	// Test case 1
	testRunner.Run("Validate should return url is empty error", func(testRunnerChild *testing.T) {
		config.Config = &config.Configuration{
			MSName:      "test-ms-url-shortnet",
			Environment: "test",
		}
		urlDto := NewURLDto()
		urlDto.URL = nil
		respObj := urlDto.Validate()
		assert.NotNil(testRunnerChild, respObj, "Validate must return error object")
		assert.Equal(testRunnerChild, http.StatusUnprocessableEntity, *respObj.ErrorCode, "Validate must return url is empty error with 422 error code")
	})
	// Test case 2
	testRunner.Run("Validate should return url is invalid error", func(testRunnerChild *testing.T) {
		config.Config = &config.Configuration{
			MSName:      "test-ms-url-shortnet",
			Environment: "test",
		}
		urlDto := NewURLDto()
		url := "infracloud.io"
		urlDto.URL = &url
		respObj := urlDto.Validate()
		assert.NotNil(testRunnerChild, respObj, "Validate must return error object")
		assert.Equal(testRunnerChild, http.StatusUnprocessableEntity, *respObj.ErrorCode, "Validate must return url is invalid error with 422 error code")
	})
	// Test case 3
	testRunner.Run("Validate should not return any error", func(testRunnerChild *testing.T) {
		config.Config = &config.Configuration{
			MSName:      "test-ms-url-shortnet",
			Environment: "test",
		}
		urlDto := NewURLDto()
		url := "https://www.infracloud.io"
		urlDto.URL = &url
		respObj := urlDto.Validate()
		assert.Nil(testRunnerChild, respObj, "Validate should not return any error if url is valid")
	})

	// Test case 4
	testRunner.Run("Validate should return url is invalid error if http or https are not found as prefix", func(testRunnerChild *testing.T) {
		config.Config = &config.Configuration{
			MSName:      "test-ms-url-shortnet",
			Environment: "test",
		}
		urlDto := NewURLDto()
		url := "www://infracloud.io"
		urlDto.URL = &url
		respObj := urlDto.Validate()
		assert.NotNil(testRunnerChild, respObj, "Validate must return error object")
		assert.Equal(testRunnerChild, http.StatusUnprocessableEntity, *respObj.ErrorCode, "Validate must return url is invalid error with 422 error code if http or https are not found as prefix")
	})
}
