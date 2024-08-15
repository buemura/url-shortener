package domain

import "time"

type Url struct {
	ID          string
	OriginalUrl string
	ShortUrl    string
	Code        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
