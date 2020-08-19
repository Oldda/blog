package model

import(
	"log"
    "time"

	"blog/validators"
)

type Article struct{
	Id int `xorm:"id pk autoincr" json:"id"`
    Order int `xorm:"order" json:"order"`
	Cid int `xorm:"cid" json:"cate_id"`
	UserId int `xorm:"user_id" json:"user_id"`
	Title string `xorm:"title" json:"title"`
	Description string `xorm:"description" json:"desc"`
	Body string `xorm:"body" json:"content"`
	CreatedAt int64 `xorm:"created" json:"created_at"`
	UpdatedAt int64 `xorm:"updated" json:"updated_at"`
	DeletedAt int64 `json:"-"`
    Tags []*Tag `xorm:"-"json:"tags"`
    User *User `xorm:"-" json:"user"`
    Category *Category `xorm:"-" json:"category"`
}

func (article *Article)TableName()string{
	return "articles"
}

/**
    文章本身的增删改查应该和关联关系是分开的，尽量是功能单一化。  
**/
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
    article.Order = art.Order
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

func(article *Article)UpdateArticle(art validators.UpdateArticleValidator)bool{
    var err error
    session := engine.NewSession()
    defer session.Close()
    if err = session.Begin(); err != nil {
        return false
    }
    //删除关联关系
    artag := new(TagAndArticle)
    _,err = session.Where("article_id=?",art.Id).Delete(artag)
    //再添加新的关联关系
    for _,tid := range art.TagId{
        t := new(TagAndArticle)
        t.ArticleId = art.Id
        t.TagId = tid
        _,err = session.Insert(t)
        if err != nil{
            session.Rollback()
            log.Println(err)
            return false
        }
    }
    if err == nil{
        //更改文章本身
        article.Cid = art.Cid
        article.Order = art.Order
        article.UserId = art.UserId
        article.Title = art.Title
        article.Description = art.Description
        article.Body = art.Body
        _,err := session.Id(art.Id).Update(article)
        if err == nil{
            session.Commit()
            return true
        }
    }
    session.Rollback()
    return false
}

func(article *Article)GetArticleById(id int)(*Article,error){
    var err error
    _,err = engine.Id(id).Get(article)
    article.Tags = make([]*Tag,0)
    err = engine.Table("tag_and_article").Alias("ta").Join("INNER","tag","ta.tag_id=tag.id").Where("ta.article_id = ?",id).Cols("tag.*").Find(&article.Tags)
    if err != nil{
        return nil,err
    }
    return article,err
}

func (article *Article) ArticleList(page,limit int)ListData{
    list := make([]*Article,0)
    orm := engine.Where("deleted_at = 0").Desc("created_at")
    cp_orm := *orm   
    orm.Limit(limit,(page - 1)*limit).Find(&list)
    for _,value := range list{
        value.Tags = make([]*Tag,0)
        engine.Table("tag_and_article").Alias("ta").Join("INNER","tag","ta.tag_id=tag.id").Where("ta.article_id = ?",value.Id).Cols("tag.*").Find(&value.Tags)
        value.User,_ = new(User).UserGet(value.UserId)
        value.Category,_ = new(Category).GetCateById(1)
        log.Println(value)
    }
    total,_ := cp_orm.Count(article)
    return ListData{
        List: list,
        Total: total,
    }
}

func(article *Article)DeleteArticle(id int)bool{
    article.DeletedAt = time.Now().Unix()
    _,err := engine.Id(id).MustCols("deleted_at").Update(article)
    if err != nil{
        return false
    } 
    return true
}

func (article *Article) OrderArticle(id,order int)bool{
    article.Order = order
    if _,err := engine.Id(id).Update(article);err != nil{
        return false
    }
    return true
}

func(article *Article)QuickSetTag(id,tagId,actType int)bool{
    var err error
    taa := new(TagAndArticle)
    switch actType{
    case 1://添加
        taa.ArticleId = id
        taa.TagId = tagId
        _,err = engine.Insert(taa)
        break
    case 2://删除
        _,err = engine.Where("tag_id=?",tagId).And("article_id=?",id).Delete(taa)
        break
    default:
        return false
        break
    }
    if err != nil{
        return false
    }
    return true
}