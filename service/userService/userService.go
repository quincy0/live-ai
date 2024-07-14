package userService

import (
	"context"

	"github.com/quincy0/live-ai/table"
	"github.com/quincy0/qpro/cryption"
	"github.com/quincy0/qpro/qdb"
)

func UserInfo(ctx context.Context, username string) (*table.UserInfo, error) {
	var userInfo table.UserInfo
	err := qdb.Db.WithContext(ctx).
		Model(&table.UserInfo{}).
		Where("username = ?", username).
		Find(&userInfo).
		Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func Register(ctx context.Context, username string, password string) (int64, error) {
	user := &table.UserInfo{
		Username: username,
		Password: cryption.Md5Encode([]byte(password)),
	}
	err := qdb.Db.WithContext(ctx).Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.UserId, nil
}

func Verify(ctx context.Context, username string, sign string) int64 {
	user, err := UserInfo(ctx, username)
	if err != nil {
		return 0
	}
	ok := user.CheckPassword(ctx, sign)
	if !ok {
		return 0
	}
	return user.UserId
}
