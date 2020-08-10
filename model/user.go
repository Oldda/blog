package model

import(
	"log"
	"time"
)

type User struct {
	Id int `xorm:"pk autoincr" json:"id"`
	Nickname string `json:"nickname"`
	Realname string `json:"realname"`
	Avatar string `json:"avatar"`
	Email string `json:"email"`
	Password string `json:"-"`
	Phone string `json:"phone"`
	CreatedAt int64 `xorm:"created" json:"created_at"`
	UpdatedAt int64 `xorm:"updated" json:"updated_at"`
	DeletedAt int64 `json:"-"`
}

func (u *User)TableName()string{
	return "users"
}

func(u *User)CreateUser(nickname,realname,password,phone,email,avatar string)(int,string){
	u.Nickname = nickname
	u.Realname = realname
	u.Password = password
	u.Phone = phone
	u.Email = email
	u.Avatar = avatar
	_, err := engine.InsertOne(u)
	if err != nil{
		log.Println(err)
		return 0,"ADD_RESOURCE_ERR"
	}
	return u.Id,"SUCCESS"
}

func(u *User)UserList(page,limit int)ListData{
	users := make([]User,0)
	orm := engine.Where("deleted_at = 0").Desc("created_at")
	cp_orm := *orm
	err := orm.Find(&users)
	log.Println(err)
	total,_ := cp_orm.Count(u)
	return ListData{
		List: users,
		Total: total,
	}
}

func(u *User)UserUpdate(id int,nickname,realname,password,phone,email,avatar string)error{
	u.Nickname = nickname
	u.Realname = realname
	u.Password = password
	u.Phone = phone
	u.Email = email
	u.Avatar = avatar
	_,err := engine.Id(id).Update(u)
	log.Println(err)
	return err
}

func(u *User)UserDelete(id int)error{
	u.DeletedAt = time.Now().Unix()
	_,err := engine.Id(id).Update(u)
	return err
}

func(u *User)UserGet(id int)(*User,error){
	_, err := engine.Id(id).Get(u)
	if err != nil{
		return nil,err
	}
	return u,nil
}
