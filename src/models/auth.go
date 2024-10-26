package models

// register
type RegisterReq struct {
	FullName        string `json:"fullName" valid:"required~Vui lòng nhập họ tên"`
	Username        string `json:"username" valid:"required~Vui lòng nhập tên đăng nhập"`
	Email           string `json:"email" valid:"email~Email không hợp lệ"`
	Password        string `json:"password" valid:"required~Vui lòng nhập mật khẩu"`
	ConfirmPassword string `json:"confirmPassword" valid:"required"`
}

type RegisterRes struct {
	ID         string `json:"id"`
	FullName   string `json:"fullName"`
	UserName   string `json:"userName"`
	Email      string `json:"email"`
	EmailToken string `json:"emailToken"`
}

// verify-email

type VerifyEmailReq struct {
	Code string `json:"code" valid:"required~Vui lòng nhập mã xác thực"`
}

// login

type LoginReq struct {
	Username string `json:"username" valid:"required~Vui lòng nhập tên đăng nhập"`
	Password string `json:"password" valid:"required~Vui lòng nhập mật khẩu"`
}

type LoginRes struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
