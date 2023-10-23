package services

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"github.com/rayquaza1994/microservices-demo/services/auth/pkg/db"
	"github.com/rayquaza1994/microservices-demo/services/auth/pkg/models"
	"github.com/rayquaza1994/microservices-demo/services/auth/pkg/utils"
	authv2 "github.com/rayquaza1994/microservices-demo/shared/proto/auth/v2"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
}

// func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
func (s *Server) Register(ctx context.Context, req *connect.Request[authv2.RegisterRequest]) (*connect.Response[authv2.RegisterResponse], error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Msg.Email}).First(&user); result.Error == nil {
		res := connect.NewResponse(&authv2.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		})
		return res, nil
		// return &pb.RegisterResponse{
		// 	Status: http.StatusConflict,
		// 	Error:  "E-Mail already exists",
		// }, nil
	}

	user.Email = req.Msg.Email
	user.Password = utils.HashPassword(req.Msg.Password)

	s.H.DB.Create(&user)

	// return &pb.RegisterResponse{
	// 	Status: http.StatusCreated,
	// }, nil

	res := connect.NewResponse(&authv2.RegisterResponse{
		Status: http.StatusCreated,
	})
	return res, nil
}

// func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
func (s *Server) Login(ctx context.Context, req *connect.Request[authv2.LoginRequest]) (*connect.Response[authv2.LoginResponse], error) {
	var user models.User

	if result := s.H.DB.Where(&models.User{Email: req.Msg.Email}).First(&user); result.Error != nil {
		res := connect.NewResponse(&authv2.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		})
		return res, nil
		// return &pb.LoginResponse{
		// 	Status: http.StatusNotFound,
		// 	Error:  "User not found",
		// }, nil
	}

	match := utils.CheckPasswordHash(req.Msg.Password, user.Password)

	if !match {
		res := connect.NewResponse(&authv2.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		})
		return res, nil
		// return &pb.LoginResponse{
		// 	Status: http.StatusNotFound,
		// 	Error:  "User not found",
		// }, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	res := connect.NewResponse(&authv2.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	})
	return res, nil

	// return &pb.LoginResponse{
	// 	Status: http.StatusOK,
	// 	Token:  token,
	// }, nil
}

// func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
func (s *Server) Validate(ctx context.Context, req *connect.Request[authv2.ValidateRequest]) (*connect.Response[authv2.ValidateResponse], error) {
	claims, err := s.Jwt.ValidateToken(req.Msg.Token)

	if err != nil {
		res := connect.NewResponse(&authv2.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		})
		return res, nil
		// return &pb.ValidateResponse{
		// 	Status: http.StatusBadRequest,
		// 	Error:  err.Error(),
		// }, nil
	}

	var user models.User

	if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
		res := connect.NewResponse(&authv2.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		})
		return res, nil
		// return &pb.ValidateResponse{
		// 	Status: http.StatusNotFound,
		// 	Error:  "User not found",
		// }, nil
	}

	res := connect.NewResponse(&authv2.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	})
	return res, nil

	// return &pb.ValidateResponse{
	// 	Status: http.StatusOK,
	// 	UserId: user.Id,
	// }, nil
}
