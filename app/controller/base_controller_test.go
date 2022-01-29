package controller

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ms-url-shortner/app/model"
	testhelper "github.com/ms-url-shortner/app/test-helper"
	"github.com/stretchr/testify/assert"
)

func TestNewBaseController(testRunner *testing.T) {
	testRunner.Run("NewBaseController service can't be nil", func(testRunnerChild *testing.T) {
		baseCtrl := NewBaseContoller(&testhelper.MockURLShortnerServie{})
		assert.NotNil(testRunnerChild, baseCtrl.ShortenURLService, "baseCtrl.ShortenURLService can't be nil object")
	})
}

func TestPing(testRunner *testing.T) {
	testRunner.Run("ping should return status code 200", func(testRunnerChild *testing.T) {
		baseCtrl := NewBaseContoller(&testhelper.MockURLShortnerServie{})
		handler, httpRequest, responseRecorder := createGetHandleSetup(testRunnerChild, "/ping", baseCtrl.Ping)
		handler.ServeHTTP(responseRecorder, httpRequest)
		assert.Equal(testRunnerChild, http.StatusOK, responseRecorder.Code, "ping response must be 200")
	})
}

func createGetHandleSetup(testRunner *testing.T, endPoint string, handler gin.HandlerFunc) (*gin.Engine, *http.Request, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	ginEngine := gin.Default()

	ginEngine.GET(endPoint, handler)

	httpRequest, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		testRunner.Fatal(err)
	}
	responseRecorder := httptest.NewRecorder()

	return ginEngine, httpRequest, responseRecorder
}

func createPostHandleSetup(testRunner *testing.T, endPoint string, handler gin.HandlerFunc, body io.Reader) (*gin.Engine, *http.Request, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	ginEngine := gin.Default()

	ginEngine.POST(endPoint, handler)

	httpRequest, err := http.NewRequest("POST", endPoint, body)
	if err != nil {
		testRunner.Fatal(err)
	}
	requestRecorder := httptest.NewRecorder()

	return ginEngine, httpRequest, requestRecorder
}
func TestHandleURLShortner(testRunner *testing.T) {
	//Test case 1
	testRunner.Run("HandleURLShortner broken json", func(testRunnerChild *testing.T) {
		mockURLShortService := &testhelper.MockURLShortnerServie{}
		testhelper.MockConfig()
		baseCtrl := NewBaseContoller(mockURLShortService)

		handler, httpRequest, responseRecorder := createPostHandleSetup(testRunnerChild, "/getshorturl", baseCtrl.HandleURLShortner, bytes.NewBuffer([]byte(`{"url":10}`)))
		handler.ServeHTTP(responseRecorder, httpRequest)
		assert.Equal(testRunnerChild, http.StatusBadRequest, responseRecorder.Code, "HandleURLShortner must return error code 400")
	})

	//Test case 2
	testRunner.Run("HandleURLShortner empty url", func(testRunnerChild *testing.T) {
		mockURLShortService := &testhelper.MockURLShortnerServie{}
		testhelper.MockConfig()
		baseCtrl := NewBaseContoller(mockURLShortService)

		handler, httpRequest, responseRecorder := createPostHandleSetup(testRunnerChild, "/getshorturl", baseCtrl.HandleURLShortner, bytes.NewBuffer([]byte(`{"url":""}`)))
		handler.ServeHTTP(responseRecorder, httpRequest)
		assert.Equal(testRunnerChild, http.StatusUnprocessableEntity, responseRecorder.Code, "HandleURLShortner must return error code 422 for empty url")
	})

	//Test case 3
	testRunner.Run("HandleURLShortner invalid url", func(testRunnerChild *testing.T) {
		mockURLShortService := &testhelper.MockURLShortnerServie{}
		testhelper.MockConfig()
		baseCtrl := NewBaseContoller(mockURLShortService)

		handler, httpRequest, responseRecorder := createPostHandleSetup(testRunnerChild, "/getshorturl", baseCtrl.HandleURLShortner, bytes.NewBuffer([]byte(`{"url":"test.url.com"}`)))
		handler.ServeHTTP(responseRecorder, httpRequest)
		assert.Equal(testRunnerChild, http.StatusUnprocessableEntity, responseRecorder.Code, "HandleURLShortner must return error code 422 for invalid url")
	})

	//Test case 4
	testRunner.Run("HandleURLShortner should return success", func(testRunnerChild *testing.T) {
		mockURLShortService := &testhelper.MockURLShortnerServie{}
		testhelper.MockConfig()
		baseCtrl := NewBaseContoller(mockURLShortService)

		shortURL := "https://shorturl.com/ea95sh23"
		mockURLShortService.SetData(&model.URLDTO{URL: &shortURL})

		handler, httpRequest, responseRecorder := createPostHandleSetup(testRunnerChild, "/getshorturl", baseCtrl.HandleURLShortner, bytes.NewBuffer([]byte(`{"url":"https://www.infracloud.io"}`)))
		handler.ServeHTTP(responseRecorder, httpRequest)
		assert.Equal(testRunnerChild, http.StatusOK, responseRecorder.Code, "HandleURLShortner must return statusok code 200 for valid url")
	})
}
