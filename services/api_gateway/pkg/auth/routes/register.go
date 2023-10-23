package routes

import (
	"context"
	"net/http"

	connect_go "github.com/bufbuild/connect-go"
	"github.com/gin-gonic/gin"
	v2 "github.com/rayquaza1994/microservices-demo/shared/proto/auth/v2"
	pb "github.com/rayquaza1994/microservices-demo/shared/proto/auth/v2/authv2connect"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	b := RegisterRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req := v2.RegisterRequest{
		Email:    b.Email,
		Password: b.Password,
	}

	res, err := c.Register(context.Background(), connect_go.NewRequest(&req))

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
