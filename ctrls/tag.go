package ctrls

import(
	"github.com/gin-gonic/gin"
	"fmt"
	"blog/model"
)

type Tag struct{
	
}

func NewTag()*Tag{
	return &Tag{
		
	}
}

func(t *Tag)Admin(ctx *gin.Context){

	user := new(model.User)
	model.DB.Engine.Sync2(user)
	fmt.Fprintln(ctx.Writer,"domain/admin")
}

func (t *Tag)Index(ctx *gin.Context){
	fmt.Fprintln(ctx.Writer,"domain")
}