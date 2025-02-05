package usecase

import (
	"context"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/google/uuid"
)

type EstateUsecase interface {
	CreateEstate(ctx context.Context, req generated.PostEstateJSONRequestBody) (resp generated.CreateEstateResponse, err error)
	CreateTree(ctx context.Context, id uuid.UUID, req generated.PostEstateIdTreeJSONRequestBody) (resp generated.CreateTreeResponse, err error)
	GetEstateStats(ctx context.Context, id uuid.UUID) (resp generated.GetEstateStatResponse, err error)
	CalculateTravelDistance(ctx context.Context, id uuid.UUID) (resp generated.GetDronePlaneResponse, err error)
}
