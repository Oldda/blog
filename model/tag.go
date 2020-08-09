package model

import(
	"time"
)

type Tag struct{
	Id int `xorm:"id pk autoincr" json:"id"`
	Order int `xorm:"order" json:"order"`
	Title string `xorm:"title" json:"title"`
	ArticleCount int `xorm:"article_count" json:"article_count"`
	CreatedAt int64 `xorm:"created" json:"created_at"`
	UpdatedAt int64 `xorm:"updated" json:"updated_at"`
	DeletedAt int64 `xorm:"deleted_at" json:"-"`
}

func (tag *Tag)GetTags(page,limit int,title string)ListData{
	orm := engine.Where("deleted_at=?",0)
	if title != ""{
		orm = orm.And("title like ?","%"+title+"%")
	}
	cp_orm := *orm
	list := make([]*Tag,0)
	orm.Limit(limit,(page-1)*limit).Desc("order").Find(&list)
	total,_ := cp_orm.Count(tag)
	return ListData{
		List: list,
		Total: total,
	}
}

func (tag *Tag)CreateTag(order int, title string)(int, string){
	tag.Order = order
	tag.Title = title
	_,err := engine.Insert(tag)
	if err != nil{
		return 0,"ADD_RESOURCE_ERR"
	}
	return tag.Id,"SUCCESS"
}

func (tag *Tag)UpdateTag(id, order int,title string)error{
	tag.Order = order
	tag.Title = title
	_,err := engine.Id(id).Update(tag)
	return err
}

func (tag *Tag)DeleteTag(id int)error{
	tag.DeletedAt = time.Now().Unix()
	_,err := engine.Id(id).Update(tag)
	return err
}

type TagAndArticle struct{
	Id int `xorm:"id pk autoincr" json:"id"`
	ArticleId int `xorm:"article_id" json:"article_id"`
	TagId int `xorm:"tag_id" json:"tag_id"`
	CreatedAt int64 `xorm:"created" json:"created_at"`
	UpdatedAt int64 `xorm:"updated" json:"updated_at"`
	DeletedAt int64 `json:"-"`
}