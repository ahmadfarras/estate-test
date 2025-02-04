package usecase

import (
	"github.com/SawitProRecruitment/UserService/repository"
)

type EstateUsecaseImpl struct {
	Repository repository.RepositoryInterface
}

func NewEstateUsecase(repo repository.RepositoryInterface) EstateUsecase {
	return &EstateUsecaseImpl{
		repo,
	}
}
