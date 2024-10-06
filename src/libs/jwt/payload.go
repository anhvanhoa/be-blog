package jwt

type PayloadUser struct {
	ID       string
	FullName string
	UserName string
	Email    string
	Roles    string
	Exp      int64
}

type PayloadVerify struct {
	Email string
	Exp   int64
}
