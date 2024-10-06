package entities

type StatusAuth string

type Auth struct {
	tableName struct{} `pg:"users"`
	ID        string
	FullName  string
	Username  string
	Email     string
	Password  string
	CodeMail  string
	Status    StatusAuth
}
