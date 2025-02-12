package core

import (
	"douyin-mall/rpc/order/internal/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitMysql() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: global.Config.Mysql.DSN(),
	}), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	global.DB = db

}
