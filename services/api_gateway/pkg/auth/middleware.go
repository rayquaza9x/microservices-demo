package auth

import (
	"context"
	"log"
	"net/http"
	"strings"

	connect_go "github.com/bufbuild/connect-go"
	"github.com/gin-gonic/gin"
	v2 "github.com/rayquaza1994/microservices-demo/shared/proto/auth/v2"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		log.Println("no token")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	req := v2.ValidateRequest{
		Token: token[1],
	}

	res, err := c.svc.Client.Validate(context.Background(), connect_go.NewRequest(&req))

	if err != nil || res.Msg.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.Msg.UserId)

	ctx.Next()
}
