package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/buemura/url-shortener/config"
	"github.com/buemura/url-shortener/internal/core/entity"
	"github.com/buemura/url-shortener/internal/core/gateway"
)

type GetShortenedUrl struct {
	cacheStorage  gateway.CacheStorage
	urlRepository gateway.UrlRepository
}

func NewGetShortenedUrl(cacheStorage gateway.CacheStorage, urlRepository gateway.UrlRepository) *GetShortenedUrl {
	return &GetShortenedUrl{cacheStorage: cacheStorage, urlRepository: urlRepository}
}

func (u *GetShortenedUrl) Execute(code string) (*entity.Url, error) {
	// Get url from cache
	urlCache, err := u.cacheStorage.Get(fmt.Sprintf("%s:%s", config.CACHE_URL_KEY_PREFIX, code))
	if err != nil {
		return nil, err
	}
	if len(urlCache) > 0 {
		return parseCachedUrl(urlCache)
	}

	// Get url from database
	urlDb, err := u.urlRepository.FindByCode(code)
	if err != nil {
		return nil, err
	}

	return urlDb, nil
}

func parseCachedUrl(urlCache string) (*entity.Url, error) {
	var url *entity.Url
	err := json.Unmarshal([]byte(urlCache), &url)
	if err != nil {
		return nil, err
	}
	return url, nil
}
