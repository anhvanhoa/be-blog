package entities

import "time"

type Blog struct {
	tableName   struct{}  `pg:"blogs"`
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	ContentMd   string    `json:"content_md"`
	ContentHtml string    `json:"content_html"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Category    string    `json:"category"`
	Views       int       `json:"views"`
	Slug        string    `json:"slug"`
	Tags        []string  `json:"tags"`
	Thumbnail   string    `json:"thumbnail"`
	Status      *bool     `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
