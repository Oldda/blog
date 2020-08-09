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

			//分类管理
			bak.GET("/category",ctrls.CategoryList)
			bak.POST("/category",ctrls.CategoryStore)
			bak.PUT("/category",ctrls.CategoryUpdate)
			bak.DELETE("/category",ctrls.CategoryDelete)

			//标签管理
			bak.GET("/tag",ctrls.TagList)
			bak.POST("/tag",ctrls.TagStore)
			bak.PUT("/tag",ctrls.TagUpdate)
			bak.DELETE("/tag",ctrls.TagDelete)

			//管理员管理
		}
	}
}