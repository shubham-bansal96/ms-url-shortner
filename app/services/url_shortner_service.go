package services

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/ms-url-shortner/app/logging"
	"github.com/ms-url-shortner/app/model"
)

const shortURL = "shorturl.com/"

var (
	urlRepo      map[string]string = make(map[string]string)
	shortURLRepo map[string]string = make(map[string]string)
)

type IShortenUrl interface {
	ShortURL(ctx context.Context, url string) *model.URLDTO
}

type ShortenUrl struct {
	UId IUID
}

func NewShortenURLService(uid IUID) IShortenUrl {
	return &ShortenUrl{
		UId: uid,
	}
}

// ShortURL = > short the URL coming in JSON request
func (su *ShortenUrl) ShortURL(ctx context.Context, url string) *model.URLDTO {
	lw := logging.LogForFunc()

	if shortURL, isExists := urlRepo[url]; isExists {
		lw.Info("short url is already exists")
		URLShortedByCache.WithLabelValues(url).Inc()
		return &model.URLDTO{URL: &shortURL}
	}

	if longURL, isExists := shortURLRepo[url]; isExists {
		lw.Info("url is already exists")
		URLShortedByCache.WithLabelValues(url).Inc()
		return &model.URLDTO{URL: &longURL}
	}

	if len(url) <= 20 {
		lw.Info("url is already shorted")
		return &model.URLDTO{URL: &url}
	}

	str := strings.Split(url, "//")
	newURL := fmt.Sprintf("%s//%s%s", str[0], shortURL, su.UId.GetUniqueID())

	URLShortedBySVC.WithLabelValues(url).Inc()
	urlRepo[url] = newURL
	shortURLRepo[newURL] = url
	lw.Info("url shorted successfully")
	return &model.URLDTO{URL: &newURL}
}

// helper method for test case
func SetDataInURLRepository(key, value string) {
	urlRepo[key] = value
}

type IUID interface {
	GetUniqueID() string
}
type UId struct{}

func NewUidService() IUID {
	return &UId{}
}

// GetUniqueID => generate unique id in 32 hexa decimal digit in the form of 8-4-4-4-12
func (uid *UId) GetUniqueID() string {
	uuid := strings.Split(uuid.New().String(), "-")
	return uuid[0]
}
