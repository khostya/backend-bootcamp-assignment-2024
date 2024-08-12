// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/usecase/mocks/flat.go
//
// Generated by this command:
//
//	mockgen -source ./internal/usecase/mocks/flat.go -destination=./internal/usecase/mocks/flat_mock.go -package=mock_usecase
//

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	domain "backend-bootcamp-assignment-2024/internal/domain"
	dto "backend-bootcamp-assignment-2024/internal/dto"
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockflatService is a mock of flatService interface.
type MockflatService struct {
	ctrl     *gomock.Controller
	recorder *MockflatServiceMockRecorder
}

// MockflatServiceMockRecorder is the mock recorder for MockflatService.
type MockflatServiceMockRecorder struct {
	mock *MockflatService
}

// NewMockflatService creates a new mock instance.
func NewMockflatService(ctrl *gomock.Controller) *MockflatService {
	mock := &MockflatService{ctrl: ctrl}
	mock.recorder = &MockflatServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockflatService) EXPECT() *MockflatServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockflatService) Create(ctx context.Context, param dto.CreateFlatParam) (domain.Flat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, param)
	ret0, _ := ret[0].(domain.Flat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockflatServiceMockRecorder) Create(ctx, param any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockflatService)(nil).Create), ctx, param)
}

// Update mocks base method.
func (m *MockflatService) Update(ctx context.Context, param dto.UpdateFlatParam) (domain.Flat, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, param)
	ret0, _ := ret[0].(domain.Flat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockflatServiceMockRecorder) Update(ctx, param any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockflatService)(nil).Update), ctx, param)
}
