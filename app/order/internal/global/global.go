package global

import (
	"douyin-mall/app/order/internal/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config config.Config
	DB     *gorm.DB
	RDB    *redis.Client
)
