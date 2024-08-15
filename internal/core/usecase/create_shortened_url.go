package usecase

import (
	"github.com/buemura/url-shortener/internal/core/entity"
	"github.com/buemura/url-shortener/internal/core/gateway"
)

type CreateShortenedUrl struct {
	urlRepo gateway.UrlRepository
}

func NewCreateShortenedUrl(urlRepo gateway.UrlRepository) *CreateShortenedUrl {
	return &CreateShortenedUrl{urlRepo: urlRepo}
}

func (u *CreateShortenedUrl) Execute(urlInput string) (*entity.Url, error) {
	url, err := entity.NewUrl(urlInput)
	if err != nil {
		return nil, err
	}

	// // Save url in database
	// u.urlRepo.Create(url)

	return url, nil
}
