package model

import(
	"time"
	"log"
)

type Category struct{
	Id int `xorm:"id pk autoincr" json:"id"` //因此添加数据的时候返回了自增id
	Pid int `xorm:"pid" json:"pid"`
	Order int `xorm:"order" json:"order"`
	Title string `xorm:"title" json:"title"`
	ArticleCount int `xorm:"article_count" json:"article_count"`
	CreatedAt int64 `xorm:"created" json:"created_at"`
	UpdatedAt int64 `xorm:"updated" json:"updated_at"`
	DeletedAt int64 `json:"-"`
	Children []*Category
}

func (cate *Category)CreateCategory(pid,order int, title string)(int,string){
	cate.Order = order
	cate.Title = title
	cate.Pid = pid
	affecte, err := engine.InsertOne(cate)
	if err != nil{
		return 0,"ADD_RESOURCE_ERR"
	}
	log.Println(cate)
	log.Println(affecte)
	return cate.Id,"SUCCESS"
}

func(cate *Category)GetCategories(pid int,cates *[]*Category){
	err := engine.Cols("id","pid","order","title","article_count").Where("pid=?",pid).And("deleted_at=0").Desc("order").Find(cates)
	log.Println(err)
	for _,cate := range *cates{
		cate.GetCategories(cate.Id,&cate.Children)
	}
}

func (cate *Category)UpdateCategory(id,pid,order int,title string)error{
	cate.Pid = pid
	cate.Order = order
	cate.Title = title
	_,err := engine.Id(id).Update(cate)
	return err
}

func(cate *Category)DeleteCategory(id int)error{
	cate.DeletedAt = time.Now().Unix()
	log.Println(cate)
	log.Println(id)
	_,err := engine.Id(id).Update(cate)
	return err
}