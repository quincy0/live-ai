package dto

type DouyinDeleteCard struct {
	DouyinId  string `form:"douyinId" binding:"required"`
	ProductId string `form:"productId" binding:"required"`
}

type DouyinAddCard struct {
	DouyinId  string `form:"douyinId" binding:"required"`
	ProductId string `form:"productId" binding:"required"`
}
