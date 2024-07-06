package goodsService

import (
	"context"
	"fmt"

	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/table"
	"github.com/quincy0/qpro/qdb"
)

func List(ctx context.Context, page dto.PageParam) ([]table.GoodsInfo, int64) {
	var list []table.GoodsInfo
	err := qdb.Db.WithContext(ctx).
		Model(&table.GoodsInfo{}).
		Order("goods_id").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).
		Scan(&list).
		Error
	if err != nil {
		fmt.Println(err)
	}
	var count int64
	err = qdb.Db.WithContext(ctx).Model(&table.GoodsInfo{}).Count(&count).Error
	if err != nil {
		fmt.Println(err)
	}
	return list, count
}

/*
func BatchGetLiveData(ctx context.Context, status []string) []LiveStreamAndPlay {
	var list []LiveStreamAndPlay
	database.Db.WithContext(ctx).
		Model(&table.LiveStream{}).
		Select("live_stream.*, play_info.*").
		Joins("join play_info on play_info.id = live_stream.play_id").
		Where("live_stream.live_status in ?", status).
		Order("play_info.id").
		Limit(100).
		Scan(&list)
	return list
}

func GetLiveStream(ctx context.Context, roomId string) *table.LiveStream {
	var stream table.LiveStream
	database.Db.WithContext(ctx).Find(&stream, "room_id = ?", roomId)
	return &stream
}
*/
