package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/service/audioService"
	"github.com/quincy0/live-ai/service/goodsService"
	"github.com/quincy0/live-ai/service/scriptService"
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
	url, err := tts.CreateChatAudio(ctxNew, params.Recreate, params.Spk, params.Text)
	if err != nil {
		app.Error(c, 800001, err)
		return
	}
	app.OK(c, url)
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
	list, count, err := scriptService.List(ctx, params)
	app.PageOK(c, list, count, params.PageNum, params.PageSize)
}

func ScriptUpsert(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto.CreateScriptParam
	err := c.ShouldBindQuery(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	if params.GoodsId == 0 && params.ScriptId == 0 {
		qLog.TraceError(ctx, "参数错误")
		app.Error(c, 100001, errors.New("商品不存在"))
		return
	}
	scene, err := scriptService.Edit(ctx, params)
	if err != nil {
		qLog.TraceError(ctx, "剧本更新失败")
		app.Error(c, 200000, err)
		return
	}
	app.OK(c, scene)
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
	info, err := scriptService.Info(ctx, params.ScriptId)
	if err != nil {
		qLog.TraceError(ctx, "get script info failed", zap.Int64("scriptId", params.ScriptId), zap.Error(err))
		app.Error(c, 200002, err)
		return
	}
	app.OK(c, info)
}

/*
1. 创建剧本
GET  script/create?goodsId=g2332 返回 scriptId
2. 上传音频
POST audio/upload  {"scriptId": 111, "audio": "https://muqi...."}  返回success
3. 获取剧本信息
GET  script/info?scriptId=111  返回剧本音频列表
4. 创建直播间
POST room/create [scriptId-111,scriptId-222,scriptId-333] 返回 success roomId
5. 获取直播间列表
GET  room/list   返回[roomId1, roomId2]
6. 获取直播间音频列表
GET  room/info  返回[{scriptId,goodsId, [audio1, audio2]},{scriptId,goodsId, [audio1, audio2]}]
*/
