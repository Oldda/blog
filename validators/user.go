package validators

import(
	"regexp"
	"github.com/go-playground/validator/v10"
)

type UserStoreValidator struct{
	Nickname string `form:"nickname" binding:"max=15"`
	Realname string `form:"realname" binding:"max=10"`
	Password string `form:"password" binding:"required"`
	Phone string `form:"phone" binding:"required,validphone"`
	Email string `form:"email" binding:"required,validemail"`
	Avatar string `form:"avatar"`
}

type UserUpdateValidator struct{
	Id int `form:"id" binding:"required"`
	Nickname string `form:"nickname" binding:"max=15"`
	Realname string `form:"realname" binding:"max=10"`
	Password string `form:"password" binding:"required"`
	Phone string `form:"phone" binding:"required,validphone"`
	Email string `form:"email" binding:"required,validemail"`
	Avatar string `form:"avatar"`
}

//验证邮箱
var Validemail = func(fl validator.FieldLevel)bool{
	if value,ok := fl.Field().Interface().(string);ok{
		if value == ""{
			return true
		}
		pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

		reg := regexp.MustCompile(pattern)
		return reg.MatchString(value)
	}
	return false
}

//验证手机号
var Validphone = func(fl validator.FieldLevel)bool{
	if value,ok := fl.Field().Interface().(string);ok{
		if value == ""{
			return true 
		}
		pattern := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`

		reg := regexp.MustCompile(pattern)
		return reg.MatchString(value)
	}
	return false
}