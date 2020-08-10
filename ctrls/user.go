package ctrls

import(
	"strconv"
	"crypto/md5"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"blog/business"
	"blog/util"
	"blog/validators"
)

// @Summary 登录
// @Description 统一登录接口，暂时接收账号密码登录和手机验证码登录
// @Tags 用户信息
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param type formData int true "登录类型1:用户名密码登录，2手机号验证码登录"
// @Param email formData string false "登录邮箱"
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
			pwd := fmt.Sprintf("%x",md5.Sum([]byte(ctx.PostForm("password"))))
			userBaseInfo = business.LoginByUsername(ctx.PostForm("email"),pwd)
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

// @Summary 管理员新增
// @Description 新增管理员
// @Tags 用户信息
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param nickname formData string false "用户昵称"
// @Param realname formData string false "用户真实姓名-用于登录使用"
// @Param password formData string true "密码"
// @Param phone formData string true "手机号-可用于登录"
// @Param email formData string true "邮箱"
// @Param avatar formData string false "用户头像"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/user [post]
func UserStore(ctx *gin.Context){
	var user validators.UserStoreValidator
	if err := ctx.ShouldBindWith(&user,binding.FormPost);err != nil{
		ctx.JSON(util.NewApiResp("PARAMS_VALIDATE_ERR"))
		return
	}
	
	//密码加密
	pwd := fmt.Sprintf("%x",md5.Sum([]byte(user.Password)))
	lastId,resp := business.CreateUser(user.Nickname,user.Realname,pwd,user.Phone,user.Email,user.Avatar)
	if lastId == 0{
		ctx.JSON(util.NewApiResp(resp))
		return
	}
	ctx.JSON(util.NewApiResp("SUCCESS",lastId))
}

// @Summary 管理员列表
// @Description 管理员列表
// @Tags 用户信息
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param page query int false "页码"
// @Param limit query int true "每页显示条数"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/user [get]
func UserList(ctx *gin.Context){
	page,_ := strconv.Atoi(ctx.DefaultQuery("page","1"))
	limit,_ := strconv.Atoi(ctx.DefaultQuery("limit","10"))
	ctx.JSON(util.NewApiResp("SUCCESS",business.UserList(page,limit))) 
}

// @Summary 管理员更新
// @Description 编辑管理员
// @Tags 用户信息
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param id query int true "用户id"
// @Param nickname formData string false "用户昵称"
// @Param realname formData string false "用户真实姓名-用于登录使用"
// @Param password formData string true "密码"
// @Param phone formData string true "手机号-可用于登录"
// @Param email formData string true "邮箱"
// @Param avatar formData string false "用户头像"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/user [put]
func UserUpdate(ctx *gin.Context){
	var user validators.UserUpdateValidator
	if err := ctx.ShouldBindWith(&user,binding.FormPost);err != nil{
		ctx.JSON(util.NewApiResp("PARAMS_VALIDATE_ERR"))
		return
	}
	pwd := fmt.Sprintf("%x",md5.Sum([]byte(user.Password)))
	err := business.UserUpdate(user.Id,user.Nickname,user.Realname,pwd,user.Phone,user.Email,user.Avatar)
	if err != nil{
		ctx.JSON(util.NewApiResp("UPDATE_RESOURCE_ERR"))
		return
	}
	ctx.JSON(util.NewApiResp("SUCCESS"))
}

// @Summary 管理员删除
// @Description 删除管理员-暂只支持单个删除
// @Tags 用户信息
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param id query int true "用户id"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/user [delete]
func UserDelete(ctx *gin.Context){
	strId := ctx.Query("id")
	if strId == ""{
		ctx.JSON(util.NewApiResp("MUST_PARAMS_LOST"))
		return
	}
	id,_ := strconv.Atoi(strId)
	if err := business.UserDelete(id);err != nil{
		ctx.JSON(util.NewApiResp("DELETE_RESOURCE_ERR"))
		return
	}
	ctx.JSON(util.NewApiResp("SUCCESS"))
}

// @Summary 管理员详情
// @Description 根据主键ID获取管理员详情
// @Tags 用户信息
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param id query int true "用户id"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/user/:id [get]
func UserGet(ctx *gin.Context){
	strId := ctx.Param("id")
	if strId == ""{
		ctx.JSON(util.NewApiResp("MUST_PARAMS_LOST"))
		return
	}
	id,_ := strconv.Atoi(strId)
	user,err := business.UserGet(id)
	if err != nil{
		ctx.JSON(util.NewApiResp("QUERY_RESOURCE_ERR"))
		return
	}
	ctx.JSON(util.NewApiResp("SUCCESS",user))
}