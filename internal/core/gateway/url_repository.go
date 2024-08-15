package gateway

import "github.com/buemura/url-shortener/internal/core/entity"

type UrlRepository interface {
	FindByCode(code string) (*entity.Url, error)
	Create(url *entity.Url) (*entity.Url, error)
}
