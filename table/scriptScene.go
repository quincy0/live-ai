package table

type ScriptScene struct {
	Id       int64  `json:"id" gorm:"PRIMARY_KEY;column:id"`
	ScriptId int64  `json:"script_id" gorm:"column:script_id"`
	SceneId  int64  `json:"scene_id" gorm:"column:scene_id"`
	GoodsId  int64  `json:"goods_id" gorm:"column:goods_id"`
	Audio    string `json:"audio" gorm:"column:audio"`
}

func (p *ScriptScene) TableName() string {
	return "script_scene"
}
