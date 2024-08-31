package table

import "time"

type GoodsInfo struct {
	GoodsId int64  `json:"goods_id" gorm:"PRIMARY_KEY;column:goods_id"`
	Name    string `json:"name" gorm:"column:name"`
}

func (p *GoodsInfo) TableName() string {
	return "goods_info"
}

type GoodsWindow struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment;column:id"`
	Name        string    `json:"name" gorm:"column:name;type:varchar(255);not null"`
	PromotionID string    `json:"promotion_id" gorm:"column:promotion_id;type:varchar(255);not null"`
	ProductID   string    `json:"product_id" gorm:"column:product_id;type:varchar(255);not null;unique"`
	Title       string    `json:"title" gorm:"column:title;type:varchar(255);not null"`
	Cover       string    `json:"cover" gorm:"column:cover;type:varchar(255);not null"`
	Price       int       `json:"price" gorm:"column:price;not null"`
	CosFee      int       `json:"cos_fee" gorm:"column:cos_fee;not null"`
	CosRatio    int       `json:"cos_ratio" gorm:"column:cos_ratio;not null"`
	ShopID      string    `json:"shop_id" gorm:"column:shop_id;type:varchar(255);not null"`
	ShopName    string    `json:"shop_name" gorm:"column:shop_name;type:varchar(255);not null"`
	RoomID      string    `json:"room_id" gorm:"column:room_id;type:varchar(255);not null"`
	Type        string    `json:"type" gorm:"column:type;type:varchar(255);not null"`
	AccountID   string    `json:"account_id" gorm:"column:account_id;type:varchar(255);not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
	Credit      int       `json:"credit" gorm:"column:credit"`
	TotalAmount int       `json:"total_amount" gorm:"column:total_amount"`
	CouponID    string    `json:"coupon_id" gorm:"column:coupon_id;type:varchar(100)"`
}

func (GoodsWindow) TableName() string {
	return "goods_window"
}
