package dto

type AudioCreateParam struct {
	Recreate int    `form:"recreate"`
	Spk      string `form:"spk" binding:"required"`
	Text     string `form:"text" binding:"required"`
}

type AudioListParam struct {
	RoomId string `form:"roomId" binding:"required"`
}

type AudioNotifyParam struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		TaskId   string `json:"task_id"`
		VideoUrl string `json:"video_url"`
	} `json:"data"`
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

type ChapterCreateParam struct {
	ChapterTitle string           `form:"chapterTitle"`
	ScriptTag    string           `form:"scriptTag"`
	ProductTag   string           `form:"productTag"`
	Summary      string           `form:"summary"`
	Paragraph    []*ParagraphItem `form:"paragraph"`
}

type ChapterInfoParam struct {
	ChapterId int64 `form:"chapterId"`
}

type ParagraphItem struct {
	ParagraphTitle string `form:"paragraphTitle"`
	Content        string `form:"content"`
}

type ParagraphEditParam struct {
	ParagraphId    int64  `form:"paragraphId"`
	ParagraphTitle string `form:"paragraphTitle"`
	Content        string `form:"content"`
}

type ScriptCreateParam struct {
	ChapterId   int64  `form:"chapterId"`
	ScriptTitle string `form:"scriptTitle"`
	Timbre      string `form:"timbre"`
}

type SceneEditParam struct {
	SceneId int64  `form:"sceneId" binding:"required"`
	Audio   string `form:"audio" binding:"required"`
}

type RoomCreateParam struct {
	RoomName   string `form:"roomName" binding:"required"`
	ProductTag string `form:"productTag" binding:"required"`
	Timbre     string `form:"timbre" binding:"required"`
	TemplateId int    `form:"templateId" binding:"required"`
}

type RoomInfoParam struct {
	RoomId int64 `form:"roomId" binding:"required"`
}

type LoginParam struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type TimbreParam struct {
	Key  string `form:"key" json:"key" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}
