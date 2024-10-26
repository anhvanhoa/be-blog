package entities

import (
	"time"
)

type Comment struct {
	tableName       struct{}  `pg:"comments"`
	ID              string    `json:"id"`
	Content         string    `json:"content"`
	AuthorId        string    `json:"authorId"`
	BlogId          string    `json:"blogId"`
	ParentCommentId string    `json:"parent_comment_id"`
	Status          *bool     `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}
