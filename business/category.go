package business

import(
	"blog/model"
)

func CreateCategory(pid,order int,title string)(int,string){
	cate := new(model.Category)
	return cate.CreateCategory(pid,order,title)
}

func GetCategories()[]*model.Category{
	cate := new(model.Category)
	cates := make([]*model.Category,0)
	cate.GetCategories(0,&cates)
	return cates
}

func UpdateCategory(id,pid,order int,title string)error{
	cate := new(model.Category)
	return cate.UpdateCategory(id,pid,order,title)
}

func DeleteCategory(id int)error{
	cate := new(model.Category)
	return cate.DeleteCategory(id)
}