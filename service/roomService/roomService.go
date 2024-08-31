package roomService

import (
	"context"
	"errors"

	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/table"
	"github.com/quincy0/qpro/qLog"
	"github.com/quincy0/qpro/qdb"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TagItem struct {
	Count int                  `json:"count"`
	Data  []*table.ScriptTable `json:"data"`
}

func RoomCreate(ctx context.Context, userId int64, params dto.RoomCreateParam) (int64, error) {
	var scriptList []table.ScriptTable

	err := qdb.Db.WithContext(ctx).
		Model(&table.ScriptTable{}).
		Where("script_id in ?", params.Scripts).
		Scan(&scriptList).
		Error
	if err != nil {
		return 0, err
	}
	if len(scriptList) == 0 {
		return 0, errors.New("illegal scripts")
	}

	roomInfo := &table.RoomTable{
		RoomId:     0,
		UserId:     userId,
		RoomName:   params.RoomName,
		ProductTag: scriptList[0].ProductTag,
		Timbre:     scriptList[0].Timbre,
		TemplateId: 0,
	}

	roomScriptList := make([]*table.RoomScriptTable, len(scriptList))

	err = qdb.Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(roomInfo).Error; err != nil {
			return err
		}
		for k, v := range scriptList {
			roomScriptList[k] = &table.RoomScriptTable{
				RoomId:      roomInfo.RoomId,
				ScriptId:    v.ScriptId,
				ScriptTitle: v.ScriptTitle,
				ProductTag:  v.ProductTag,
				Timbre:      v.Timbre,
				ScriptTag:   v.ScriptTag,
				Sequence:    k,
			}
		}
		if err := tx.CreateInBatches(roomScriptList, 100).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return roomInfo.RoomId, nil

}

func RoomList(ctx context.Context, userId int64, page dto.PageParam) ([]*table.RoomTable, int64, error) {
	var list []*table.RoomTable
	err := qdb.Db.WithContext(ctx).
		Model(&table.RoomTable{}).
		Where("user_id = ?", userId).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).
		Scan(&list).
		Error
	if err != nil {
		qLog.Error("find room list error", zap.Int64("userId", userId), zap.Error(err))
		return nil, 0, err
	}
	var count int64
	err = qdb.Db.WithContext(ctx).Model(&table.RoomTable{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

type RoomDetail struct {
	RoomId     int64         `json:"room_id"`
	RoomName   string        `json:"room_name"`
	ProductTag string        `json:"product_tag"`
	Timbre     string        `json:"timbre"`
	TemplateId int           `json:"template_id"`
	Scripts    []*ScriptItem `json:"scripts"`
}

type ScriptItem struct {
	ScriptId    int64        `json:"script_id"`
	ScriptTitle string       `json:"script_title"`
	Scenes      []*SceneItem `json:"scenes"`
}

type SceneItem struct {
	SceneId   int64  `json:"scene_id"`
	SceneName string `json:"scene_name"`
	Audio     string `json:"audio"`
}

func RoomInfo(ctx context.Context, userId int64, roomId int64) (*RoomDetail, error) {
	var roomInfo table.RoomTable
	err := qdb.Db.WithContext(ctx).
		Model(&table.RoomTable{}).
		Where("room_id = ?", roomId).
		Find(&roomInfo).
		Error
	if err != nil {
		return nil, err
	}
	if roomInfo.UserId != userId {
		return nil, errors.New("无权限访问此直播间")
	}

	roomDetail := &RoomDetail{
		RoomId:     roomInfo.RoomId,
		RoomName:   roomInfo.RoomName,
		ProductTag: roomInfo.ProductTag,
		Timbre:     roomInfo.Timbre,
		TemplateId: roomInfo.TemplateId,
		Scripts:    nil,
	}

	var roomScript []*table.RoomScriptTable
	err = qdb.Db.WithContext(ctx).
		Model(&table.RoomScriptTable{}).
		Where("room_id = ?", roomId).
		Order("sequence").
		Scan(&roomScript).
		Error
	if err != nil {
		return nil, err
	}

	scripts := make([]*ScriptItem, len(roomScript))
	for index, script := range roomScript {
		var list []*table.SceneTable
		err = qdb.Db.WithContext(ctx).Model(&table.SceneTable{}).
			Where("script_id = ?", script.ScriptId).
			Scan(&list).
			Error
		if err != nil {
			return nil, err
		}
		scenes := make([]*SceneItem, len(list))
		for k, v := range list {
			scenes[k] = &SceneItem{
				SceneId:   v.SceneId,
				SceneName: v.SceneName,
				Audio:     v.Audio,
			}
		}
		scripts[index] = &ScriptItem{
			ScriptId:    script.ScriptId,
			ScriptTitle: script.ScriptTitle,
			Scenes:      scenes,
		}
	}
	roomDetail.Scripts = scripts

	return roomDetail, nil
}

func RoomDelete(ctx context.Context, userId int64, roomId int64) error {
	room := table.RoomTable{
		RoomId: roomId,
		UserId: userId,
	}
	return qdb.Db.WithContext(ctx).Delete(room).Error
}
