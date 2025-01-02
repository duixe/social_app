package repository

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type FollowerRepository struct {
	db *sql.DB
}

func (s *FollowerRepository) Follow(ctx context.Context, followerID int64, userID int64) error {
	query := `
		INSERT INTO followers (user_id, follower_id) VALUES ($1, $2)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {
			return ErrConflict
		}
	}
	return err
}

func (s *FollowerRepository) UnFollow(ctx context.Context, followerID int64, userID int64) error {
	query := `
		DELETE FROM followers 
		WHERE user_id = $1 AND follower_id = $2
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()
	
	_, err := s.db.ExecContext(ctx, query, userID, followerID)
	return err
}

