package models

import "time"

type TagBlog struct {
	ID        string `json:"id"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	Variables string `json:"variables"`
}

type BlogReq struct {
	ID          string    `json:"id"`
	Title       string    `json:"title" valid:"required~Tên không được để trống"`
	ContentMd   string    `json:"contentMd" valid:"required~Nội dung không được để trống"`
	ContentHtml string    `json:"contentHtml" valid:"required~Nội dung html không được để trống"`
	Description string    `json:"description" valid:"required~Mô tả không được để trống"`
	AuthorId    string    `json:"authorId"`
	CategoryId  string    `json:"categoryId" valid:"required~Danh mục không được để trống"`
	Slug        string    `json:"slug" valid:"required~Slug không được để trống"`
	Tags        []TagBlog `json:"tags"`
	Thumbnail   string    `json:"thumbnail" valid:"required~Ảnh không được để trống"`
	Status      *bool     `json:"status"`
}

type Blog struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Slug         string    `json:"slug"`
	SlugCategory string    `json:"slugCategory"`
	Category     string    `json:"category"`
	Tags         []TagBlog `pg:",array" json:"tags"`
	Thumbnail    string    `json:"thumbnail"`
	Status       bool      `json:"status"`
	UpdatedAt    time.Time `json:"updatedAt"`
	CreatedAt    time.Time `json:"createdAt"`
}

type BlogBySlug struct {
	ContentMd    string `json:"contentMd"`
	ContentHtml  string `json:"contentHtml"`
	BlogsRelated []Blog `json:"blogsRelated"`
	Security     string `json:"security"`
	Blog
}

type BlogByID struct {
	ContentMd   string    `json:"contentMd"`
	ContentHtml string    `json:"contentHtml"`
	Security    string    `json:"security"`
	CategoryId  string    `json:"categoryId"`
	CreatedAt   time.Time `json:"createdAt"`
	Blog
}

type BlogCategory struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Thumbnail string `json:"thumbnail"`
	Blogs     []Blog `json:"blogs"`
}
