package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/duixe/social_app/internal/models"
	"github.com/lib/pq"
)

type PostRepository struct {
	db *sql.DB
}

func (s *PostRepository) Create(ctx context.Context, post *models.Post) error {
	query := `
	INSERT INTO posts (content, title, user_id, tags) 
	VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserId,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostRepository) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	query := `
		SELECT id, user_id, title, content, created_at, updated_at, tags, version
		FROM posts 
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	var post models.Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.UserId,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
		pq.Array(&post.Tags),
		&post.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}

func (s *PostRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM posts WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *PostRepository) Update(ctx context.Context, post *models.Post) error {
	query := `
		UPDATE posts
		SET title = $1, content = $2, version = version + 1
		WHERE id = $3 AND version = $4 
		RETURNING version
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Title,
		post.Content,
		post.ID,
		post.Version,
	).Scan(&post.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrNotFound
		default:
			return err
		}
	}

	return nil
}

func (s *PostRepository) GetUserFeed(ctx context.Context, userID int64, fq PaginatedFeedQuery) ([]models.PostWithMetadata, error) {
	query := `
		SELECT
			p.id,
			p.user_id,
			p.title,
			p.content,
			p.tags,
			u.email,
			p.version,
			p.created_at,
			COUNT(c.id) AS comments_count
		FROM
			posts p
			LEFT JOIN comments c ON c.post_id = p.id
			LEFT JOIN users u on u.id = p.user_id
			JOIN followers f ON f.follower_id = p.user_id OR p.user_id = $1
			WHERE f.user_id = $1 OR p.user_id = $1
		GROUP BY
			p.id, u.email
		ORDER BY
			p.created_at ` + fq.Sort +`
		LIMIT $2 OFFSET $3
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, userID, fq.Limit, fq.Offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var feed []models.PostWithMetadata
	for rows.Next() {
		var p models.PostWithMetadata
		err := rows.Scan(
			&p.ID,
			&p.UserId,
			&p.Title,
			&p.Content,
			pq.Array(&p.Tags),
			&p.User.Email,
			&p.Version,
			&p.CreatedAt,
			&p.CommentsCount,
		)
		if err != nil {
			return nil, err
		}

		feed = append(feed, p)
	}

	return feed, nil
}
