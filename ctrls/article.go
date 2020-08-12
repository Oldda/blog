package ctrls

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"blog/business"
	"blog/validators"
	"blog/util"
)

// @Summary 文章新增
// @Description 新增文章和分类的关系是一个分类对应多个文章，和标签的关系是多对多。注：请求格式为json
// @Tags 文章管理
// @accept json
// @Produce  json
// @Param cid body int true "分类id"
// @Param tag_id body int false "标签id"
// @Param title body string true "标题"
// @Param description body string true "描述"
// @Param body body string true "内容"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/article [post]
func CreateArticle(ctx *gin.Context){
	var article validators.CreateArticleValidator
	if err := ctx.ShouldBindWith(&article,binding.JSON);err != nil{
		ctx.JSON(util.NewApiResp("PARAMS_VALIDATE_ERR"))
		return
	}
	if claims,exist := ctx.Get("claims");exist{
		if value,ok := claims.(*util.CustomClaims);ok{
			article.UserId = value.UserId
			lastId,err := business.CreateArticle(article)
			if err != ""{
				ctx.JSON(util.NewApiResp(err))
				return
			}
			ctx.JSON(util.NewApiResp("SUCCESS",lastId))
			return
		}
	}
	ctx.JSON(util.NewApiResp("AUTHORIZATION_TOKEN_ERROR"))
	return
}