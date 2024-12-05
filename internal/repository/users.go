package repository

import (
	"context"
	"database/sql"

	"github.com/duixe/social_app/internal/models"
)

type UsersRepository struct {
	db *sql.DB
}

func (s *UsersRepository) Create (ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (first_name, last_name, username, password, email) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at
	`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}