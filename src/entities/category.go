package entities

import "time"

type Category struct {
	tableName struct{}  `pg:"categories"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Parent    string    `json:"parent"`
	Thumbnail string    `json:"thumbnail"`
	Order     int       `json:"order"`
	Status    *bool     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
