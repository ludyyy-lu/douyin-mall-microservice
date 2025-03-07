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
	OrderItems   []OrderItem `gorm:"dforeignKey:OrderID;references:OrderID"`
	Email        string      `gorm:"type:varchar(100)"`
	Address      Address     `gorm:"embedded"`
	Paid         bool
	Expired      bool
}

// Address 订单地址
type Address struct {
	StreetAddress string `gorm:"column:street_address"`
	City          string `gorm:"column:city"`
	State         string `gorm:"column:state"`
	Country       string `gorm:"column:country"`
	ZipCode       int32  `gorm:"column:zip_code"`
}
