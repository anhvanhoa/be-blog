package models

type Category struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Parent    string `json:"parent"`
	Status    bool   `json:"status"`
	Slug      string `json:"slug"`
	OrderC    int    `pg:"order_c" json:"order_c"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
