package main

import(
	"github.com/gin-gonic/gin"
	"blog/route"
	"blog/model"
)

func main(){
	r := gin.Default()

	//注册路由
	router := route.NewRouter(r)
	router.API()

	//释放野兽
	model.InitDB()

	r.Run("127.0.0.1:8080")
}