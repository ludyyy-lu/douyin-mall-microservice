package repository

import (
	"gorm.io/gorm"
)

// Order 数据库模型
type Order struct {
	gorm.Model
	OrderID      string      `gorm:"type:varchar(100);uniqueIndex"`
	UserID       uint32      `gorm:"type:int(11)"`
	UserCurrency string      `gorm:"type:varchar(10)"`
	Email        string      `gorm:"column:email"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderID;references:OrderID"`
	Address      Address     `gorm:"embedded"`
}

// OrderItem 订单项
type OrderItem struct {
	ID      uint32
	OrderID string  `gorm:"type:varchar(100)"`
	ItemID  uint32  `gorm:"column:item_id"`
	Cost    float32 `gorm:"column:cost"`
}

// Address 订单地址
type Address struct {
	Email         string `gorm:"type:varchar(100)"`
	StreetAddress string `gorm:"column:street_address"`
	City          string `gorm:"column:city"`
	State         string `gorm:"column:state"`
	Country       string `gorm:"column:country"`
	ZipCode       int32  `gorm:"column:zip_code"`
}
