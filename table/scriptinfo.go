package table

type ScriptInfo struct {
	ScriptId   int64  `json:"script_id" gorm:"PRIMARY_KEY;column:script_id"`
	ScriptName string `json:"script_name" gorm:"column:script_name"`
	GoodsId    int64  `json:"goods_id" gorm:"column:goods_id"`
}

func (p *ScriptInfo) TableName() string {
	return "script_info"
}
