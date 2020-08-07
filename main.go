package main

import(
	"context"

	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/files" // swagger embed files
	"github.com/gin-gonic/gin"

	"blog/route"
	"blog/model"
	"blog/util"
	"blog/docs" // docs is generated by Swag CLI, you have to import it.
)

func main(){
	// programatically set swagger info
    docs.SwaggerInfo.Title = "个人博客系统API"
    docs.SwaggerInfo.Description = "仅供个人学习使用，不做任何商业活动。"
    docs.SwaggerInfo.Version = "1.0.0"
    docs.SwaggerInfo.Host = "blog.oldda.cn"
    docs.SwaggerInfo.BasePath = "/api"
    docs.SwaggerInfo.Schemes = []string{"http", "https"}

	svr := gin.Default()

	//注册路由
	route.API(svr)

	//swagger api文档
	svr.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//释放野兽
	model.InitDB()

	//redis
	//空 context。这只能用于高等级（在 main 或顶级请求处理中）
	ctx := context.Background()
	util.NewRedisEngine("39.106.116.186:6379","Aa890223",8,ctx)

	svr.Run("127.0.0.1:8080")
}