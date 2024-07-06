package table

type RoomInfo struct {
	RoomId int64  `json:"room_id" gorm:"PRIMARY_KEY;column:room_id"`
	Name   string `json:"name" gorm:"column:name"`
}

func (p *RoomInfo) TableName() string {
	return "room_info"
}
