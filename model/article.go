package model

import(
	"log"

	"blog/validators"
)

type Article struct{
	Id int `xorm:"id pk autoincr" json:"id"`
	Cid int `xorm:"cate_id" json:"cate_id"`
	UserId int `xorm:"user_id" json:"user_id"`
	Title string `xorm:"title" json:"title"`
	Description string `xorm:"description" json:"desc"`
	Body string `xorm:"body" json:"content"`
	CreatedAt int `xorm:"created" json:"created_at"`
	UpdatedAt int `xorm:"updated" json:"updated_at"`
	DeletedAt int `json:"-"`
}

func (article *Article)TableName()string{
	return "articles"
}

func (article *Article)CreateArticle(art validators.CreateArticleValidator)(int,string){
	//事务处理
	var err error
	session := engine.NewSession()
    defer session.Close()
    if err = session.Begin(); err != nil {
    	log.Println(err)
        return 0,"ADD_RESOURCE_ERR"
    }
    //存储文章
    article.Cid = art.Cid
    article.UserId = art.UserId
    article.Title = art.Title
    article.Description = art.Description
    article.Body = art.Body
    _,err = session.Insert(article)
    if err != nil{
    	session.Rollback()
    	log.Println(err)
    	return 0,"ADD_RESOURCE_ERR"
    }
    //存储关联关系
    for _,tid := range art.TagId{
    	t := new(TagAndArticle)
    	t.ArticleId = article.Id
    	t.TagId = tid
    	_,err = session.Insert(t)
    	if err != nil{
    		session.Rollback()
    		log.Println(err)
        	return 0,"ADD_RESOURCE_ERR"
    	}
    }
    log.Println(err)
	session.Commit()
	return article.Id,""
}