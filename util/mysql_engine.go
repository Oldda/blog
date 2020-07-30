package util

import(
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
)

type DB struct{
	Engine *xorm.Engine
	IsShowSql bool
	MaxIdleConns int
	MaxOpenConns int
	TableMapper core.IMapper
	ColumnMapper core.IMapper
}

//初始化引擎
func NewDB(addr,port,uname,password,database,char string)(*DB,error){
	//"root:Aa890223@tcp(127.0.0.1:3308)/blog?charset=utf8")
	engine,err := xorm.NewEngine("mysql", uname+":"+password+"@tcp("+addr+":"+port+")/"+database+"?charset="+char)
	if err != nil{
		return nil,err
	}
	return &DB{
		Engine : engine,
	},nil
}

//测试链接
func(db *DB)TestConnect()error{
	return db.Engine.Ping()
}

func(db *DB)SetShowSql(isShow bool){
	db.IsShowSql = isShow
	db.Engine.ShowSQL(isShow)
}

func(db *DB)SetMaxIdleConns(num int){
	db.MaxIdleConns = num
	db.Engine.SetMaxIdleConns(num)
}

func(db *DB)SetMaxOpenConns(num int){
	db.MaxOpenConns = num
	db.Engine.SetMaxOpenConns(num)
}

func(db *DB)SetTableMapper(mapper core.IMapper){
	db.TableMapper = mapper
	db.Engine.SetTableMapper(mapper)
}

func(db *DB)SetColumnMapper(mapper core.IMapper){
	db.ColumnMapper = mapper
	db.Engine.SetColumnMapper(mapper)
}