package main

import (
	"context"
	auth "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/auth"
	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/biz/service"
)

// AuthServiceImpl implements the last service interface defined in the IDL.
type AuthServiceImpl struct{}

// DeliverTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq) (resp *auth.DeliveryResp, err error) {
	resp, err = service.NewDeliverTokenByRPCService(ctx).Run(req)

	return resp, err
}

// VerifyTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq) (resp *auth.VerifyResp, err error) {
	resp, err = service.NewVerifyTokenByRPCService(ctx).Run(req)

	return resp, err
}

// RenewTokenByRPC implements the AuthServiceImpl interface.
func (s *AuthServiceImpl) RenewTokenByRPC(ctx context.Context, req *auth.RenewTokenReq) (resp *auth.DeliveryResp, err error) {
	resp, err = service.NewRenewTokenByRPCService(ctx).Run(req)

	return resp, err
}
