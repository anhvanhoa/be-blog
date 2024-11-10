package entities

import (
	"time"
)

type User struct {
	tableName struct{}  `pg:"users"`
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Avatar    string    `json:"avatar"`
	FullName  string    `json:"fullName"`
	Bio       string    `json:"bio"`
	Gender    string    `json:"gender"`
	Status    string    `json:"status"`
	Roles     []string  `pg:",array" json:"roles"`
	Birthday  time.Time `json:"birthday"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserUpdate struct {
	tableName struct{}  `pg:"users"`
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FullName  string    `json:"fullName"`
	Bio       string    `json:"bio"`
	Gender    string    `json:"gender"`
	Birthday  time.Time `json:"birthday"`
	UpdatedAt time.Time `json:"updated_at"`
}
