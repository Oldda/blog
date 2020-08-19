package business

import(
	"blog/validators"
	"blog/model"
)

func CreateArticle(article validators.CreateArticleValidator)(int,string){
	arts := new(model.Article)
	return arts.CreateArticle(article)
}

func UpdateArticle(article validators.UpdateArticleValidator)bool{
	arts := new(model.Article)
	return arts.UpdateArticle(article)
}

func GetArticleById(id int)(*model.Article,error){
	arts := new(model.Article)
	return arts.GetArticleById(id)
}

func ArticleList(page,limit int)model.ListData{
	arts := new(model.Article)
	return arts.ArticleList(page,limit)
}

func DeleteArticle(id int)bool{
	arts := new(model.Article)
	return arts.DeleteArticle(id)
}

func OrderArticle(id,order int)bool{
	arts := new(model.Article)
	return arts.OrderArticle(id,order)
}

func TagSetArticle(id int,tagId int,actType int)bool{
	arts := new(model.Article)
	return arts.QuickSetTag(id,tagId,actType)
}