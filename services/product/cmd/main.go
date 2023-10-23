package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/rayquaza1994/microservices-demo/services/product/pkg/config"
	"github.com/rayquaza1994/microservices-demo/services/product/pkg/db"
	services "github.com/rayquaza1994/microservices-demo/services/product/pkg/services"
	productv2connect "github.com/rayquaza1994/microservices-demo/shared/proto/product/v2/productv2connect"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	fmt.Println("Product Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	mux := http.NewServeMux()
	path, handler := productv2connect.NewProductServiceHandler(&s)
	mux.Handle(path, handler)
	http.ListenAndServe(
		fmt.Sprintf("0.0.0.0%s", c.Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

	// grpcServer := grpc.NewServer()

	// pb.RegisterProductServiceServer(grpcServer, &s)

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalln("Failed to serve:", err)
	// }
}
