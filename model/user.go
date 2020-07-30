package model

import(
	"time"
)

type User struct {
    Id       int       `xorm:"not null pk autoincr INT(11)"`
    Username string    `xorm:"not null VARCHAR(32)"`
    Birthday time.Time `xorm:"DATE"`
    Sex      string    `xorm:"CHAR(1)"`
    Address  string    `xorm:"VARCHAR(256)"`
}

func (u *User)TableName()string{
	return "users"
}
