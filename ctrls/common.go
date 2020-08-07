package ctrls

import(
	"github.com/gin-gonic/gin"
	"github.com/dchest/captcha"
	"blog/util"
	"math/rand"
	"time"
	"log"
	"fmt"
)

// @Summary 短信接口
// @Description 发送短信接口
// @Tags 公共接口
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param phone formData string true "手机号"
// @Param captcha formData string true "图形验证码的值"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /sms [post]
func SendSMS(ctx *gin.Context){
	phone := ctx.PostForm("phone")
	captcha := ctx.PostForm("captcha")
	//验证图形验证码
	if !verifyCaptcha(phone,captcha){
		ctx.JSON(util.NewApiResp("CAPTCHA_ERR_OR_EXPIRED"))
		return
	}
	//发送短信
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(9999-1000)+1000
	log.Println(num)
	util.RdbSetString(phone+"_sms",num,time.Minute*5) //redis
	err := util.AliSendSMS(phone,fmt.Sprintf("%d",num)) //发送短信
	if err != nil{
		log.Println(err)
		ctx.JSON(util.NewApiResp("SMS_SEND_ERROR"))
		return
	}
	ctx.JSON(util.NewApiResp("SUCCESS"))
}

// @Summary 图形验证码
// @Description 获取图形验证码图片
// @Tags 公共接口
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param phone query string true "手机号"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /captcha [get]
func GetCaptcha(ctx *gin.Context){
	phone := ctx.Query("phone")
	if phone == ""{
		ctx.JSON(util.NewApiResp("MUST_PARAMS_LOST"))
		return
	}

	id := captcha.NewLen(5)
	//存到redis
	e := util.RdbSetString(phone,id,time.Minute*5) //5分钟有效。
	log.Println(e)
	captcha.WriteImage(ctx.Writer,id,captcha.StdWidth,captcha.StdHeight)
}

//验证图形验证码
func verifyCaptcha(phone string,pattern string)bool{
	id,err := util.RdbGetString(phone)
	if err == nil{
		if captcha.VerifyString(id, pattern){
			return true
		}
	}
	return false
}