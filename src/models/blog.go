package models

import "time"

type TagBlog struct {
	ID        string `json:"id"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	Variables string `json:"variables"`
}

type Blog struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Slug         string    `json:"slug"`
	SlugCategory string    `json:"slugCategory"`
	Tags         []TagBlog `pg:",array" json:"tags"`
	Thumbnail    string    `json:"thumbnail"`
	Status       bool      `json:"status"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type BlogBySlug struct {
	ContentMd    string `json:"contentMd"`
	ContentHtml  string `json:"contentHtml"`
	BlogsRelated []Blog `json:"blogsRelated"`
	Blog
}

type BlogCategory struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Thumbnail string `json:"thumbnail"`
	Blogs     []Blog `json:"blogs"`
}
