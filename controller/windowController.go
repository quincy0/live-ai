package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quincy0/live-ai/dto"
	"github.com/quincy0/live-ai/service/douyinService"
	"github.com/quincy0/live-ai/util"
	"github.com/quincy0/qpro/app"
	"github.com/quincy0/qpro/qLog"
)

func GetWindow(c *gin.Context) {

	user := util.ParseUser(c)
	qLog.Info(fmt.Sprintf("AddWindow userId: %d", user.UserId))

	//ctx := c.Request.Context()
	var params dto.WindowGetParam
	err := c.ShouldBindJSON(&params)

	result, err := douyinService.GetWindow(params.DouyinId)
	if err != nil {
		app.Error(c, 100410, err)
		return
	}
	app.OK(c, result)
}

func GetWindowFromDB(c *gin.Context) {

	user := util.ParseUser(c)
	qLog.Info(fmt.Sprintf("GetWindowFromDB userId: %d", user.UserId))
	ctx := c.Request.Context()

	//ctx := c.Request.Context()
	var params dto.WindowGetParam
	err := c.ShouldBindJSON(&params)

	result, err := douyinService.GetWindowFromDB(ctx, params.DouyinId)
	if err != nil {
		app.Error(c, 100410, err)
		return
	}
	app.OK(c, result)
}

func AddWindow(c *gin.Context) {

	user := util.ParseUser(c)
	qLog.Info(fmt.Sprintf("AddWindow userId: %d", user.UserId))

	//ctx := c.Request.Context()
	var params dto.WindowAddParam
	err := c.ShouldBindJSON(&params)

	result, err := douyinService.AddWindow(params.DouyinId, params.ProductId)
	if err != nil {
		app.Error(c, 100410, err)
		return
	}
	app.OK(c, result)
}

func DeleteWindow(c *gin.Context) {

	user := util.ParseUser(c)
	qLog.Info(fmt.Sprintf("DeleteWindow userId: %d", user.UserId))

	//ctx := c.Request.Context()
	var params dto.WindowAddParam
	err := c.ShouldBindJSON(&params)

	result, err := douyinService.DeleteWindow(params.DouyinId, params.ProductId)
	if err != nil {
		app.Error(c, 100410, err)
		return
	}
	app.OK(c, result)
}

func UpdateWindow(c *gin.Context) {

	user := util.ParseUser(c)
	qLog.Info(fmt.Sprintf("UpdateWindow userId: %d", user.UserId))
	ctx := c.Request.Context()

	//ctx := c.Request.Context()
	var params dto.WindowUpdateParam
	err := c.ShouldBindJSON(&params)

	result, err := douyinService.UpdateWindow(ctx, params)
	if err != nil {
		app.Error(c, 100410, err)
		return
	}
	app.OK(c, result)
}
