package models

type Follower struct {
	UserID int64 `json:"user_id"`
	FollowerID int64 `json:"follower_id"`
	CreatedAt int64 `json:"created_at"`
}