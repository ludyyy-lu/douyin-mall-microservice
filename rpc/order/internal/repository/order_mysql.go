package repository

import (
	"douyin-mall/rpc/order/internal/repository/model"
	"gorm.io/gorm"
)

type OrderRepo struct {
	*gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	return &OrderRepo{
		db,
	}
}
func (r *OrderRepo) CreateOrder(order model.Order) (string, error) {
	if r.Create(&order).Error != nil {
		return "", r.Error
	}
	return order.OrderID, nil

}
