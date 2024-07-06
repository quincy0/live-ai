package table

type RoomScript struct {
	Id       int64  `json:"id" gorm:"PRIMARY_KEY;column:id"`
	RoomId   int64  `json:"room_id" gorm:"column:room_id"`
	ScriptId int64  `json:"script_id" gorm:"column:script_id"`
	Audio    string `json:"audio" gorm:"column:audio"`
}

func (p *RoomScript) TableName() string {
	return "room_script"
}
