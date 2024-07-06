package dto

type AudioCreateParam struct {
	Recreate int    `form:"recreate"`
	Spk      string `form:"spk" binding:"required"`
	Text     string `form:"text" binding:"required"`
}

type AudioListParam struct {
	RoomId string `form:"roomId" binding:"required"`
}

type AudioUploadParam struct {
	ScriptId int64  `form:"scriptId" binding:"required"`
	Audio    string `form:"audio" binding:"required"`
}

type CreateScriptParam struct {
	GoodsId    int64  `form:"goodsId"`
	ScriptId   int64  `form:"scriptId"`
	ScriptName string `form:"scriptName"`
	Audio      string `form:"audio" binding:"required"`
}

type ScriptInfoParam struct {
	ScriptId int64 `form:"scriptId" binding:"required"`
}

type PageParam struct {
	PageSize int `form:"pageSize"`
	PageNum  int `form:"pageNum"`
}
