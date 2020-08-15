package ctrls

import(
	"strconv"

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
	ctx.JSON(util.NewApiResp("UPDATE_RESOURCE_ERR"))
	return
}

// @Summary 文章更新
// @Description 编辑文章和分类的关系是一个分类对应多个文章，和标签的关系是多对多。注：请求格式为json
// @Tags 文章管理
// @accept json
// @Produce  json
// @Param id body int true "文章id"
// @Param cid body int true "分类id"
// @Param tag_id body int false "标签id"
// @Param title body string true "标题"
// @Param description body string true "描述"
// @Param body body string true "内容"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/article [post]
func UpdateArticle(ctx *gin.Context){
	var article validators.UpdateArticleValidator
	if err := ctx.ShouldBindWith(&article,binding.JSON);err != nil{
		ctx.JSON(util.NewApiResp("PARAMS_VALIDATE_ERR"))
		return
	}
	if claims,exist := ctx.Get("claims");exist{
		if value,ok := claims.(*util.CustomClaims);ok{
			article.UserId = value.UserId
			isTrue := business.UpdateArticle(article)
			if isTrue{
				ctx.JSON(util.NewApiResp("SUCCESS"))
				return
			}
		}
	}
	ctx.JSON(util.NewApiResp("UPDATE_RESOURCE_ERR"))
	return
}

// @Summary 文章详情
// @Description 获取文章详情。和分类的关系是一个分类对应多个文章，和标签的关系是多对多。
// @Tags 文章管理
// @accept x-www-form-urlencoded
// @Produce  json
// @Param id path int true "文章id"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/article/:id [get]
func GetArticleById(ctx *gin.Context){
	article_id,_ := strconv.Atoi(ctx.Param("id"))
	article,err := business.GetArticleById(article_id)
	if err != nil{
		ctx.JSON(util.NewApiResp("QUERY_RESOURCE_ERR"))
		return
	}
	ctx.JSON(util.NewApiResp("SUCCESS",article))
}

// @Summary 文章列表
// @Description 文章列表。和分类的关系是一个分类对应多个文章，和标签的关系是多对多。
// @Tags 文章管理
// @accept x-www-form-urlencoded
// @Produce  json
// @Param page path int false "页码"
// @Param limit path int false "每页显示条数"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/article [get]
func ArticleList(ctx *gin.Context){
	var err error
	page,err := strconv.Atoi(ctx.DefaultQuery("page","1"))
	limit,err := strconv.Atoi(ctx.DefaultQuery("limit","10"))

	if err != nil{
		ctx.JSON(util.NewApiResp("PARAMS_VALIDATE_ERR"))
		return
	}
	articles := business.ArticleList(page,limit)
	ctx.JSON(util.NewApiResp("SUCCESS",articles))
}

// @Summary 文章删除
// @Description 删除文章。和分类的关系是一个分类对应多个文章，和标签的关系是多对多。
// @Tags 文章管理
// @accept x-www-form-urlencoded
// @Produce  json
// @Param id path int true "文章id"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/article [delete]
func DeleteArticle(ctx *gin.Context){
	id,_ := strconv.Atoi(ctx.Query("id"))
	if id == 0{
		ctx.JSON(util.NewApiResp("MUST_PARAMS_LOST"))
		return
	}
	isTrue := business.DeleteArticle(id)
	if !isTrue{
		ctx.JSON(util.NewApiResp("DELETE_RESOURCE_ERR"))
		return
	}
	ctx.JSON(util.NewApiResp("SUCCESS"))
}

// @Summary 文章排序
// @Description 修改文章排序。和分类的关系是一个分类对应多个文章，和标签的关系是多对多。注：请求格式为json
// @Tags 文章管理
// @accept x-www-form-urlencoded
// @Produce  json
// @Param id path int true "文章id"
// @Param order formData int false "顺序，数值越大越靠前"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/article_order/:id [get]
func OrderArticle(ctx *gin.Context){

}

// @Summary 文章标签
// @Description 给文章设置标签，请设置请求格式为json
// @Tags 文章管理
// @accept json
// @Produce  json
// @Param id path int true "文章id"
// @Param tag_id body array true "标签id列表"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/article_tag/:id [put]
func TagSetArticle(ctx *gin.Context){

}