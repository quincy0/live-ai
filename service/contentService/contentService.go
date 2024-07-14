package contentService

import (
	"context"
	"errors"

	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/table"
	"github.com/quincy0/qpro/qdb"
	"gorm.io/gorm"
)

type ChapterDetail struct {
	*table.ChapterTable
	List []*table.ParagraphTable `json:"list"`
}

type ScriptDetail struct {
	*table.ScriptTable
	List []*table.SceneTable `json:"list"`
}

func ChapterCreate(ctx context.Context, userId int64, params *dto.ChapterCreateParam) (int64, error) {
	chapter := &table.ChapterTable{
		ChapterTitle: params.ChapterTitle,
		UserId:       userId,
		ScriptTag:    params.ScriptTag,
		ProductTag:   params.ProductTag,
		Summary:      params.Summary,
	}
	paragraphList := make([]*table.ParagraphTable, len(params.Paragraph))
	for k, v := range params.Paragraph {
		paragraphList[k] = &table.ParagraphTable{
			ParagraphId:    0,
			ChapterId:      0,
			ParagraphTitle: v.ParagraphTitle,
			Content:        v.Content,
		}
	}
	err := qdb.Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(chapter).Error
		if err != nil {
			return err
		}
		for _, v := range paragraphList {
			v.ChapterId = chapter.ChapterId
		}
		err = tx.CreateInBatches(paragraphList, 100).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return chapter.ChapterId, nil
}

func ChapterList(ctx context.Context, userId int64, page dto.PageParam) ([]*table.ChapterTable, int64, error) {
	var list []*table.ChapterTable
	err := qdb.Db.WithContext(ctx).Model(&table.ChapterTable{}).
		Where("user_id = ?", userId).
		Order("chapter_id").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).
		Scan(&list).
		Error
	if err != nil {
		return nil, 0, err
	}
	var count int64
	err = qdb.Db.WithContext(ctx).Model(&table.ChapterTable{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func ChapterInfo(ctx context.Context, userId int64, chapterId int64) (*ChapterDetail, error) {
	var chapter table.ChapterTable
	err := qdb.Db.WithContext(ctx).Model(&table.ChapterTable{}).
		Where("chapter_id = ?", chapterId).
		Find(&chapter).
		Error
	if err != nil {
		return nil, err
	}
	if chapter.UserId != userId {
		return nil, errors.New("无权访问该文章")
	}
	var paragraphs []*table.ParagraphTable
	err = qdb.Db.WithContext(ctx).
		Model(&table.ParagraphTable{}).
		Where("chapter_id = ?", chapterId).
		Scan(&paragraphs).
		Error
	if err != nil {
		return nil, err
	}
	return &ChapterDetail{
		ChapterTable: &chapter,
		List:         paragraphs,
	}, nil
}

func ParagraphEdit(ctx context.Context, params dto.ParagraphEditParam) error {
	return qdb.Db.WithContext(ctx).
		Model(&table.ParagraphTable{}).
		Where("paragraph_id = ?", params.ParagraphId).
		Select("paragraph_title", "content").
		Updates(map[string]interface{}{"paragraph_title": params.ParagraphTitle, "content": params.Content}).
		Error
}

func ScriptCreate(ctx context.Context, userId int64, params *dto.ScriptCreateParam) (int64, error) {
	chapter, err := ChapterInfo(ctx, userId, params.ChapterId)
	if err != nil {
		return 0, err
	}
	script := &table.ScriptTable{
		ScriptId:    0,
		UserId:      userId,
		ScriptTitle: params.ScriptTitle,
		ScriptTag:   chapter.ScriptTag,
		ProductTag:  chapter.ProductTag,
		Summary:     chapter.Summary,
		Timbre:      params.Timbre,
		LastPlay:    0,
	}
	list := make([]*table.SceneTable, len(chapter.List))
	for k, v := range chapter.List {
		list[k] = &table.SceneTable{
			SceneId:   0,
			ScriptId:  0,
			SceneName: v.ParagraphTitle,
			Content:   v.Content,
			Audio:     "",
		}
	}

	err = qdb.Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Create(script).Error
		if err != nil {
			return err
		}
		for _, v := range list {
			v.ScriptId = script.ScriptId
		}
		err = tx.CreateInBatches(list, 100).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return script.ScriptId, nil
}

func ScriptList(ctx context.Context, userId int64, page dto.PageParam) ([]*table.ScriptTable, int64, error) {
	var list []*table.ScriptTable
	err := qdb.Db.WithContext(ctx).Model(&table.ScriptTable{}).
		Where("user_id = ?", userId).
		Order("script_id").
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).
		Scan(&list).
		Error
	if err != nil {
		return nil, 0, err
	}
	var count int64
	err = qdb.Db.WithContext(ctx).Model(&table.ScriptTable{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}

func ScriptInfo(ctx context.Context, userId int64, scriptId int64) (*ScriptDetail, error) {
	var script table.ScriptTable
	err := qdb.Db.WithContext(ctx).Model(&table.ScriptTable{}).
		Where("script_id = ?", scriptId).
		Find(&script).
		Error
	if err != nil {
		return nil, err
	}
	if script.UserId != userId {
		return nil, errors.New("无权访问该剧本")
	}
	var list []*table.SceneTable
	err = qdb.Db.WithContext(ctx).Model(&table.SceneTable{}).
		Where("script_id = ?", scriptId).
		Scan(&list).
		Error
	if err != nil {
		return nil, err
	}
	return &ScriptDetail{
		ScriptTable: &script,
		List:        list,
	}, nil
}

func SceneEdit(ctx context.Context, sceneId int64, audioURL string) error {
	return qdb.Db.WithContext(ctx).
		Model(&table.SceneTable{}).
		Where("scene_id = ?", sceneId).
		Update("audio", audioURL).
		Error
}
