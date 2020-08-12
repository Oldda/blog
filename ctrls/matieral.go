package ctrls

import(
	"log"

	"github.com/gin-gonic/gin"

	"blog/util"
)

// @Summary 获取上传token
// @Description 上传图片获取七牛上传token-暂时仅支持上传，后期改为素材管理
// @Tags 素材管理
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/matieral/token [get]
func MatieralGetToken(ctx *gin.Context){
	token := util.GetQNUploadToken()
	log.Println(token)
	ctx.JSON(util.NewApiResp("SUCCESS",token))
}