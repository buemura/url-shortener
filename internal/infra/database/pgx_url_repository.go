package database

import (
	"context"
	"errors"

	"github.com/buemura/url-shortener/internal/core/entity"
	"github.com/jackc/pgx/v5"
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

func (r *PgxUrlRepository) FindByCode(code string) (*entity.Url, error) {
	rows, err := r.db.Query(
		context.Background(),
		`SELECT id, original_url, code, created_at, updated_at
		FROM "url" 
		WHERE code = $1`,
		code,
	)
	if err != nil {
		return nil, err
	}

	url, err := pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByPos[entity.Url])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return url, nil
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
