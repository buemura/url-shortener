package entity

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/lucsky/cuid"
)

type Url struct {
	ID          string
	OriginalUrl string
	Code        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewUrl(originalUrl string) (*Url, error) {
	cuid, err := cuid.NewCrypto(rand.Reader)
	if err != nil {
		return nil, err
	}

	code, err := generateCode(5)
	if err != nil {
		return nil, err
	}

	return &Url{
		ID:          cuid,
		OriginalUrl: originalUrl,
		Code:        code,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func generateCode(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buffer)[:length], nil
}
