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

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

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

func (s *UsersRepository) GetByID(ctx context.Context, userID int64) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, username, email, password, created_at
		FROM users
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	user := &models.User{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}