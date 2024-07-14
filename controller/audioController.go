package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/service/audioService"
	"github.com/quincy0/live-ai/service/goodsService"
	"github.com/quincy0/live-ai/tts"
	"github.com/quincy0/live-ai/util"
	"github.com/quincy0/qpro/app"
	"github.com/quincy0/qpro/qLog"
	"go.uber.org/zap"
)

func AudioCreate(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto.AudioCreateParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}

	ctxNew := util.InitContextWithSameTrace(ctx)
	url, err := tts.CreateChatAudioV2(ctxNew, params.Recreate, params.Spk, params.Text)
	if err != nil {
		app.Error(c, 800001, err)
		return
	}
	if len(url) == 0 {
		url = "音频生成中"
	}
	app.OK(c, url)
}

func AudioNotify(c *gin.Context) {
	ctx := c.Request.Context()
	sum := c.Param("sum")
	var params dto.AudioNotifyParam
	err := c.ShouldBindQuery(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	err = tts.CreateChatAudioNotify(ctx, sum, params)
	if err != nil {
		app.Error(c, 800003, err)
		return
	}
	app.OK(c, "success")
}

func AudioList(c *gin.Context) {
	ctx := c.Request.Context()

	var params dto.AudioListParam
	err := c.ShouldBindQuery(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	roomData, err := audioService.AudioList(ctx, params.RoomId)
	if err != nil {
		app.Error(c, 100002, err)
		return
	}
	app.OK(c, roomData)

}

func GoodsList(c *gin.Context) {
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
	list, count := goodsService.List(ctx, params)
	app.PageOK(c, list, count, params.PageNum, params.PageSize)
}
