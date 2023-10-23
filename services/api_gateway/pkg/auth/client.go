package auth

import (
	"fmt"
	"net/http"

	"github.com/rayquaza1994/microservices-demo/services/api_gateway/pkg/config"
	pb "github.com/rayquaza1994/microservices-demo/shared/proto/auth/v2/authv2connect"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	// using WithInsecure() because no SSL running
	// cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	// if err != nil {
	// 	fmt.Println("Could not connect:", err)
	// }

	// return pb.NewAuthServiceClient(cc)
	return pb.NewAuthServiceClient(
		http.DefaultClient,
		fmt.Sprintf("http://%s", c.AuthSvcUrl),
	)
}
