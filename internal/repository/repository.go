package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/duixe/social_app/internal/models"
)

//👇 i.e like defining a class public property that can be accessed even outide the class
var (
	ErrNotFound = errors.New("resource not found")
	QueryTimeOutDuration = time.Second * 5
)

type Repository struct {
	Posts interface {
		GetByID(context.Context, int64) (*models.Post, error)
		Create(context.Context, *models.Post) error
		Delete(context.Context, int64) error
		Update(context.Context, *models.Post) error
	}
	Users interface {
		Create(context.Context, *models.User) error
	}
	Comments interface {
		GetByPostID(context.Context, int64) ([]models.Comment, error)
	}
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Posts: &PostRepository{db},
		Users: &UsersRepository{db},
		Comments: &CommentRepository{db},
	}
}