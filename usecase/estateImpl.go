package usecase

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// CreateEstate implements EstateUsecase.
func (e *EstateUsecaseImpl) CreateEstate(
	ctx context.Context, req generated.PostEstateJSONRequestBody,
) (resp generated.CreateEstateResponse, err error) {
	estateInput := repository.CreateEstateInput{
		ID:     uuid.New(),
		Length: req.Length,
		Width:  req.Width,
	}

	err = e.Repository.CreateEstate(ctx, estateInput)
	if err != nil {
		logrus.WithError(err).Error("failed to create estate")
		return generated.CreateEstateResponse{}, err
	}

	resp = generated.CreateEstateResponse{
		Id: estateInput.ID,
	}

	return resp, nil
}

// CreateTree implements EstateUsecase.
func (e *EstateUsecaseImpl) CreateTree(
	ctx context.Context, id uuid.UUID, req generated.PostEstateIdTreeJSONRequestBody,
) (resp generated.CreateTreeResponse, err error) {
	estate, err := e.Repository.GetEstateById(ctx, repository.GetEstateByIdInput{ID: id})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get estate by id")
		return generated.CreateTreeResponse{}, err
	}

	treeInput := repository.CreateTreeInput{
		ID:       uuid.New(),
		EstateID: estate.ID,
		Height:   req.Height,
		X:        req.X,
		Y:        req.Y,
	}

	err = e.Repository.CreateTree(ctx, treeInput)
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to create tree")
		return generated.CreateTreeResponse{}, err
	}

	resp = generated.CreateTreeResponse{
		Id: treeInput.ID,
	}

	return resp, nil
}
