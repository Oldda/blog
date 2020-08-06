package util

import(
	"context"
	"github.com/go-redis/redis/v8" //redis
	"log"
	"time"
)

var c context.Context

var Rdb *redis.Client

func NewRedisEngine(addr string,pwd string,db int,ctx context.Context){
	c = ctx
	Rdb = redis.NewClient(&redis.Options{
	    Addr:     addr,
	    Password: pwd, // no password set
	    DB:       db,  // use default DB
	})
	if _,err := Rdb.Ping(c).Result();err != nil{
		log.Println("redis connect failed! addr:"+addr)
	}
}

//设置字符串
func RdbSetString(key string,value interface{},expire time.Duration)error{
	//0表示不设置时间
	return Rdb.Set(c,key,value,expire).Err()
}

//获取字符串
func RdbGetString(key string)(string,error){
	return Rdb.Get(c,key).Result()
}