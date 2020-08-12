package business

import(
	"blog/validators"
	"blog/model"
)

func CreateArticle(article validators.CreateArticleValidator)(int,string){
	arts := new(model.Article)
	return arts.CreateArticle(article)
}