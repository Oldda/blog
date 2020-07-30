package route

import(
	"github.com/gin-gonic/gin"
	"blog/ctrls"
)


type Router struct{
	engine *gin.Engine
	tagCtrl *ctrls.Tag
}

//初始化引擎和控制器
func NewRouter(engine *gin.Engine)*Router{
	return &Router{
		engine : engine,
		tagCtrl : ctrls.NewTag(),
	}
}

//路由分发
func (ro *Router)API(){
	//后台管理
	bak := ro.engine.Group("/admin")
	{
		//标签管理
		tag := bak.Group("/tag")
		{
			//标签列表
			tag.GET("/list",ro.tagCtrl.Index)
		}
	}

	//前台展示
	ro.engine.GET("/",ro.tagCtrl.Admin)
}