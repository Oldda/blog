package business

import(
	"blog/model"
)

func GetTags(page,limit int,title string)model.ListData{
	tag := new(model.Tag)
	return tag.GetTags(page,limit,title)
}

func CreateTag(order int,title string)(int,string){
	tag := new(model.Tag)
	return tag.CreateTag(order,title)
}

func UpdateTag(id, order int,title string)error{
	tag := new(model.Tag)
	return tag.UpdateTag(id,order,title)
}

func DeleteTag(id int)error{
	tag := new(model.Tag)
	return tag.DeleteTag(id)
}