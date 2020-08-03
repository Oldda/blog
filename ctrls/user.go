package ctrls

import(
	"github.com/gin-gonic/gin"
	"blog/business"
	"blog/util"
)
// @Summary 登录
// @Description 统一登录接口，暂时接收账号密码登录和手机验证码登录
// @Tags 用户信息
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param type formData int true "登录类型1:用户名密码登录，2手机号验证码登录"
// @Param username formData string false "用户真实姓名"
// @Param password formData string false "密码"
// @Param phone formData string false "手机号"
// @Param captcha formData string false "手机验证码"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/login [post]
func Login(ctx *gin.Context){
	var userBaseInfo *business.UserBaseInfo
	switch ctx.PostForm("type"){
		case "1": //账号密码登录
			userBaseInfo = business.LoginByUsername(ctx.PostForm("username"),ctx.PostForm("password"))
			break
		case "2": //手机号验证码登录
			userBaseInfo = business.LoginByPhone(ctx.PostForm("phone"),ctx.PostForm("captcha"))
			break
		default:
			ctx.JSON(util.NewApiResp("LOGIN_METHOD_NOT_ALLOWD"))
			return
			break 
	}
	if userBaseInfo == nil{
		ctx.JSON(util.NewApiResp("UNAME_OR_PWD_ERROR"))
		return
	}
	ctx.JSON(util.NewApiResp("SUCCESS",userBaseInfo))
}