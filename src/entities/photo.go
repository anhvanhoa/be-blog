package entities

import (
	"time"
)

type Photo struct {
	tableName struct{}  `pg:"photos"`
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	AuthorId  string    `json:"authorId"`
	PublicID  string    `json:"publicId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PhotoUpdate struct {
	tableName struct{}  `pg:"photos"`
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	AuthorId  string    `json:"author_id"`
	UpdatedAt time.Time `json:"updated_at"`
}
