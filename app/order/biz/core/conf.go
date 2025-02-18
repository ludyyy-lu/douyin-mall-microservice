package core

import (
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/global"
	"github.com/Mmx233/EnvConfig"
)

func InitConfig() {
	global.Config.Init()
	EnvConfig.Load("ORDER_SERVICE_", &global.Config.ServiceInfo)
}
