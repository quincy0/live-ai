package dto

type WindowGetParam struct {
	RoomId   string `form:"roomId"`
	DouyinId string `form:"douyinId"`
}

type WindowAddParam struct {
	RoomId    string `form:"roomId"`
	DouyinId  string `form:"douyinId"`
	ProductId string `form:"productId"`
}
