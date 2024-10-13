package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/quincy0/live-ai/service/douyinService"
	"github.com/quincy0/live-ai/table"
	"github.com/quincy0/live-ai/util"
	"github.com/quincy0/qpro/app"
	"github.com/quincy0/qpro/qLog"
)

func GetGoodsWindowCouponStrategyFromDB(c *gin.Context) {

	user := util.ParseUser(c)
	qLog.Info(fmt.Sprintf("GetGoodsWindowCouponStrategyFromDB userId: %d", user.UserId))
	ctx := c.Request.Context()

	var params table.GoodsWindowCouponStrategy
	err := c.ShouldBindJSON(&params)

	result, err := douyinService.GetGoodsWindowCouponStrategyFromDB(ctx, params.AccountID)
	if err != nil {
		app.Error(c, 100410, err)
		return
	}
	app.OK(c, result)
}

func AddGoodsWindowCouponStrategyToDB(c *gin.Context) {

	user := util.ParseUser(c)
	qLog.Info(fmt.Sprintf("AddGoodsWindowCouponStrategy userId: %d", user.UserId))

	var params table.GoodsWindowCouponStrategy
	err := c.ShouldBindJSON(&params)

	ctx := c.Request.Context()

	result, err := douyinService.AddGoodsWindowCouponStrategyToDB(ctx, params)
	if err != nil {
		app.Error(c, 100410, err)
		return
	}
	app.OK(c, result)
}

func DeleteGoodsWindowCouponStrategy(c *gin.Context) {

	user := util.ParseUser(c)
	qLog.Info(fmt.Sprintf("DeleteGoodsWindowCouponStrategy userId: %d", user.UserId))

	var params table.GoodsWindowCouponStrategy
	err := c.ShouldBindJSON(&params)

	qLog.Info(fmt.Sprintf("DeleteGoodsWindowCouponStrategy params: %s", params))

	ctx := c.Request.Context()

	result, err := douyinService.DeleteGoodsWindowCouponStrategyToDB(ctx, params)
	if err != nil {
		app.Error(c, 100410, err)
		return
	}
	app.OK(c, result)
}

func UpdateGoodsWindowCouponStrategy(c *gin.Context) {

	user := util.ParseUser(c)
	qLog.Info(fmt.Sprintf("UpdateGoodsWindowCouponStrategy userId: %d", user.UserId))

	var params table.GoodsWindowCouponStrategy
	err := c.ShouldBindJSON(&params)
	ctx := c.Request.Context()

	result, err := douyinService.UpdateGoodsWindowCouponStrategyToDB(ctx, params)
	if err != nil {
		app.Error(c, 100410, err)
		return
	}
	app.OK(c, result)
}
