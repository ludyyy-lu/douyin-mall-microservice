package rpc

import (
	"fmt"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/v2/resolver"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	//保证只能初始化一次
	once sync.Once
)

func Init() {
	once.Do(func() {

		iniProductClient()
	})
}
func iniProductClient() {

	r, err := resolver.NewDefaultNacosResolver()
	fmt.Println("product服务发现", r.Name())
	if err != nil {
		hlog.Fatal(err)
	}
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}

}
