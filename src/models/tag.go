package models

type TagReq struct {
	ID        string `json:"id"`
	Name      string `json:"name" valid:"required~Tên không được để trống"`
	Slug      string `json:"slug" valid:"required~Slug không được để trống"`
	Variables string `json:"variables" valid:"required~Màu không được để trống"`
}
