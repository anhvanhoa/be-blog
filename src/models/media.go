package models

type PhotoReq struct {
	Id       string `json:"id" valid:"required~Id không được để trống"`
	Title    string `json:"title" valid:"required~Tên ảnh không được để trống"`
	AuthorId string `json:"author_id"`
}
