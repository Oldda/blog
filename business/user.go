package business

import(
	"blog/model"
	"fmt"
)

type User struct{
	*Common
	user *model.User
}

func NewUser()*User{
	return &User{
		user:new(model.User),
	}
}

func(u *User)LoginByUsername(username,password string)(*model.User,error){
	has,_ := model.DB.Engine.Where("realname=? and password=?",username,password).Get(u.user)
	if has {
		token,err := u.GenerateToken(u.user.Id)
		fmt.Println(token)
		if err != nil{
			return nil,err
		}
	}
	return u.user,nil
}

func(u *User)LoginByPhone(phone,captcha string)(*model.User,error){
	return nil,nil
}