package core

import (
	"douyin-mall/rpc/order/internal/global"
	"github.com/Mmx233/EnvConfig"
)

func InitConfig() {
	global.Config.Init()
	EnvConfig.Load("ORDER_SERVICE_", &global.Config.ServiceInfo)
}
