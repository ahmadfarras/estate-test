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
						Estate: model.Estate{ID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"), Length: 5, Width: 1},
					}, nil).Times(1)
					r.EXPECT().GetAllTreesByEstateID(gomock.Any(), gomock.Any()).Return(repository.GetAllTreesByEstateIDOutput{
						Trees: []model.Tree{
							{Height: 5, X: 2, Y: 1},
							{Height: 3, X: 3, Y: 1},
							{Height: 4, X: 4, Y: 1},
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
				Distance: 54,
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

func Test_mappingTreeByItsCoordinate(t *testing.T) {
	type args struct {
		trees []model.Tree
	}
	tests := []struct {
		name string
		args args
		want map[string]model.Tree
	}{
		{
			name: "basic test",
			args: args{
				trees: []model.Tree{
					{Height: 10, X: 1, Y: 1},
					{Height: 15, X: 2, Y: 3},
					{Height: 20, X: 3, Y: 2},
				},
			},
			want: map[string]model.Tree{
				"1,1": {Height: 10, X: 1, Y: 1},
				"2,3": {Height: 15, X: 2, Y: 3},
				"3,2": {Height: 20, X: 3, Y: 2},
			},
		},
		{
			name: "empty list",
			args: args{
				trees: []model.Tree{},
			},
			want: map[string]model.Tree{},
		},
		{
			name: "duplicate coordinates",
			args: args{
				trees: []model.Tree{
					{Height: 10, X: 1, Y: 1},
					{Height: 15, X: 1, Y: 1},
				},
			},
			want: map[string]model.Tree{
				"1,1": {Height: 15, X: 1, Y: 1}, // The last tree with the same coordinates should overwrite the previous one
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mappingTreeByItsCoordinate(tt.args.trees); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mappingTreeByItsCoordinate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getTreeByCoordinates(t *testing.T) {
	type args struct {
		x       int
		y       int
		treeMap map[string]model.Tree
	}
	tests := []struct {
		name  string
		args  args
		want  *model.Tree
		want1 bool
	}{
		{
			name: "tree exists",
			args: args{
				x: 2,
				y: 3,
				treeMap: map[string]model.Tree{
					"2,3": {Height: 15, X: 2, Y: 3},
				},
			},
			want:  &model.Tree{Height: 15, X: 2, Y: 3},
			want1: true,
		},
		{
			name: "tree does not exist",
			args: args{
				x: 1,
				y: 1,
				treeMap: map[string]model.Tree{
					"2,3": {Height: 15, X: 2, Y: 3},
				},
			},
			want:  nil,
			want1: false,
		},
		{
			name: "empty map",
			args: args{
				x:       1,
				y:       1,
				treeMap: map[string]model.Tree{},
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getTreeByCoordinates(tt.args.x, tt.args.y, tt.args.treeMap)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTreeByCoordinates() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getTreeByCoordinates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_processCalculateTotalDistance(t *testing.T) {
	type args struct {
		estate  model.Estate
		treeMap map[string]model.Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic test",
			args: args{
				estate: model.Estate{Length: 5, Width: 1},
				treeMap: map[string]model.Tree{
					"2,1": {Height: 5, X: 2, Y: 1},
					"3,1": {Height: 3, X: 3, Y: 1},
					"4,1": {Height: 4, X: 4, Y: 1},
				},
			},
			want: 54,
		},
		{
			name: "estate not found",
			args: args{
				estate: model.Estate{Length: 5, Width: 1},
				treeMap: map[string]model.Tree{
					"2,1": {Height: 5, X: 2, Y: 1},
					"3,1": {Height: 3, X: 3, Y: 1},
					"4,1": {Height: 4, X: 4, Y: 1},
				},
			},
			want: 54,
		},
		{
			name: "no trees found",
			args: args{
				estate:  model.Estate{Length: 10, Width: 20},
				treeMap: map[string]model.Tree{},
			},
			want: 1992,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processCalculateTotalDistance(tt.args.estate, tt.args.treeMap); got != tt.want {
				t.Errorf("processCalculateTotalDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
