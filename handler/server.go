package handler

import (
	"github.com/SawitProRecruitment/UserService/usecase"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	EstateUsecase usecase.EstateUsecase
	Validator     *validator.Validate
}

type NewServerOptions struct {
	EstateUsecase usecase.EstateUsecase
	Validator     *validator.Validate
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		opts.EstateUsecase,
		opts.Validator,
	}
}
