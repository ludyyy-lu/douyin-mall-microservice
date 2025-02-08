package middleware

import (
	"context"
	"net/http"

	"github.com/All-Done-Right/douyin-mall-microservice/app/frontend/biz/utils"
	authcenter "github.com/All-Done-Right/douyin-mall-microservice/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

// JWTMiddleware 是用于校验 JWT 令牌的中间件
func JWTMiddleware(h *LoginService) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.GetHeader("Authorization")
		if token == "" {
			utils.SendErrResponse(ctx, c, http.StatusUnauthorized, "Missing Authorization header")
			c.Abort()
			return
		}

		// 去掉 "Bearer " 前缀
		// 假设前端传递的令牌格式为 "Bearer <token>"
		// 如果不是这种格式，需要根据实际情况调整
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		verifyReq := &authcenter.VerifyTokenReq{
			Token: token,
		}
		// 通过此客户端 访问 认证服务 的服务端
		tokenResp, err := rpc.AuthClient.DeliverTokenByRPC(h.Context, verifyReq)
		if err != nil {
			return "", "", err
		}
		if err != nil || !verifyResp.Res {
			utils.SendErrResponse(ctx, c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		c.Next(ctx)
	}
}
