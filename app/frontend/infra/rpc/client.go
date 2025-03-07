
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


// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"context"
	"sync"

	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/conf"
	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/infra/mtl"
	frontendutils "github.com/All-Done-Right/douyin-mall-microservice/app/frontend/utils"
	"github.com/All-Done-Right/douyin-mall-microservice/common/clientsuite"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/auth/authservice"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/checkout/checkoutservice"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/order/orderservice"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	prometheus "github.com/kitex-contrib/monitor-prometheus"
)

var (
	AuthClient     authservice.Client
	ProductClient  productcatalogservice.Client
	UserClient     userservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
	err            error
	registryAddr   string
	commonSuite    client.Option
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Hertz.RegistryAddr
		commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
			RegistryAddr:       registryAddr,
			CurrentServiceName: frontendutils.ServiceName,
		})
		initAuthClient()
		initProductClient()
		initUserClient()
		initCartClient()
		initCheckoutClient()
		initOrderClient()
    
	})
}

// func initProductClient() {
// 	var opts []client.Option

// 	cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
// 		return circuitbreak.RPCInfo2Key(ri)
// 	})
// 	cbs.UpdateServiceCBConfig("shop-frontend/product/GetProduct", circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2})

// 	opts = append(opts, commonSuite, client.WithCircuitBreaker(cbs), client.WithFallback(fallback.NewFallbackPolicy(fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
// 		methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
// 		if err == nil {
// 			return resp, err
// 		}
// 		if methodName != "ListProducts" {
// 			return resp, err
// 		}
// 		return &product.ListProductsResp{
// 			Products: []*product.Product{
// 				{
// 					Price:       6.6,
// 					Id:          3,
// 					Picture:     "/static/image/t-shirt.jpeg",
// 					Name:        "T-Shirt",
// 					Description: "CloudWeGo T-Shirt",
// 				},
// 			},
// 		}, nil
// 	}))))
// 	opts = append(opts, client.WithTracer(prometheus.NewClientTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))))

// 	ProductClient, err = productcatalogservice.NewClient("product", opts...)
// 	frontendutils.MustHandleError(err)
// }

func initAuthClient() {
	AuthClient, err = authservice.NewClient("auth", commonSuite)
	frontendutils.MustHandleError(err)
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", commonSuite)
	frontendutils.MustHandleError(err)
}

func initCartClient() {
	CartClient, err = cartservice.NewClient("cart", commonSuite)
	frontendutils.MustHandleError(err)
}

func initCheckoutClient() {
	CheckoutClient, err = checkoutservice.NewClient("checkout", commonSuite)
	frontendutils.MustHandleError(err)
}

func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", commonSuite)
	frontendutils.MustHandleError(err)

}


  
