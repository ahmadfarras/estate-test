package usecase

import (
	"context"
	"reflect"
	"testing"

	globalError "github.com/SawitProRecruitment/UserService/error"
	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestEstateUsecaseImpl_CreateEstate(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		req generated.PostEstateJSONRequestBody
	}
	tests := []struct {
		name     string
		e        *EstateUsecaseImpl
		args     args
		wantResp generated.CreateEstateResponse
		wantErr  bool
	}{
		{
			name: "successful creation",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().CreateEstate(gomock.Any(), gomock.Any()).Return(nil).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				req: generated.PostEstateJSONRequestBody{
					Length: 10,
					Width:  20,
				},
			},
			wantErr: false,
		},
		{
			name: "Error creation",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().CreateEstate(gomock.Any(), gomock.Any()).Return(assert.AnError).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				req: generated.PostEstateJSONRequestBody{
					Length: 10,
					Width:  20,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.e.CreateEstate(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("EstateUsecaseImpl.CreateEstate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEstateUsecaseImpl_CreateTree(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		id  uuid.UUID
		req generated.PostEstateIdTreeJSONRequestBody
	}
	tests := []struct {
		name     string
		e        *EstateUsecaseImpl
		args     args
		wantResp generated.CreateTreeResponse
		wantErr  bool
	}{
		{
			name: "successful creation",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().GetEstateById(gomock.Any(), gomock.Any()).Return(&repository.GetEstateByIdOutput{
						Estate: model.Estate{ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), Length: 10, Width: 20},
					}, nil).Times(1)
					r.EXPECT().GetTreeByEstateIDAndCoordinate(gomock.Any(), gomock.Any()).Return(nil, nil).Times(1)
					r.EXPECT().CreateTree(gomock.Any(), gomock.Any()).Return(nil).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				id:  uuid.New(),
				req: generated.PostEstateIdTreeJSONRequestBody{
					X:      5,
					Y:      5,
					Height: 10,
				},
			},
			wantErr: false,
		},
		{
			name: "request x coordinate out of boundary",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().GetEstateById(gomock.Any(), gomock.Any()).Return(&repository.GetEstateByIdOutput{
						Estate: model.Estate{ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), Length: 10, Width: 20},
					}, nil).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				id:  uuid.New(),
				req: generated.PostEstateIdTreeJSONRequestBody{
					X:      100,
					Y:      100,
					Height: 10,
				},
			},
			wantErr: true,
		},
		{
			name: "request x coordinate negative",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().GetEstateById(gomock.Any(), gomock.Any()).Return(&repository.GetEstateByIdOutput{
						Estate: model.Estate{ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), Length: 10, Width: 20},
					}, nil).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				id:  uuid.New(),
				req: generated.PostEstateIdTreeJSONRequestBody{
					X:      -100,
					Y:      -100,
					Height: 10,
				},
			},
			wantErr: true,
		},
		{
			name: "request x heigh negative",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().GetEstateById(gomock.Any(), gomock.Any()).Return(&repository.GetEstateByIdOutput{
						Estate: model.Estate{ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), Length: 10, Width: 20},
					}, nil).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				id:  uuid.New(),
				req: generated.PostEstateIdTreeJSONRequestBody{
					X:      10,
					Y:      10,
					Height: -10,
				},
			},
			wantErr: true,
		},
		{
			name: "estate not found",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().GetEstateById(gomock.Any(), gomock.Any()).Return(nil, globalError.ErrEstateNotFound).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				id:  uuid.New(),
				req: generated.PostEstateIdTreeJSONRequestBody{
					X:      5,
					Y:      5,
					Height: 10,
				},
			},
			wantResp: generated.CreateTreeResponse{},
			wantErr:  true,
		},
		{
			name: "tree already planted",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().GetEstateById(gomock.Any(), gomock.Any()).Return(&repository.GetEstateByIdOutput{
						Estate: model.Estate{ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), Length: 10, Width: 20},
					}, nil).Times(1)
					r.EXPECT().GetTreeByEstateIDAndCoordinate(gomock.Any(), gomock.Any()).Return(
						&repository.GetTreeByEstateIDAndCoordinateOutput{
							Tree: model.Tree{
								ID:       "123e4567-e89b-12d3-a456-426614174000",
								EstateID: "123e4567-e89b-12d3-a456-426614174000",
								Height:   10,
								X:        5,
								Y:        5,
							},
						}, nil).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				id:  uuid.New(),
				req: generated.PostEstateIdTreeJSONRequestBody{
					X:      5,
					Y:      5,
					Height: 10,
				},
			},
			wantResp: generated.CreateTreeResponse{},
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.e.CreateTree(tt.args.ctx, tt.args.id, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("EstateUsecaseImpl.CreateTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestEstateUsecaseImpl_GetEstateStats(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name     string
		e        *EstateUsecaseImpl
		args     args
		wantResp generated.GetEstateStatResponse
		wantErr  bool
	}{
		{
			name: "successful stats retrieval",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().GetEstateById(gomock.Any(), gomock.Any()).Return(&repository.GetEstateByIdOutput{
						Estate: model.Estate{ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), Length: 10, Width: 20},
					}, nil).Times(1)
					r.EXPECT().GetAllTreesByEstateID(gomock.Any(), gomock.Any()).Return(repository.GetAllTreesByEstateIDOutput{
						Trees: []model.Tree{
							{Height: 5},
							{Height: 10},
							{Height: 15},
						},
						Total: 3,
					}, nil).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				id:  uuid.New(),
			},
			wantResp: generated.GetEstateStatResponse{
				Count:  3,
				Max:    15,
				Min:    5,
				Median: 10,
			},
			wantErr: false,
		},
		{
			name: "estate not found",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().GetEstateById(gomock.Any(), gomock.Any()).Return(nil, globalError.ErrEstateNotFound).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				id:  uuid.New(),
			},
			wantResp: generated.GetEstateStatResponse{},
			wantErr:  true,
		},
		{
			name: "no trees found",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().GetEstateById(gomock.Any(), gomock.Any()).Return(&repository.GetEstateByIdOutput{
						Estate: model.Estate{ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), Length: 10, Width: 20},
					}, nil).Times(1)
					r.EXPECT().GetAllTreesByEstateID(gomock.Any(), gomock.Any()).Return(repository.GetAllTreesByEstateIDOutput{
						Trees: []model.Tree{},
						Total: 0,
					}, nil).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				id:  uuid.New(),
			},
			wantResp: generated.GetEstateStatResponse{},
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.e.GetEstateStats(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("EstateUsecaseImpl.GetEstateStats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("EstateUsecaseImpl.GetEstateStats() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestEstateUsecaseImpl_CalculateTravelDistance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx context.Context
		id  uuid.UUID
	}
	tests := []struct {
		name     string
		e        *EstateUsecaseImpl
		args     args
		wantResp generated.GetDronePlaneResponse
		wantErr  bool
	}{
		{
			name: "successful distance calculation",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().GetEstateById(gomock.Any(), gomock.Any()).Return(&repository.GetEstateByIdOutput{
						Estate: model.Estate{ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), Length: 10, Width: 20},
					}, nil).Times(1)
					r.EXPECT().GetAllTreesByEstateID(gomock.Any(), gomock.Any()).Return(repository.GetAllTreesByEstateIDOutput{
						Trees: []model.Tree{
							{Height: 5},
							{Height: 10},
							{Height: 15},
						},
						Total: 3,
					}, nil).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				id:  uuid.New(),
			},
			wantResp: generated.GetDronePlaneResponse{
				Distance: 2022,
			},
			wantErr: false,
		},
		{
			name: "estate not found",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().GetEstateById(gomock.Any(), gomock.Any()).Return(nil, globalError.ErrEstateNotFound).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				id:  uuid.New(),
			},
			wantResp: generated.GetDronePlaneResponse{},
			wantErr:  true,
		},
		{
			name: "no trees found",
			e: &EstateUsecaseImpl{
				Repository: func() repository.RepositoryInterface {
					r := repository.NewMockRepositoryInterface(ctrl)
					r.EXPECT().GetEstateById(gomock.Any(), gomock.Any()).Return(&repository.GetEstateByIdOutput{
						Estate: model.Estate{ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), Length: 10, Width: 20},
					}, nil).Times(1)
					r.EXPECT().GetAllTreesByEstateID(gomock.Any(), gomock.Any()).Return(repository.GetAllTreesByEstateIDOutput{
						Trees: []model.Tree{},
						Total: 0,
					}, nil).Times(1)
					return r
				}(),
			},
			args: args{
				ctx: context.TODO(),
				id:  uuid.New(),
			},
			wantResp: generated.GetDronePlaneResponse{
				Distance: 1992,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.e.CalculateTravelDistance(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("EstateUsecaseImpl.CalculateTravelDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("EstateUsecaseImpl.CalculateTravelDistance() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
