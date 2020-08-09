package validators

type CreateCateValidator struct{
	Order int `form:"order" json:"order" binding:"min=0"`
	Title string `form:"title" json:"title" binding:"required"`
	Pid int `form:"pid" json:"pid" binding:"min=0"`
}

type UpdateCateValidator struct{
	Id int `form:"id" binding:"required,min=0"`
	Pid int `form:"pid" binding:"min=0"`
	Title string `form:"title" binding:"required"`
	Order int `form:"order" binding:"min=0"`
}