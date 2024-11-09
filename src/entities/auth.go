package entities

type StatusAuth string

type Auth struct {
	tableName struct{} `pg:"users"`
	ID        string
	Avatar    string
	FullName  string
	Username  string
	Email     string
	Password  string
	CodeMail  string
	Roles     []string `pg:",array"`
	Status    StatusAuth
}

type UserResetPass struct {
	tableName struct{} `pg:"users"`
	Email     string
	Password  string
}
