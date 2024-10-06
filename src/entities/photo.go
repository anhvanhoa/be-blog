package entities

import (
	"time"
)

type Photo struct {
	tableName struct{}  `pg:"photos"`
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Url       string    `json:"url"`
	Author    string    `json:"author"`
	Status    *bool     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
