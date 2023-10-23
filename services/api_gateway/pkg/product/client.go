package product

import (
	"fmt"
	"net/http"

	"github.com/rayquaza1994/microservices-demo/services/api_gateway/pkg/config"
	pb "github.com/rayquaza1994/microservices-demo/shared/proto/product/v2/productv2connect"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	// using WithInsecure() because no SSL running
	// cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())

	// if err != nil {
	// 	fmt.Println("Could not connect:", err)
	// }

	return pb.NewProductServiceClient(
		http.DefaultClient,
		fmt.Sprintf("http://%s", c.ProductSvcUrl),
	)
}
