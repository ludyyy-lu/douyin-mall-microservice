package global

import (
	"douyin-mall/rpc/order/internal/config"
	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
)
