package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/rayquaza1994/microservices-demo/services/auth/pkg/config"
	"github.com/rayquaza1994/microservices-demo/services/auth/pkg/db"
	authv2connect "github.com/rayquaza1994/microservices-demo/shared/proto/auth/v2/authv2connect"

	"github.com/rayquaza1994/microservices-demo/services/auth/pkg/services"
	"github.com/rayquaza1994/microservices-demo/services/auth/pkg/utils"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "go-grpc-auth-svc",
		ExpirationHours: 24 * 365,
	}

	fmt.Println("Auth Svc on", c.Port)

	s := services.Server{
		H:   h,
		Jwt: jwt,
	}

	mux := http.NewServeMux()
	path, handler := authv2connect.NewAuthServiceHandler(&s)
	mux.Handle(path, handler)
	http.ListenAndServe(
		fmt.Sprintf("0.0.0.0%s", c.Port),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

}
