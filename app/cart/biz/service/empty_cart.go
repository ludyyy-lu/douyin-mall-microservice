package service

import (
	"context"

	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/dal/mysql"
	"github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/model"

	cart "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/cart"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

/*
	type EmptyCartService struct {
		ctx context.Context
	}
*/
type EmptyCartService struct {
	//CartStore model.CartStore
	Ctx       context.Context
}

// NewEmptyCartService new EmptyCartService
func NewEmptyCartService(Ctx context.Context) *EmptyCartService {
	return &EmptyCartService{Ctx: Ctx}
}

// Run create note info

func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	err = model.EmptyCart(s.Ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50001, err.Error())
	}
	return &cart.EmptyCartResp{}, nil
}
