package models

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserId    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	Comments  []Comment `json:"comments"`
	Version   int       `json:"version"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	User      User      `json:"user"`
}

//post composition -  it's like a computed property of the post struct
type PostWithMetadata struct {
	Post
	CommentsCount int `json:"comments_count"`
}
