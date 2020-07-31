package route

import(
	"github.com/gin-gonic/gin"
	"blog/ctrls"
	"blog/middlewares"
)


type Router struct{
	engine *gin.Engine
	tagCtrl *ctrls.Tag
	userCtrl *ctrls.User
}

//初始化引擎和控制器
func NewRouter(engine *gin.Engine)*Router{
	return &Router{
		engine : engine,
		tagCtrl : ctrls.NewTag(),
		userCtrl : ctrls.NewUser(),
	}
}

//路由分发
func (ro *Router)API(){
	//后台管理
	ro.engine.POST("/admin/login",ro.userCtrl.Login) //登录
	//找回密码
	bak := ro.engine.Group("/admin",middlewares.JWTAuth())
	{
		//标签管理
		tag := bak.Group("/tag")
		{
			//标签列表
			tag.GET("/list",ro.tagCtrl.Index)
		}
		//用户管理
		user := bak.Group("/user")
		{
			user.GET("/list",ro.userCtrl.List)
		}
	}

	//前台展示
	ro.engine.GET("/",ro.tagCtrl.Admin)
}