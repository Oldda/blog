package model

import(
	
)

type User struct {
	Id int
	Nickname string
	Realname string
	Avatar string
	Email string
	Password string
	Phone string
	CreatedAt int
	UpdatedAt int
	DeletedAt int
}

func (u *User)TableName()string{
	return "users"
}
