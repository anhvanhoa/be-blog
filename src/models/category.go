package models

import "time"

type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Thumbnail string    `json:"thumbnail"`
	Parent    string    `json:"parent"`
	Status    bool      `json:"status"`
	Slug      string    `json:"slug"`
	OrderC    int       `pg:"order_c" json:"orderC"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CategoryReq struct {
	ID        string `json:"id"`
	Name      string `json:"name" valid:"required~Tên danh mục không được để trống"`
	Thumbnail string `json:"thumbnail" valid:"required~Ảnh đại diện không được để trống"`
	Parent    string `json:"parent"`
	Slug      string `json:"slug" valid:"required~Slug không được để trống"`
	OrderC    int    `pg:"order_c" json:"orderC"`
}

type CategoryManager struct {
	Category
	CountBlog int `json:"countBlog"`
}
