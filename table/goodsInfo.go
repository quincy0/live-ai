package table

type GoodsInfo struct {
	GoodsId int64  `json:"goods_id" gorm:"PRIMARY_KEY;column:goods_id"`
	Name    string `json:"name" gorm:"column:name"`
}

func (p *GoodsInfo) TableName() string {
	return "goods_info"
}
