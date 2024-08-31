package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/service/douyinService"
	"github.com/quincy0/live-ai/util"
	"github.com/quincy0/qpro/app"
	"github.com/quincy0/qpro/qLog"
	"go.uber.org/zap"
)

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
