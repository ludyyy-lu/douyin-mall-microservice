package main

import (
	core2 "douyin-mall/app/order/internal/core"
	"douyin-mall/app/order/internal/global"
	order "douyin-mall/app/order/kitex_gen/order/orderservice"
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/sirupsen/logrus"
	"log"
	"net"
)

func main() {
	core2.InitConfig()
	fmt.Println(global.Config)
	core2.InitMysql()
	core2.InitRedis()
	r, err := etcd.NewEtcdRegistry([]string{global.Config.Etcd.Addr()})
	if err != nil {
		log.Fatal(err)
	}
	logrus.Infoln(r)
	addr, err := net.ResolveTCPAddr("tcp", global.Config.ServiceInfo.Addr())
	if err != nil {
		logrus.Fatal(err)
	}
	svr := order.NewServer(new(OrderServiceImpl),
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
