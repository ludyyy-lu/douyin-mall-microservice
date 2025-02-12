package config

import (
	projectConfig "douyin-mall/config"
)

type Config struct {
	projectConfig.Config
	ServiceInfo ServiceInfo
}
