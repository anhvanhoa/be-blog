package entities

import (
	"time"
)

type Comment struct {
	tableName struct{}  `pg:"comments"`
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Status    *bool     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
