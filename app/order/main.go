package main

import (
	"fmt"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/core"
	"github.com/All-Done-Right/douyin-mall-microservice/app/order/biz/global"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/sirupsen/logrus"
	"log"
	"net"
)

func main() {
	core.InitConfig()
	fmt.Println(global.Config)
	core.InitMysql()
	core.InitRedis()
	r, err := consul.NewConsulRegister(global.Config.Consul.Addr())
	if err != nil {
		klog.Fatal(err)
	}

	logrus.Infoln(r)
	addr, err := net.ResolveTCPAddr("tcp", ":8889")
	if err != nil {
		logrus.Fatal(err)
	}
	svr := orderservice.NewServer(new(OrderServiceImpl),
		server.WithServiceAddr(addr),
		//指定 Registry 与服务基本信息
		server.WithRegistry(r),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: global.Config.ServiceInfo.Name,
			},
		),
	)
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
