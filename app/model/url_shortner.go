package model

import (
	"net/http"
	"strings"

	"github.com/ms-url-shortner/app/logging"
)

type URLDTO struct {
	URL *string `json:"url"`
}

func (url *URLDTO) Validate() *Error {
	lw := logging.LogForFunc()

	if url.URL == nil || *url.URL == "" {
		lw.Warn("url is empty")
		return NewError(http.StatusUnprocessableEntity, "url is empty")
	}

	str := strings.Split(*url.URL, "//")
	if len(str) < 0 {
		lw.Error("url is invalid")
		return NewError(http.StatusUnprocessableEntity, "invalid url")
	}

	if str[0] == "https:" || str[0] == "http:" {
		lw.Error("url is valid")
		return nil
	}

	lw.Error("url is invalid")
	return NewError(http.StatusUnprocessableEntity, "invalid url")
}
