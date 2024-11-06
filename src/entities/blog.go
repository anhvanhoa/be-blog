package entities

import "time"

type Blog struct {
	tableName   struct{} `pg:"blogs"`
	ID          string
	Title       string
	ContentMd   string
	ContentHtml string
	Description string
	AuthorId    string
	CategoryId  string
	Views       int
	Slug        string
	Tags        []string `pg:",array"`
	Thumbnail   string
	Status      *bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
type BlogUpdate struct {
	tableName   struct{} `pg:"blogs"`
	ID          string
	Title       string
	ContentMd   string
	ContentHtml string
	Description string
	AuthorId    string
	CategoryId  string
	Views       int
	Slug        string
	Tags        []string `pg:",array"`
	Thumbnail   string
	Status      *bool
	UpdatedAt   time.Time
}

type BlogComment struct {
	tableName struct{} `pg:"blogs"`
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Email     string   `json:"email"`
}
