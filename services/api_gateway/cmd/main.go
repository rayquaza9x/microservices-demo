package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rayquaza1994/microservices-demo/services/api_gateway/pkg/auth"
	"github.com/rayquaza1994/microservices-demo/services/api_gateway/pkg/config"
	// "github.com/rayquaza1994/microservices-demo/pkg/order"
	"github.com/rayquaza1994/microservices-demo/services/api_gateway/pkg/product"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	log.Println(c)

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	// order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
