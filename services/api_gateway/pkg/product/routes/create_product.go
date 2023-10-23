package routes

import (
	"context"
	"net/http"

	connect_go "github.com/bufbuild/connect-go"
	"github.com/gin-gonic/gin"
	v2 "github.com/rayquaza1994/microservices-demo/shared/proto/product/v2"
	pb "github.com/rayquaza1994/microservices-demo/shared/proto/product/v2/productv2connect"
)

type CreateProductRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

func CreateProduct(ctx *gin.Context, c pb.ProductServiceClient) {
	b := CreateProductRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req := v2.CreateProductRequest{
		Name:  b.Name,
		Stock: b.Stock,
		Price: b.Price,
	}

	res, err := c.CreateProduct(context.Background(), connect_go.NewRequest(&req))

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
