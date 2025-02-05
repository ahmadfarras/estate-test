package usecase

import (
	"context"
	"sort"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// CreateEstate implements EstateUsecase.
func (e *EstateUsecaseImpl) CreateEstate(ctx context.Context, req generated.PostEstateJSONRequestBody,
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
func (e *EstateUsecaseImpl) CreateTree(ctx context.Context, id uuid.UUID, req generated.PostEstateIdTreeJSONRequestBody,
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

// GetEstateStats implements EstateUsecase.
func (e *EstateUsecaseImpl) GetEstateStats(ctx context.Context, id uuid.UUID,
) (resp generated.GetEstateStatResponse, err error) {
	estateTrees, err := e.Repository.GetAllTreesByEstateID(ctx, repository.GetAllTreesByEstateIDInput{EstateID: id})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get all trees by estate id")
		return generated.GetEstateStatResponse{}, err
	}

	if estateTrees.Total == 0 {
		return generated.GetEstateStatResponse{}, nil
	}

	var (
		medianHeight float32
	)

	sort.Slice(estateTrees.Trees, func(i, j int) bool {
		return estateTrees.Trees[i].Height < estateTrees.Trees[j].Height
	})

	maxHeight := estateTrees.Trees[len(estateTrees.Trees)-1].Height
	minHeight := estateTrees.Trees[0].Height
	totalCountOfTrees := estateTrees.Total

	if totalCountOfTrees%2 == 0 {
		medianHeight = float32(estateTrees.Trees[totalCountOfTrees/2-1].Height+estateTrees.Trees[totalCountOfTrees/2].Height) / 2
	} else {
		medianHeight = float32(estateTrees.Trees[totalCountOfTrees/2].Height)
	}

	resp = generated.GetEstateStatResponse{
		Count:  totalCountOfTrees,
		Max:    maxHeight,
		Min:    minHeight,
		Median: medianHeight,
	}

	return resp, nil

}
