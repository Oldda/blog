package ctrls

import(
	"github.com/gin-gonic/gin"
	"fmt"
)

type User struct{

}

func NewUser()*User{
	return &User{}
}

//登录
func(u *User)Login(ctx *gin.Context){
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	captcha := ctx.PostForm("captcha")
	if password != ""{ //账号密码登录

	}
}

func(u *User)List(ctx *gin.Context){

}