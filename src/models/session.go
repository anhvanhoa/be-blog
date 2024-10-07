package models

import "time"

type Session struct {
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	IP        string    `json:"ip"`
	ExpiredAt int64     `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
}
