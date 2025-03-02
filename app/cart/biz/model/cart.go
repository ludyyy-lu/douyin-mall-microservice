package model

import (
	"context"
	"database/sql"
	"errors"

	//"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product/productcatalogservice"

	"gorm.io/gorm"
)

type CartStore interface {
	AddItem(ctx context.Context, db *sql.DB, item *Cart) error
	EmptyCart(ctx context.Context, db *sql.DB, userID int64) error
	GetCartByUserId(ctx context.Context, db *sql.DB, userID int64) ([]*Cart, error)
}

type Cart struct {
	gorm.Model
	UserID    uint32 `gorm:"type:int(11);not null;index:idx_user_id"`
	ProductID uint32 `gorm:"type:int(11);not null"`
	Qty       uint32 `gorm:"type:int(11);not null"`
}

func (Cart) TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, cart *Cart) error {
	var row Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserID: cart.UserID, ProductID: cart.ProductID}).
		First(&row).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if row.ID > 0 {
		return db.WithContext(ctx).
			Model(&Cart{}).
			Where(&Cart{UserID: cart.UserID, ProductID: cart.ProductID}).
			UpdateColumn("qty", gorm.Expr("qty+?", cart.Qty)).Error
	}
	return db.WithContext(ctx).Create(cart).Error
}

func EmptyCart(ctx context.Context, db *gorm.DB, userId uint32) error {
	if userId == 0 {
		return errors.New("user id is required")
	}
	return db.WithContext(ctx).Delete(&Cart{}, "user_id =?", userId).Error
}

func GetCartByUserId(ctx context.Context, db *gorm.DB, userId uint32) ([]*Cart, error) {
	var rows []*Cart
	err := db.WithContext(ctx).
		Model(&Cart{}).
		Where(&Cart{UserID: userId}).
		Find(&rows).Error
	return rows, err
}
