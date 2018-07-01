package models

type User struct {
	Id    	uint64
	Login	string
	Name	string
	Email	string
	Date	int64
	Rank	uint8
	Pass  	string
	Year	string
	Token 	string
	Allow_map	bool
}

func NewUser(
	id uint64,
	login,
	name,
	email string,
	date int64,
	rank uint8,
	pass,
	year,
	token string,
	allow_map bool,
	) *User {

	return &User{id, login, name, email, date, rank,pass, year, token, allow_map}

}
