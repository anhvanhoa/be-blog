package jwt

type PayloadUser struct {
	ID       string `json:"id"`
	Avatar   string `json:"avatar"`
	FullName string `json:"fullName"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Roles    string `json:"roles"`
	Exp      int64
}

type PayloadVerify struct {
	Email string
	Exp   int64
}

type PayloadResetPass struct {
	Email string
	Exp   int64
}
