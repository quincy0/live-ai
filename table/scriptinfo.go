package table

type ScriptInfo struct {
	ScriptId   int64  `json:"script_id" gorm:"PRIMARY_KEY;column:script_id"`
	ScriptName string `json:"script_name" gorm:"column:script_name"`
	ScriptTag  string `json:"script_tag" gorm:"column:script_tag"`
	ProductTag string `json:"product_tag" gorm:"column:product_tag"`
	Summary    string `json:"summary" gorm:"column:summary"`
	Content    string `json:"content" gorm:"column:content"`
	Timbre     string `json:"timbre" gorm:"column:timbre"`
	Audio      string `json:"audio" gorm:"column:audio"`
	LastPlay   int64  `json:"last_play" gorm:"column:last_play"`
}

func (p *ScriptInfo) TableName() string {
	return "script_info"
}
