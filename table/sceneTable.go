package table

type SceneTable struct {
	SceneId   int64  `json:"scene_id" gorm:"PRIMARY_KEY;column:scene_id"`
	ScriptId  int64  `json:"script_id" gorm:"column:script_id"`
	SceneName string `json:"scene_name" gorm:"column:scene_name"`
	Content   string `json:"content" gorm:"column:content"`
	Audio     string `json:"audio" gorm:"column:audio"`
}

func (t *SceneTable) TableName() string {
	return "scene"
}
