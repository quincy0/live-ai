package douyinService

import (
	"context"
	"errors"
	"github.com/quincy0/live-ai/table"
	"github.com/quincy0/qpro/qdb"
	"gorm.io/gorm"
)

func GetGoodsWindowCouponStrategyFromDB(ctx context.Context, douyinId string) ([]table.GoodsWindowCouponStrategy, error) {
	var goodsWindowCouponStrategy []table.GoodsWindowCouponStrategy

	// 根据 douyinId 查询 GoodsWindow 表
	err := qdb.Db.WithContext(ctx).
		Model(&table.GoodsWindowCouponStrategy{}).
		Where("account_id = ?", douyinId).
		Find(&goodsWindowCouponStrategy).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("no records found")
		}
		return nil, err
	}

	// 返回查询到的 GoodsWindow 切片
	return goodsWindowCouponStrategy, nil
}

func AddGoodsWindowCouponStrategyToDB(ctx context.Context, goodsWindowCouponStrategy table.GoodsWindowCouponStrategy) (*table.GoodsWindowCouponStrategy, error) {
	// 先检查数据库中是否已经存在该记录
	var existingStrategy table.GoodsWindowCouponStrategy
	err := qdb.Db.WithContext(ctx).
		Model(&table.GoodsWindowCouponStrategy{}).
		Where("account_id = ?", goodsWindowCouponStrategy.AccountID).
		First(&existingStrategy).Error

	if err != nil {
		// 如果记录不存在，或者查询出现错误
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		// 如果记录不存在，则执行插入操作
	} else {
		// 记录已经存在
		return nil, errors.New("record already exists")
	}

	// 将新记录插入数据库
	err = qdb.Db.WithContext(ctx).
		Create(&goodsWindowCouponStrategy).Error

	if err != nil {
		return nil, err
	}

	// 返回新添加的 goodsWindowCouponStrategy 实例
	return &goodsWindowCouponStrategy, nil
}

func UpdateGoodsWindowCouponStrategyToDB(ctx context.Context, goodsWindowCouponStrategy table.GoodsWindowCouponStrategy) (*table.GoodsWindowCouponStrategy, error) {

	var goodsWindowCouponStrategyTemp table.GoodsWindowCouponStrategy

	// 查询数据库是否存在该记录
	// 查询数据库是否存在该记录，同时根据 product_id 和 douyin_id 进行查询
	err := qdb.Db.WithContext(ctx).
		Model(&table.GoodsWindowCouponStrategy{}).
		Where("id = ?", goodsWindowCouponStrategy.ID).
		First(&goodsWindowCouponStrategyTemp).Error

	if err != nil {
		// 记录不存在，或者查询出现错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	// 如果记录存在，更新 credit 和 totalAmount
	err = qdb.Db.WithContext(ctx).
		Model(&table.GoodsWindowCouponStrategy{}).
		Where("id = ?", goodsWindowCouponStrategyTemp.ID).
		Updates(table.GoodsWindowCouponStrategy{
			CreditStrategy:      goodsWindowCouponStrategy.CreditStrategy,
			TotalAmountStrategy: goodsWindowCouponStrategy.TotalAmountStrategy,
		}).Error

	if err != nil {
		return nil, err
	}

	// 更新后，再次查询数据库获取最新的 goodsWindow 数据
	err = qdb.Db.WithContext(ctx).
		Model(&table.GoodsWindowCouponStrategy{}).
		Where("id = ?", goodsWindowCouponStrategyTemp.ID).
		First(&goodsWindowCouponStrategyTemp).Error

	if err != nil {
		return nil, err
	}

	// 返回更新后的 goodsWindow 实例
	return &goodsWindowCouponStrategyTemp, nil
}

func DeleteGoodsWindowCouponStrategyToDB(ctx context.Context, goodsWindowCouponStrategy table.GoodsWindowCouponStrategy) (error, error) {
	// 先检查数据库中是否存在该记录
	var goodsWindowCouponStrategyTemp table.GoodsWindowCouponStrategy
	err := qdb.Db.WithContext(ctx).
		Model(&table.GoodsWindowCouponStrategy{}).
		Where("id = ?", goodsWindowCouponStrategy.ID).
		First(&goodsWindowCouponStrategyTemp).Error

	if err != nil {
		// 如果记录不存在，返回错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("record not found"), nil
		}
		// 查询发生其他错误时，返回错误信息
		return err, nil
	}

	// 如果找到记录，执行删除操作
	err = qdb.Db.WithContext(ctx).
		Where("id = ?", goodsWindowCouponStrategyTemp.ID).
		Delete(&goodsWindowCouponStrategyTemp).Error

	if err != nil {
		// 删除失败，返回错误
		return err, nil
	}

	// 删除成功，返回 nil 表示无错误
	return nil, nil
}
