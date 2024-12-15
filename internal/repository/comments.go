package repository

import (
	"context"
	"database/sql"

	"github.com/duixe/social_app/internal/models"
)

type CommentRepository struct {
	db *sql.DB
}

func (s *CommentRepository) GetByPostID(ctx context.Context, postID int64) ([]models.Comment, error){
	query := `
		SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, users.first_name, users.id FROM comments c
		JOIN users on users.id = c.user_id
		WHERE c.post_id = $1
		ORDER BY c.created_at DESC
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeOutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	comments := []models.Comment{}
	for rows.Next() {
		var c models.Comment
		c.User = models.User{}
		err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt, &c.User.FirstName, &c.User.ID)
		if err != nil {
			return nil, err
		}

		comments = append(comments, c)
	}

	return comments, nil
}