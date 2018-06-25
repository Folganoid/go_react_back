package models

import "time"

type User struct {
	Id    	uint64
	Login	string
	Name	string
	Email	string
	Date	time.Time
	Rank	uint8
	Pass  	string
	Year	string
	Token 	string
}

func NewUser(
	id uint64,
	login,
	name,
	email string,
	date time.Time,
	rank uint8,
	pass,
	year,
	token string) *User {

	return &User{id, login, name, email, date, rank,pass, year, token}

}
