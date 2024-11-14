package models

type PayloadWs struct {
	Event     string `json:"event"`
	Data      string `json:"data"`
	NameSpace string `json:"namespace"`
}

type UserWs struct {
	ID       string    `json:"id"`
	UserName string `json:"username"`
	FullName string `json:"full_name"`
	Avatar   string `json:"avatar"`
}
