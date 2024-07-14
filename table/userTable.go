package table

import "context"

type UserInfo struct {
	UserId   int64  `json:"user_id" gorm:"PRIMARY_KEY;column:user_id"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}

func (user *UserInfo) TableName() string {
	return "user_info"
}

func (user *UserInfo) CheckPassword(ctx context.Context, sign string) bool {
	return user.Password == sign
}
