package validators

type CreateArticleValidator struct{
	Cid int `json:"cid" binding:"required"`
	Order int `json:"order" binding:"min=0"`
	TagId []int `json:"tag_id"`
	UserId int
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required,max=255"`
	Body string `json:"body" binding:"required"`
}

type UpdateArticleValidator struct{
	Id int `json:"id" binding:"required"`
	Order int `json:"order" binding="min=0"`
	Cid int `json:"cid" binding:"required"`
	TagId []int `json:"tag_id"`
	UserId int
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required,max=255"`
	Body string `json:"body" binding:"required"`
}