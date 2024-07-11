package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quincy0/live-ai/consts"
	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/service/contentService"
	"github.com/quincy0/qpro/app"
	"github.com/quincy0/qpro/qLog"
	"go.uber.org/zap"
)

func ProductTagList(c *gin.Context) {
	app.OK(c, consts.ProductTagList)
}

func ChapterCreate(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto.ChapterCreateParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	chapterId, err := contentService.ChapterCreate(ctx, &params)
	if err != nil {
		qLog.TraceError(ctx, "chapter create failed", zap.Error(err))
		app.Error(c, 400000, err)
		return
	}
	res := map[string]int64{"chapterId": chapterId}
	app.OK(c, res)
}

func ChapterList(c *gin.Context) {
	ctx := c.Request.Context()
	var params = dto.PageParam{
		PageSize: 100,
		PageNum:  1,
	}
	err := c.BindQuery(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	list, count, err := contentService.ChapterList(ctx, params)
	if err != nil {
		qLog.TraceError(ctx, "chapter list failed", zap.Error(err))
		app.Error(c, 400001, err)
		return
	}
	app.PageOK(c, list, count, params.PageNum, params.PageSize)
}

func ChapterInfo(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto.ChapterInfoParam
	err := c.ShouldBindQuery(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	detail, err := contentService.ChapterInfo(ctx, params.ChapterId)
	if err != nil {
		qLog.TraceError(ctx, "get chapter failed", zap.Error(err))
		app.Error(c, 400002, err)
		return
	}
	app.OK(c, detail)
}

func ParagraphEdit(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto.ParagraphEditParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	err = contentService.ParagraphEdit(ctx, params)
	if err != nil {
		qLog.TraceError(ctx, "paragraph edit failed", zap.Error(err))
		app.Error(c, 400003, err)
		return
	}
	app.OK(c, "success")
}

func ScriptCreate(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto.ScriptCreateParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	scriptId, err := contentService.ScriptCreate(ctx, &params)
	if err != nil {
		qLog.TraceError(ctx, "chapter create failed", zap.Error(err))
		app.Error(c, 400100, err)
		return
	}
	res := map[string]int64{"scriptId": scriptId}
	app.OK(c, res)
}

func ScriptList(c *gin.Context) {
	ctx := c.Request.Context()
	var params = dto.PageParam{
		PageSize: 100,
		PageNum:  1,
	}
	err := c.BindQuery(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	list, count, err := contentService.ScriptList(ctx, params)
	if err != nil {
		qLog.TraceError(ctx, "script list failed", zap.Error(err))
		app.Error(c, 400101, err)
		return
	}
	app.PageOK(c, list, count, params.PageNum, params.PageSize)
}

func ScriptInfo(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto.ScriptInfoParam
	err := c.ShouldBindQuery(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	detail, err := contentService.ScriptInfo(ctx, params.ScriptId)
	if err != nil {
		qLog.TraceError(ctx, "get script failed", zap.Error(err))
		app.Error(c, 400102, err)
		return
	}
	app.OK(c, detail)
}

func SceneEdit(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto.SceneEditParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	err = contentService.SceneEdit(ctx, params.SceneId, params.Audio)
	if err != nil {
		qLog.TraceError(ctx, "scene edit failed", zap.Error(err))
		app.Error(c, 400103, err)
		return
	}
	app.OK(c, "success")
}
