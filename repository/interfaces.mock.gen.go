// Code generated by MockGen. DO NOT EDIT.
// Source: repository/interfaces.go
//
// Generated by this command:
//
//	mockgen -source=repository/interfaces.go -destination=repository/interfaces.mock.gen.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepositoryInterface is a mock of RepositoryInterface interface.
type MockRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryInterfaceMockRecorder
	isgomock struct{}
}

// MockRepositoryInterfaceMockRecorder is the mock recorder for MockRepositoryInterface.
type MockRepositoryInterfaceMockRecorder struct {
	mock *MockRepositoryInterface
}

// NewMockRepositoryInterface creates a new mock instance.
func NewMockRepositoryInterface(ctrl *gomock.Controller) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryInterface) EXPECT() *MockRepositoryInterfaceMockRecorder {
	return m.recorder
}

// CreateEstate mocks base method.
func (m *MockRepositoryInterface) CreateEstate(ctx context.Context, input CreateEstateInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEstate", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEstate indicates an expected call of CreateEstate.
func (mr *MockRepositoryInterfaceMockRecorder) CreateEstate(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEstate", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateEstate), ctx, input)
}

// CreateTree mocks base method.
func (m *MockRepositoryInterface) CreateTree(ctx context.Context, input CreateTreeInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTree", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTree indicates an expected call of CreateTree.
func (mr *MockRepositoryInterfaceMockRecorder) CreateTree(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTree", reflect.TypeOf((*MockRepositoryInterface)(nil).CreateTree), ctx, input)
}

// GetAllTreesByEstateID mocks base method.
func (m *MockRepositoryInterface) GetAllTreesByEstateID(ctx context.Context, input GetAllTreesByEstateIDInput) (GetAllTreesByEstateIDOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTreesByEstateID", ctx, input)
	ret0, _ := ret[0].(GetAllTreesByEstateIDOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTreesByEstateID indicates an expected call of GetAllTreesByEstateID.
func (mr *MockRepositoryInterfaceMockRecorder) GetAllTreesByEstateID(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTreesByEstateID", reflect.TypeOf((*MockRepositoryInterface)(nil).GetAllTreesByEstateID), ctx, input)
}

// GetEstateById mocks base method.
func (m *MockRepositoryInterface) GetEstateById(ctx context.Context, input GetEstateByIdInput) (GetEstateByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEstateById", ctx, input)
	ret0, _ := ret[0].(GetEstateByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEstateById indicates an expected call of GetEstateById.
func (mr *MockRepositoryInterfaceMockRecorder) GetEstateById(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEstateById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetEstateById), ctx, input)
}

// GetTestById mocks base method.
func (m *MockRepositoryInterface) GetTestById(ctx context.Context, input GetTestByIdInput) (GetTestByIdOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTestById", ctx, input)
	ret0, _ := ret[0].(GetTestByIdOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTestById indicates an expected call of GetTestById.
func (mr *MockRepositoryInterfaceMockRecorder) GetTestById(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTestById", reflect.TypeOf((*MockRepositoryInterface)(nil).GetTestById), ctx, input)
}
