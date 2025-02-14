package middleware

import (
	"context"
	"github.com/All-Done-Right/douyin-mall-microservice/app/auth/biz/service"
	"github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 从请求头中获取 token
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			hlog.CtxErrorf(c, "Authorization header is missing")
			ctx.JSON(401, map[string]string{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		// 调用 VerifyTokenByRPC 进行 token 验证
		verifyReq := &auth.VerifyTokenReq{
			Token: token,
		}
		verifyService := service.NewVerifyTokenByRPCService(c)
		verifyResp, err := verifyService.Run(verifyReq)
		if err != nil || !verifyResp.Res {
			hlog.CtxErrorf(c, "Token verification failed: %v", err)
			ctx.JSON(401, map[string]string{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		// 验证通过，继续处理请求
		ctx.Next(c)
	}
}
