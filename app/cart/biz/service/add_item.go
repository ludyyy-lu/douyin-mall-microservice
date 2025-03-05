package service

import (
	"context"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/model"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/rpc"
	cart "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"
	product "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product"

	RPCproduct "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/product/productcatalogservice"

	//RPCproduct "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/rpc/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ProductClient RPCproduct.Client
	//CartStore     model.CartStore
	Ctx           context.Context
}

// NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{Ctx: ctx}
}

// Run create note info

func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	productResp, err := rpc.ProductClient.GetProduct(s.Ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}
	if productResp == nil || productResp.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not found")
	}
	cartItem := &model.Cart{
		UserID:    req.UserId,
		ProductID: req.Item.ProductId,
		Qty:       req.Item.Quantity,
	}
	err = model.AddCart(mysql.DB,s.Ctx,cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	return &cart.AddItemResp{}, nil
}
