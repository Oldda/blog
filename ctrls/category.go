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

// @Summary 分类列表
// @Description 分类列表，无分页，递归实现。
// @Tags 分类管理
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/category [get]
func CategoryList(ctx *gin.Context){
	cates := business.GetCategories()
	ctx.JSON(util.NewApiResp("SUCCESS",cates))
}

// @Summary 分类新增
// @Description 新增分类
// @Tags 分类管理
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param pid formData int true "父级id"
// @Param order formData int false "排序，数值越大越靠前"
// @Param title formData string true "分类名称"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/category [post]
func CategoryStore(ctx *gin.Context){
	var cate validators.CreateCateValidator
	if err := ctx.ShouldBindWith(&cate,binding.FormPost);err != nil{
		ctx.JSON(util.NewApiResp("PARAMS_VALIDATE_ERR"))
		return
	}
	lastId,resp := business.CreateCategory(cate.Pid,cate.Order,cate.Title)
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
// @Router /admin/category [put]
func CategoryUpdate(ctx *gin.Context){
	var cate validators.UpdateCateValidator
	if err := ctx.ShouldBindWith(&cate,binding.Form);err != nil{
		ctx.JSON(util.NewApiResp("PARAMS_VALIDATE_ERR"))
		return
	}
	err := business.UpdateCategory(cate.Id,cate.Pid,cate.Order,cate.Title)
	if err != nil{
		ctx.JSON(util.NewApiResp("UPDATE_RESOURCE_ERR"))
		return
	}
	ctx.JSON(util.NewApiResp("SUCCESS"))
}

// @Summary 分类删除
// @Description 删除分类 暂仅支持单个删除
// @Tags 分类管理
// @accept mpfd,x-www-form-urlencoded
// @Produce  json
// @Param id query int true "分类id"
// @Success 200 {object} util.ApiResp
// @Failure 302 {object} util.ApiResp
// @Router /admin/category [delete]
func CategoryDelete(ctx *gin.Context){
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
	is_del := business.DeleteCategory(idInt)
	if is_del != nil{
		ctx.JSON(util.NewApiResp("DELETE_RESOURCE_ERR"))
		return
	}
	ctx.JSON(util.NewApiResp("SUCCESS"))
}