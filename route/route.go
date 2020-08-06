package route

import(
	"github.com/gin-gonic/gin"
	"blog/ctrls"
	"blog/middlewares"
)

//路由分发
func API(svr *gin.Engine){
	api := svr.Group("/api")
	{
		//后台管理
		api.POST("/admin/login",ctrls.Login) //登录
		api.POST("/sms",ctrls.SendSMS)//发送短信
		api.GET("/captcha",ctrls.GetCaptcha) //图形验证码
		bak := api.Group("/admin",middlewares.JWTAuth())
		{
			bak.GET("/index",ctrls.AdminIndex)//首页
		}
	}
}