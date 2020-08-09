package model

import(
	"time"
	"blog/util"
	"github.com/go-xorm/xorm"
	"log"
)

var engine *xorm.Engine

var DB *util.DB

func InitDB(){
	var err error
	DB,err = util.NewDB("127.0.0.1","3308","root","Aa890223","blog","utf8")
	if err != nil{
		log.Println(err)
	}
	engine = DB.Engine
	engine.TZLocation, _ = time.LoadLocation("Asia/Shanghai")
}

type ListData struct{
	List interface{} `json:"list"`
	Total int64 `json:"count"`
}