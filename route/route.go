package route

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"blog/ctrls"
	"blog/middlewares"
	vd "blog/validators"
)

//路由分发
func API(svr *gin.Engine){
	api := svr.Group("/api",middlewares.Cors())
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
			bak.GET("/user",ctrls.UserList)
			bak.POST("/user",ctrls.UserStore)
			bak.PUT("/user",ctrls.UserUpdate)
			bak.DELETE("/user",ctrls.UserDelete)
			bak.GET("/user/:id",ctrls.UserGet)

			//素材管理
			bak.GET("/matieral/token",ctrls.MatieralGetToken)

			//文章管理
			bak.POST("/article",ctrls.CreateArticle)
		}
	}

	//注册自定义验证器--gin默认的v8验证器有错误，报没有实例化validator的错误。
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        v.RegisterValidation("validemail", vd.Validemail)
        v.RegisterValidation("validphone", vd.Validphone)
	}
}