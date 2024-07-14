package table

type RoomTable struct {
	RoomId     int64  `json:"room_id" gorm:"PRIMARY_KEY;column:room_id"`
	UserId     int64  `json:"user_id" gorm:"column:user_id"`
	RoomName   string `json:"room_name" gorm:"column:room_name"`
	ProductTag string `json:"product_tag" gorm:"column:product_tag"`
	Timbre     string `json:"timbre" gorm:"column:timbre"`
	TemplateId int    `json:"template_id" gorm:"column:template_id"`
}

func (t *RoomTable) TableName() string {
	return "room_info"
}

type RoomScriptTable struct {
	Id          int64  `json:"id" gorm:"PRIMARY_KEY;column:id"`
	RoomId      int64  `json:"room_id" gorm:"column:room_id"`
	ScriptId    int64  `json:"script_id" gorm:"column:script_id"`
	ScriptTitle string `json:"script_title" gorm:"column:script_title"`
	ProductTag  string `json:"product_tag" gorm:"column:product_tag"`
	Timbre      string `json:"timbre" gorm:"column:timbre"`
	ScriptTag   string `json:"script_tag" gorm:"column:script_tag"`
	Sequence    int    `json:"sequence" gorm:"column:sequence"`
}

func (t *RoomScriptTable) TableName() string {
	return "room_script"
}
