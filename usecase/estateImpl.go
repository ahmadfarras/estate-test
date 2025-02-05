package usecase

import (
	"context"
	"math"
	"sort"
	"strconv"

	globalError "github.com/SawitProRecruitment/UserService/error"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/model"
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

	if estate == nil {
		return generated.CreateTreeResponse{}, globalError.ErrEstateNotFound
	}

	if req.X > estate.Length || req.Y > estate.Width {
		return generated.CreateTreeResponse{}, globalError.ErrTreePositionOutOfBoundary
	}

	if req.X < 0 || req.Y < 0 {
		return generated.CreateTreeResponse{}, globalError.ErrTreePositionNegative
	}

	if req.Height < 0 {
		return generated.CreateTreeResponse{}, globalError.ErrTreeHeightNegative
	}

	plantedTree, err := e.Repository.GetTreeByEstateIDAndCoordinate(ctx, repository.GetTreeByEstateIDAndCoordinateInput{
		EstateID: id, X: req.X, Y: req.Y})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get tree by estate id and coordinate")
		return generated.CreateTreeResponse{}, err
	}
	if plantedTree != nil {
		return generated.CreateTreeResponse{}, globalError.ErrTreeAlreadyPlanted
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
	estate, err := e.Repository.GetEstateById(ctx, repository.GetEstateByIdInput{ID: id})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get estate by id")
		return generated.GetEstateStatResponse{}, err
	}

	if estate == nil {
		return generated.GetEstateStatResponse{}, globalError.ErrEstateNotFound
	}

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

// CalculateTravelDistance implements EstateUsecase.
func (e *EstateUsecaseImpl) CalculateTravelDistance(ctx context.Context, id uuid.UUID,
) (resp generated.GetDronePlaneResponse, err error) {
	estate, err := e.Repository.GetEstateById(ctx, repository.GetEstateByIdInput{ID: id})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get estate by id")
		return generated.GetDronePlaneResponse{}, err
	}

	if estate == nil {
		return generated.GetDronePlaneResponse{}, globalError.ErrEstateNotFound
	}

	estateTrees, err := e.Repository.GetAllTreesByEstateID(ctx, repository.GetAllTreesByEstateIDInput{EstateID: id})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to get all trees by estate id")
		return generated.GetDronePlaneResponse{}, err
	}

	treeMap := mappingTreeByItsCoordinate(estateTrees.Trees)

	totalDistance := processCalculateTotalDistance(estate.Estate, treeMap)

	resp = generated.GetDronePlaneResponse{
		Distance: float32(totalDistance),
	}

	return resp, nil
}

func mappingTreeByItsCoordinate(trees []model.Tree) map[string]model.Tree {
	treeMap := make(map[string]model.Tree)
	for _, tree := range trees {
		key := strconv.Itoa(tree.X) + "," + strconv.Itoa(tree.Y)
		treeMap[key] = tree
	}
	return treeMap
}

func getTreeByCoordinates(x, y int, treeMap map[string]model.Tree) (*model.Tree, bool) {
	key := strconv.Itoa(x) + "," + strconv.Itoa(y)
	tree, exists := treeMap[key]
	if !exists {
		return nil, false
	}
	return &tree, true
}

func processCalculateTotalDistance(estate model.Estate, treeMap map[string]model.Tree) int {
	totalDistance := 2
	previousTreeHeight := 0
	estateArea := estate.Length * estate.Width
	for i := 1; i <= estateArea; i++ {
		if i != estateArea {
			totalDistance += 10
		}

		var c int
		x := i / estate.Width
		if x%estate.Width == 0 {
			c = i % estate.Width
		} else {
			c = estate.Width - 1 - (i % estate.Width)
		}
		y := c + 1
		tree, exists := getTreeByCoordinates(x, y, treeMap)
		if exists {
			totalDistance += int(math.Abs(float64(tree.Height) - float64(previousTreeHeight)))
			previousTreeHeight = tree.Height
		} else {
			totalDistance += previousTreeHeight
			previousTreeHeight = 0
		}
	}
	return totalDistance
}
