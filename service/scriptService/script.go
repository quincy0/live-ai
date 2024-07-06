package scriptService

import (
	"context"
	"fmt"
	"time"

	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/table"
	"github.com/quincy0/qpro/qdb"
	"gorm.io/gorm"
)

func Edit(ctx context.Context, params dto.CreateScriptParam) (*table.ScriptScene, error) {

	scene := &table.ScriptScene{
		ScriptId: params.ScriptId,
		SceneId:  time.Now().UnixMilli(),
		GoodsId:  params.GoodsId,
		Audio:    params.Audio,
	}

	if params.ScriptId != 0 {
		err := qdb.Db.WithContext(ctx).Create(scene).Error
		if err != nil {
			return nil, err
		}
	} else {
		if params.ScriptId == 0 && len(params.ScriptName) == 0 {
			params.ScriptName = fmt.Sprintf("未命名-%s", time.Now().Format(time.RFC3339))
		}
		script := &table.ScriptInfo{
			ScriptName: params.ScriptName,
			GoodsId:    params.GoodsId,
		}
		err := qdb.Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(script).Error; err != nil {
				return err
			}
			scene.ScriptId = script.ScriptId
			if err := tx.Create(scene).Error; err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return scene, nil
}

type ScriptDetail struct {
	ScriptId   int64        `json:"script_id"`
	ScriptName string       `json:"script_name"`
	GoodsId    int64        `json:"goods_id"`
	Scenes     []*SceneItem `json:"scenes"`
}

type SceneItem struct {
	SceneId int64  `json:"scene_id"`
	Audio   string `json:"audio"`
}

func Info(ctx context.Context, scriptId int64) (*ScriptDetail, error) {
	var script table.ScriptInfo
	var scenes []table.ScriptScene
	err := qdb.Db.WithContext(ctx).Model(&table.ScriptInfo{}).Where("script_id = ?", scriptId).Find(&script).Error
	if err != nil {
		return nil, err
	}
	err = qdb.Db.WithContext(ctx).
		Model(&table.ScriptScene{}).
		Where("script_id = ?", scriptId).
		Limit(200).
		Scan(&scenes).
		Error
	if err != nil {
		return nil, err
	}

	list := make([]*SceneItem, len(scenes))
	for k, v := range scenes {
		list[k] = &SceneItem{
			SceneId: v.SceneId,
			Audio:   v.Audio,
		}
	}
	return &ScriptDetail{
		ScriptId:   script.ScriptId,
		ScriptName: script.ScriptName,
		GoodsId:    script.GoodsId,
		Scenes:     list,
	}, nil
}

func List(ctx context.Context, page dto.PageParam) ([]table.ScriptInfo, int64, error) {
	var list []table.ScriptInfo
	err := qdb.Db.WithContext(ctx).
		Model(&table.ScriptInfo{}).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).
		Scan(&list).
		Error
	if err != nil {
		return nil, 0, err
	}
	var count int64
	err = qdb.Db.WithContext(ctx).
		Model(&table.ScriptInfo{}).
		Count(&count).
		Error
	if err != nil {
		return nil, 0, err
	}

	return list, count, nil
}
