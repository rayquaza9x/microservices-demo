package routes

import (
	"context"
	"net/http"
	"strconv"

	connect_go "github.com/bufbuild/connect-go"
	"github.com/gin-gonic/gin"
	v2 "github.com/rayquaza1994/microservices-demo/shared/proto/product/v2"
	pb "github.com/rayquaza1994/microservices-demo/shared/proto/product/v2/productv2connect"
)

func FineOne(ctx *gin.Context, c pb.ProductServiceClient) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	req := v2.FindOneRequest{
		Id: int64(id),
	}

	res, err := c.FindOne(context.Background(), connect_go.NewRequest(&req))

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res.Msg)
}
