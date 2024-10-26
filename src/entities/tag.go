package entities

import (
	"time"
)

type Tag struct {
	tableName struct{}  `pg:"tags"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Count     int       `json:"count"`
	Status    *bool     `json:"status"`
	Variables string    `json:"variables"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type TagUpdate struct {
	tableName struct{}  `pg:"tags"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	Variables string    `json:"variables"`
	UpdatedAt time.Time `json:"updatedAt"`
}
