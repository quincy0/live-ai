package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quincy0/live-ai/consts"
	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/service/contentService"
	"github.com/quincy0/live-ai/service/douyinService"
	"github.com/quincy0/live-ai/service/roomService"
	"github.com/quincy0/live-ai/service/userService"
	"github.com/quincy0/live-ai/util"
	"github.com/quincy0/qpro/app"
	"github.com/quincy0/qpro/qLog"
	"go.uber.org/zap"
)

func ProductTagList(c *gin.Context) {
	app.OK(c, consts.ProductTagList)
}

func ScriptTagList(c *gin.Context) {
	app.OK(c, consts.ScriptTagList)
}

func RoomTemplateList(c *gin.Context) {
	app.OK(c, consts.RoomTemplateList)
}

func ChapterCreate(c *gin.Context) {
	user := util.ParseUser(c)
	ctx := c.Request.Context()
	var params dto.ChapterCreateParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	chapterId, err := contentService.ChapterCreate(ctx, user.UserId, &params)
	if err != nil {
		qLog.TraceError(ctx, "chapter create failed", zap.Error(err))
		app.Error(c, 400000, err)
		return
	}
	res := map[string]int64{"chapterId": chapterId}
	app.OK(c, res)
}

func ChapterList(c *gin.Context) {
	user := util.ParseUser(c)
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
	list, count, err := contentService.ChapterList(ctx, user.UserId, params)
	if err != nil {
		qLog.TraceError(ctx, "chapter list failed", zap.Error(err))
		app.Error(c, 400001, err)
		return
	}
	app.PageOK(c, list, count, params.PageNum, params.PageSize)
}

func ChapterInfo(c *gin.Context) {
	user := util.ParseUser(c)
	ctx := c.Request.Context()
	var params dto.ChapterInfoParam
	err := c.ShouldBindQuery(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	detail, err := contentService.ChapterInfo(ctx, user.UserId, params.ChapterId)
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
	user := util.ParseUser(c)
	ctx := c.Request.Context()
	var params dto.ScriptCreateParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	scriptId, err := contentService.ScriptCreate(ctx, user.UserId, &params)
	if err != nil {
		qLog.TraceError(ctx, "chapter create failed", zap.Error(err))
		app.Error(c, 400100, err)
		return
	}
	res := map[string]int64{"scriptId": scriptId}
	app.OK(c, res)
}

func ScriptList(c *gin.Context) {
	user := util.ParseUser(c)
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
	list, count, err := contentService.ScriptList(ctx, user.UserId, params)
	if err != nil {
		qLog.TraceError(ctx, "script list failed", zap.Error(err))
		app.Error(c, 400101, err)
		return
	}
	app.PageOK(c, list, count, params.PageNum, params.PageSize)
}

func ScriptInfo(c *gin.Context) {
	user := util.ParseUser(c)
	ctx := c.Request.Context()
	var params dto.ScriptInfoParam
	err := c.ShouldBindQuery(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	detail, err := contentService.ScriptInfo(ctx, user.UserId, params.ScriptId)
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

func RoomCreate(c *gin.Context) {
	user := util.ParseUser(c)
	ctx := c.Request.Context()
	var params dto.RoomCreateParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	roomId, err := roomService.RoomCreate(ctx, user.UserId, params)
	if err != nil {
		qLog.TraceError(ctx, "create room failed", zap.Error(err))
		app.Error(c, 700000, err)
		return
	}
	app.OK(c, map[string]int64{"roomId": roomId})
}

func RoomList(c *gin.Context) {
	user := util.ParseUser(c)
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
	list, count, err := roomService.RoomList(ctx, user.UserId, params)
	if err != nil {
		qLog.TraceError(ctx, "get room list failed", zap.Error(err))
		app.Error(c, 700001, err)
		return
	}
	app.PageOK(c, list, count, params.PageNum, params.PageSize)
}

func RoomInfo(c *gin.Context) {
	user := util.ParseUser(c)
	ctx := c.Request.Context()
	var params dto.RoomInfoParam
	err := c.ShouldBindQuery(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	roomInfo, err := roomService.RoomInfo(ctx, user.UserId, params.RoomId)
	if err != nil {
		qLog.TraceError(ctx, "get room info failed", zap.Error(err))
		app.Error(c, 700002, err)
		return
	}
	app.OK(c, roomInfo)
}

func Hello(c *gin.Context) {
	//ctx := c.Request.Context()
	user := util.ParseUser(c)
	app.OK(c, user)
}

func Register(c *gin.Context) {
	ctx := c.Request.Context()
	var params dto.LoginParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	userId, err := userService.Register(ctx, params.Username, params.Password)
	if err != nil {
		qLog.TraceError(ctx, "register failed", zap.Error(err))
		app.Error(c, 100005, err)
		return
	}
	app.OK(c, map[string]int64{"userId": userId})
}

func TimbreCreate(c *gin.Context) {
	user := util.ParseUser(c)
	ctx := c.Request.Context()
	var params dto.TimbreParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}
	err = contentService.TimbreCreate(ctx, user.UserId, params)
	if err != nil {
		app.Error(c, 100409, err)
		return
	}
	app.OK(c, "添加成功")
}

func TimbreList(c *gin.Context) {
	user := util.ParseUser(c)
	ctx := c.Request.Context()
	list, err := contentService.TimbreList(ctx, user.UserId)
	if err != nil {
		app.Error(c, 100410, err)
		return
	}
	app.OK(c, list)
}

func AddCard(c *gin.Context) {
	ctx := c.Request.Context()
	user := util.ParseUser(c)

	// 打印 UserId
	qLog.Info(fmt.Sprintf("begin AddCard %d", user.UserId))

	var params dto.DouyinAddCard
	err := c.ShouldBindJSON(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}

	qLog.Info(fmt.Sprintf("DeleteCard DeleteCard DouyinId %s", params.DouyinId))
	//ctx := c.Request.Context()
	result, err := douyinService.AddCard(params.DouyinId, params.ProductId)
	if err != nil {
		app.Error(c, 100410, err)
		return
	}
	app.OK(c, result)
}

func DeleteCard(c *gin.Context) {
	ctx := c.Request.Context()
	user := util.ParseUser(c)
	// 打印 UserId
	qLog.Info(fmt.Sprintf("DeleteCard DeleteCard %d", user.UserId))

	var params dto.DouyinDeleteCard
	err := c.ShouldBindJSON(&params)
	if err != nil {
		qLog.TraceError(ctx, "get params failed", zap.Error(err))
		app.Error(c, 100000, err)
		return
	}

	qLog.Info(fmt.Sprintf("DeleteCard DeleteCard DouyinId %s", params.DouyinId))

	//ctx := c.Request.Context()
	result, err := douyinService.DeleteCard(params.DouyinId, params.ProductId)
	if err != nil {
		app.Error(c, 100410, err)
		return
	}
	app.OK(c, result)
}
