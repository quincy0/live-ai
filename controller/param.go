package controller

type AudioCreateParam struct {
	Recreate int    `form:"recreate"`
	Spk      string `form:"spk" binding:"required"`
	Text     string `form:"text" binding:"required"`
}

type AudioListParam struct {
	RoomId string `form:"roomId" binding:"required"`
}
