package gateway

import "github.com/buemura/url-shortener/internal/core/entity"

type UrlRepository interface {
	Create(url *entity.Url) (*entity.Url, error)
}
