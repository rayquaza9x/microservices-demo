package services

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/rayquaza1994/microservices-demo/services/product/pkg/db"
	"github.com/rayquaza1994/microservices-demo/services/product/pkg/models"
	productv2 "github.com/rayquaza1994/microservices-demo/shared/proto/product/v2"
)

type Server struct {
	H db.Handler
}

func (s *Server) CreateProduct(ctx context.Context, req *connect.Request[productv2.CreateProductRequest]) (*connect.Response[productv2.CreateProductResponse], error) {
	var product models.Product

	product.Name = req.Msg.Name
	product.Stock = req.Msg.Stock
	product.Price = req.Msg.Price

	if result := s.H.DB.Create(&product); result.Error != nil {
		// return &pb.CreateProductResponse{
		// 	Status: http.StatusConflict,
		// 	Error:  result.Error.Error(),
		// }, nil
		res := connect.NewResponse(&productv2.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		})
		return res, nil
	}

	// return &pb.CreateProductResponse{
	// 	Status: http.StatusCreated,
	// 	Id:     product.Id,
	// }, nil
	res := connect.NewResponse(&productv2.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.Id,
	})
	return res, nil
}

func (s *Server) FindOne(ctx context.Context, req *connect.Request[productv2.FindOneRequest]) (*connect.Response[productv2.FindOneResponse], error) {
	var product models.Product

	if result := s.H.DB.First(&product, req.Msg.Id); result.Error != nil {
		// return &pb.FindOneResponse{
		// 	Status: http.StatusNotFound,
		// 	Error:  result.Error.Error(),
		// }, nil
		res := connect.NewResponse(&productv2.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		})
		return res, nil
	}

	data := &productv2.FindOneData{
		Id:    product.Id,
		Name:  product.Name,
		Stock: product.Stock,
		Price: product.Price,
	}

	// return &pb.FindOneResponse{
	// 	Status: http.StatusOK,
	// 	Data:   data,
	// }, nil
	res := connect.NewResponse(&productv2.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	})
	return res, nil
}

func (s *Server) DecreaseStock(ctx context.Context, req *connect.Request[productv2.DecreaseStockRequest]) (*connect.Response[productv2.DecreaseStockResponse], error) {
	var product models.Product

	if result := s.H.DB.First(&product, req.Msg.Id); result.Error != nil {
		// return &pb.DecreaseStockResponse{
		// 	Status: http.StatusNotFound,
		// 	Error:  result.Error.Error(),
		// }, nil
		res := connect.NewResponse(&productv2.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		})
		return res, nil
	}

	if product.Stock <= 0 {
		// return &pb.DecreaseStockResponse{
		// 	Status: http.StatusConflict,
		// 	Error:  "Stock too low",
		// }, nil
		res := connect.NewResponse(&productv2.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock too low",
		})
		return res, nil
	}

	var log models.StockDecreaseLog

	if result := s.H.DB.Where(&models.StockDecreaseLog{OrderId: req.Msg.OrderId}).First(&log); result.Error == nil {
		// return &pb.DecreaseStockResponse{
		// 	Status: http.StatusConflict,
		// 	Error:  "Stock already decreased",
		// }, nil
		res := connect.NewResponse(&productv2.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock already decreased",
		})
		return res, nil
	}

	product.Stock = product.Stock - 1

	s.H.DB.Save(&product)

	log.OrderId = req.Msg.OrderId
	log.ProductRefer = product.Id

	s.H.DB.Create(&log)

	// return &pb.DecreaseStockResponse{
	// 	Status: http.StatusOK,
	// }, nil
	res := connect.NewResponse(&productv2.DecreaseStockResponse{
		Status: http.StatusOK,
	})
	return res, nil
}
