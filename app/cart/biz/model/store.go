// 定义数据库操作接口（新建文件 model/store.go）
package model

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type CartStore interface {
	GetCartByUserId(db *gorm.DB, ctx context.Context, userId uint32) ([]*Cart, error)
	EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error
	AddItem(db *gorm.DB, ctx context.Context, c *Cart) error
}

type CartStoreImpl struct{}

func (s *CartStoreImpl) GetCartByUserId(db *gorm.DB, ctx context.Context, userId uint32) (cartList []*Cart, err error) {
	err = db.Debug().WithContext(ctx).Model(&Cart{}).Find(&cartList, "user_id = ?", userId).Error
	return cartList, err
}

func (s *CartStoreImpl) EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	// 原清空购物车逻辑
	if userId == 0 {
		return errors.New("user id is required")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id =?", userId).Error
}
func (s *CartStoreImpl) AddItem(db *gorm.DB, ctx context.Context, c *Cart) error {
	var find Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: c.UserID, ProductID: c.ProductID}).First(&find).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if find.ID != 0 {
		err = db.WithContext(ctx).Model(&Cart{}).Where(&Cart{UserID: c.UserID, ProductID: c.ProductID}).UpdateColumn("qty", gorm.Expr("qty+?", c.Qty)).Error
	} else {
		err = db.WithContext(ctx).Model(&Cart{}).Create(c).Error
	}
	return err
}
