package main

import (
	"context"
	"douyin-mall/rpc/order/kitex_gen/order"
	"douyin-mall/rpc/order/kitex_gen/order/orderservice"
	"fmt"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func NewOrderClient() (orderservice.Client, error) {
	// 使用时请传入真实 etcd 的服务地址，本例中为 127.0.0.1:2379
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	return orderservice.NewClient("douyin.mall.order", client.WithResolver(r)) // 指定 Resolver
}
func main() {
	c, err := NewOrderClient()
	if err != nil {
		log.Fatal(err)
	}
	req := new(order.PlaceOrderReq)

	resp, err := c.PlaceOrder(context.Background(), req)
	fmt.Println(111)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(111)
	fmt.Println(resp)
}
