package table

type ScriptTable struct {
	ScriptId    int64  `json:"script_id" gorm:"PRIMARY_KEY;column:script_id"`
	UserId      int64  `json:"user_id" gorm:"column:user_id"`
	ScriptTitle string `json:"script_title" gorm:"column:script_title"`
	ScriptTag   string `json:"script_tag" gorm:"column:script_tag"`
	ProductTag  string `json:"product_tag" gorm:"column:product_tag"`
	Summary     string `json:"summary" gorm:"column:summary"`
	Timbre      string `json:"timbre" gorm:"column:timbre"`
	LastPlay    int64  `json:"last_play" gorm:"column:last_play"`
}

func (p *ScriptTable) TableName() string {
	return "script"
}
