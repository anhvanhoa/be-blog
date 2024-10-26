package models

import "time"

type CommentReq struct {
	ID               int    `json:"id"`
	Content          string `json:"content" valid:"required"`
	UsernameReceiver string `json:"usernameReceiver"`
	AuthorId         string `json:"authorId" valid:"required"`
	BlogId           string `json:"blogId" valid:"required"`
	ParentCommentId  string `json:"parentCommentId"`
	Status           bool   `json:"status"`
}

type Comment struct {
	ID              string    `json:"id"`
	Username        string    `json:"username"`
	Avatar          string    `json:"avatar"`
	FullName        string    `json:"fullName"`
	Content         string    `json:"content" valid:"required"`
	BlogId          string    `json:"blogId" valid:"required"`
	ParentCommentId string    `json:"parentCommentId"`
	CreatedAt       time.Time `json:"createdAt"`
	Replies         []Comment `json:"replies"`
}

type AuthorComment struct {
	ID       string `json:"id"`
	Avatar   string `json:"avatar"`
	FullName string `json:"fullName"`
	Username string `json:"username"`
}
