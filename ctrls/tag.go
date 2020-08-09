package ctrls

import(
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"blog/business"
	"blog/validators"
	"blog/util"
	// "log"
)

// @Summary 标签列表
// @Description 标签列表。
// @Tags 标签管理
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param page query int false "页码"
// @Param limit query int false "每页显示条数"
// @Param title query string false "名称筛选"
// @Success 200 {object} model.ListData
// @Failure 302 {object} util.ApiResp
// @Router /admin/tag [get]
func TagList(ctx *gin.Context){
	page,_ := strconv.Atoi(ctx.DefaultQuery("page","1"))
	limit,_ := strconv.Atoi(ctx.DefaultQuery("limit","10"))
	title := ctx.DefaultQuery("title","")
	tags := business.GetTags(page,limit,title)
	ctx.JSON(util.NewApiResp("SUCCESS",tags))
}

// @Summary 标签新增
// @Description 新增标签
// @Tags 标签管理
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param order formData int false "排序，数值越大越靠前"
// @Param title formData string true "标签名称"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/tag [post]
func TagStore(ctx *gin.Context){
	var tag validators.CreateTagValidator
	if err := ctx.ShouldBindWith(&tag,binding.FormPost);err != nil{
		ctx.JSON(util.NewApiResp("PARAMS_VALIDATE_ERR"))
		return
	}
	lastId,resp := business.CreateTag(tag.Order,tag.Title)
	if lastId == 0{
		ctx.JSON(util.NewApiResp(resp))
		return
	}
	ctx.JSON(util.NewApiResp("SUCCESS",lastId))
	return
}

// @Summary 分类更新
// @Description 编辑分类
// @Tags 分类管理
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param id query int true "分类id"
// @Param pid formData int true "父级id"
// @Param order formData int false "排序，数值越大越靠前"
// @Param title formData string true "分类名称"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/tag [put]
func TagUpdate(ctx *gin.Context){
	var tag validators.UpdateTagValidator
	if err := ctx.ShouldBindWith(&tag,binding.Form);err != nil{
		ctx.JSON(util.NewApiResp("PARAMS_VALIDATE_ERR"))
		return
	}
	err := business.UpdateTag(tag.Id,tag.Order,tag.Title)
	if err != nil{
		ctx.JSON(util.NewApiResp("UPDATE_RESOURCE_ERR"))
		return
	}
	ctx.JSON(util.NewApiResp("SUCCESS"))
}

// @Summary 标签删除
// @Description 删除标签 暂仅支持单个删除
// @Tags 分类管理
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param id query int true "标签id"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/tag [delete]
func TagDelete(ctx *gin.Context){
	id := ctx.Query("id")
	if id == ""{
		ctx.JSON(util.NewApiResp("MUST_PARAMS_LOST"))
		return
	}
	idInt,err := strconv.Atoi(id)
	if err != nil{
		ctx.JSON(util.NewApiResp("PARAMS_VALIDATE_ERR"))
		return
	}
	is_del := business.DeleteTag(idInt)
	if is_del != nil{
		ctx.JSON(util.NewApiResp("DELETE_RESOURCE_ERR"))
		return
	}
	ctx.JSON(util.NewApiResp("SUCCESS"))
}