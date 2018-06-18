package models

type User struct {
	Id    uint64
	Login string
	Pass  string
	Email string
	Token string
}

func NewUser(id uint64, login, pass, email, token string) *User {
	return &User{id, login, pass, email, token}
}
