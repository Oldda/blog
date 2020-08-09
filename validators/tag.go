package validators

type CreateTagValidator struct{
	Order int `form:"order" json:"order" binding:"min=0"`
	Title string `form:"title" json:"title" binding:"required"`
}

type UpdateTagValidator struct{
	Id int `form:"id" binding:"required,min=0"`
	Title string `form:"title" binding:"required"`
	Order int `form:"order" binding:"min=0"`
}