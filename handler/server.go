package handler

import (
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/usecase"
)

type Server struct {
	EstateUsecase usecase.EstateUsecase
	Repository    repository.RepositoryInterface
}

type NewServerOptions struct {
	EstateUsecase usecase.EstateUsecase
	Repository    repository.RepositoryInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		opts.EstateUsecase,
		opts.Repository,
	}
}
