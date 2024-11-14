package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"

	"github.com/skyline-digital/api/internal/models"
)

type PostStore struct {
	db *sql.DB
}

func (s *PostStore) Create(ctx context.Context, post *models.Post) error {
	query := `
		INSERT INTO posts (title, body, user_id, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	if err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Body,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	); err != nil {
		return err
	}

	return nil
}
