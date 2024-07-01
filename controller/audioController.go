package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/quincy0/live-ai/service/audioService"
	"github.com/quincy0/live-ai/tts"
	"github.com/quincy0/live-ai/util"
	"github.com/quincy0/qpro/app"
	"github.com/quincy0/qpro/qLog"
	"go.uber.org/zap"
)

func AudioCreate(c *gin.Context) {
	ctx := c.Request.Context()
	var params AudioCreateParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}

	ctxNew := util.InitContextWithSameTrace(ctx)
	url, err := tts.CreateChatAudio(ctxNew, params.Recreate, params.Spk, params.Text)
	if err != nil {
		app.Error(c, 800001, err)
		return
	}
	app.OK(c, url)
}

func AudioList(c *gin.Context) {
	ctx := c.Request.Context()

	var params AudioListParam
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
