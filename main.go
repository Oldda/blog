package main

import(
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/files" // swagger embed files
	"blog/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/gin-gonic/gin"
	"blog/route"
	"blog/model"
)

func main(){

	// programatically set swagger info
    docs.SwaggerInfo.Title = "个人博客系统API"
    docs.SwaggerInfo.Description = "仅供个人学习使用，不做任何商业活动。"
    docs.SwaggerInfo.Version = "1.0.0"
    docs.SwaggerInfo.Host = "oldda.cn"
    docs.SwaggerInfo.BasePath = "/api"
    docs.SwaggerInfo.Schemes = []string{"http", "https"}

	svr := gin.Default()

	//注册路由
	route.API(svr)
	svr.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//释放野兽
	model.InitDB()

	svr.Run("127.0.0.1:8080")
}