package repository

import (
	"context"
	"database/sql"

	"github.com/duixe/social_app/internal/models"
)

type Repository struct {
	Posts interface {
		Create(context.Context, *models.Post) error
	}
	Users interface {
		Create(context.Context, *models.User) error
	}
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Posts: &PostsRepository{db},
		Users: &UsersRepository{db},
	}
}