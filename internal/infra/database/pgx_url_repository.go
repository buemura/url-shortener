package database

import (
	"context"

	"github.com/buemura/url-shortener/internal/core/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxUrlRepository struct {
	db *pgxpool.Pool
}

func NewPgxUrlRepository() *PgxUrlRepository {
	return &PgxUrlRepository{
		db: Conn,
	}
}

func (r *PgxUrlRepository) Create(url *entity.Url) (*entity.Url, error) {
	_, err := r.db.Query(
		context.Background(),
		`INSERT INTO "url" (id, original_url, code, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`,
		url.ID, url.OriginalUrl, url.Code, url.CreatedAt, url.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return url, nil
}
