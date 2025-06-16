package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/duixe/social_app/internal/models"
)

// 👇 i.e like defining a class public property that can be accessed even outide the class
var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exist")
	QueryTimeOutDuration = time.Second * 5
)

type Repository struct {
	Posts interface {
		GetByID(context.Context, int64) (*models.Post, error)
		Create(context.Context, *models.Post) error
		Delete(context.Context, int64) error
		Update(context.Context, *models.Post) error
		GetUserFeed(context.Context, int64, PaginatedFeedQuery) ([]models.PostWithMetadata, error)
	}
	Users interface {
		GetByID(context.Context, int64) (*models.User, error)
		Create(context.Context, *sql.Tx, *models.User) error
		CreateAndInvite(ctx context.Context, user *models.User, token string, inviationExp time.Duration) error
		Activate(context.Context, string) error
	}
	Comments interface {
		Create(context.Context, *models.Comment) error
		GetByPostID(context.Context, int64) ([]models.Comment, error)
	}
	Followers interface {
		Follow(ctx context.Context, followerID int64, userID int64) error
		UnFollow(ctx context.Context, followerID int64, userID int64) error
	}
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		Posts:     &PostRepository{db},
		Users:     &UsersRepository{db},
		Comments:  &CommentRepository{db},
		Followers: &FollowerRepository{db},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}


	return tx.Commit()
}
