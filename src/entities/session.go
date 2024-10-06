package entities

type Session struct {
	tableName struct{} `pg:"sessions"`
	ID        string   `json:"id"`
	UserID    string   `json:"user_id"`
	Token     string   `json:"token"`
	ExpiredAt int64    `json:"expired_at"`
	CreatedAt int64    `json:"created_at"`
}
