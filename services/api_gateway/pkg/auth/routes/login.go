package routes

import (
	"context"
	"net/http"

	connect_go "github.com/bufbuild/connect-go"
	"github.com/gin-gonic/gin"
	v2 "github.com/rayquaza1994/microservices-demo/shared/proto/auth/v2"
	pb "github.com/rayquaza1994/microservices-demo/shared/proto/auth/v2/authv2connect"

	log "github.com/sirupsen/logrus"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req := v2.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	}

	log.Println("Login request", req)

	res, err := c.Login(context.Background(), connect_go.NewRequest(&req))

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res.Msg)
}
