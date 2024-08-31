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

type WindowUpdateParam struct {
	ProductionId string `form:"productionId"`
	Credit       int    `form:"credit"`
	Count        int    `form:"count"`
	DouyinId     string `form:"douyinId"`
}
