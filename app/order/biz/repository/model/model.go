package model

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
	gorm.Model
	OrderID   string  `gorm:"type:varchar(100)"`
	ProductID uint32  `gorm:"type:int(11)"`
	Quantity  int32   `gorm:"type:int(11)"`
	Cost      float32 `gorm:"type:decimal(10,2)"`
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
