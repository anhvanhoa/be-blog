package entities

import "time"

type Category struct {
	tableName struct{}  `pg:"categories"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Parent    string    `json:"parent"`
	Thumbnail string    `json:"thumbnail"`
	Slug      string    `json:"slug"`
	OrderC    int       `pg:"order_c" json:"order_c"`
	Status    *bool     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CategoryUpdate struct {
	tableName struct{}  `pg:"categories"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Parent    string    `json:"parent"`
	Thumbnail string    `json:"thumbnail"`
	Slug      string    `json:"slug"`
	OrderC    int       `pg:"order_c" json:"order_c"`
	UpdatedAt time.Time `json:"updated_at"`
}
