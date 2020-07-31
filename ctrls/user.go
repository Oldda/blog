package ctrls

import(
	"github.com/gin-gonic/gin"
	"blog/business"
	"blog/model"
	"fmt"
)

type User struct{
	ub *business.User
}

func NewUser()*User{
	return &User{
		ub:business.NewUser(),
	}
}

//登录
func(u *User)Login(ctx *gin.Context){
	var user *model.User
	var err error
	switch ctx.PostForm("type"){
		case "1": //账号密码登录
			user,err = u.ub.LoginByUsername(ctx.PostForm("username"),ctx.PostForm("password"))
			break
		case "2": //手机号验证码登录
			user,err = u.ub.LoginByPhone(ctx.PostForm("phone"),ctx.PostForm("captcha"))
			break
		default:
			break 
	}
	if err != nil{
		ctx.JSON(302,gin.H{"status":false,"response_code":1000,"msg":"登录失败","data":nil})
		return
	}
	ctx.JSON(200,gin.H{"status":true,"response_code":200,"msg":"success","data":user})
}

func(u *User)List(ctx *gin.Context){
	fmt.Fprintln(ctx.Writer,"hhh")
}