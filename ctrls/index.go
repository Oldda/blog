package ctrls

import(
	// "log"

	"github.com/gin-gonic/gin"
	// "blog/business"
	// "blog/util"
)


/**
swagger的请求类型
formData
query
path
header
body
**/
func AdminIndex(ctx *gin.Context){
	ctx.String(200,"xxx")
}