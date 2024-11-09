package models

type Profile struct {
	ID       string `json:"id"`
	Avatar   string `json:"avatar" `
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender"`
}

type ProfileUpdateReq struct {
	ID       string `json:"id"`
	FullName string `json:"fullName" valid:"required~Họ tên không được để trống"`
	Username string `json:"username" valid:"required~Tên đăng nhập không được để trống"`
	Email    string `json:"email" valid:"required~Email không được để trống,email~Email không đúng định dạng"`
	Bio      string `json:"bio"`
	Birthday string `json:"birthday"`
	Gender   string `json:"gender" valid:"in(male|female)~Giới tính không đúng định dạng"`
}

type ChangePasswordReq struct {
	Password        string `json:"password" valid:"required~Mật khẩu không được để trống"`
	NewPassword     string `json:"newPassword" valid:"required~Mật khẩu mới không được để trống"`
	ConfirmPassword string `json:"confirmPassword" valid:"required~Xác nhận mật khẩu không được để trống"`
}
