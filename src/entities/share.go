package entities

import (
	"time"
)

type Share struct {
	tableName struct{}  `pg:"shares"`
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Url       string    `json:"url"`
	Author    string    `json:"author"`
	Status    *bool     `json:"status"`
	View      int       `json:"view"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
