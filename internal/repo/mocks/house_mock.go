// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repo/mocks/house.go
//
// Generated by this command:
//
//	mockgen -source ./internal/repo/mocks/house.go -destination=./internal/repo/mocks/house_mock.go -package=mock_repo
//

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	context "context"
	reflect "reflect"
	time "time"

	domain "github.com/khostya/backend-bootcamp-assignment-2024/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockhouseRepo is a mock of houseRepo interface.
type MockhouseRepo struct {
	ctrl     *gomock.Controller
	recorder *MockhouseRepoMockRecorder
}

// MockhouseRepoMockRecorder is the mock recorder for MockhouseRepo.
type MockhouseRepoMockRecorder struct {
	mock *MockhouseRepo
}

// NewMockhouseRepo creates a new mock instance.
func NewMockhouseRepo(ctrl *gomock.Controller) *MockhouseRepo {
	mock := &MockhouseRepo{ctrl: ctrl}
	mock.recorder = &MockhouseRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockhouseRepo) EXPECT() *MockhouseRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockhouseRepo) Create(ctx context.Context, house domain.House) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, house)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockhouseRepoMockRecorder) Create(ctx, house any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockhouseRepo)(nil).Create), ctx, house)
}

// GetByID mocks base method.
func (m *MockhouseRepo) GetByID(ctx context.Context, id uint) (domain.House, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(domain.House)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockhouseRepoMockRecorder) GetByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockhouseRepo)(nil).GetByID), ctx, id)
}

// GetFullByID mocks base method.
func (m *MockhouseRepo) GetFullByID(ctx context.Context, id uint, flatStatus *domain.FlatStatus) (domain.House, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullByID", ctx, id, flatStatus)
	ret0, _ := ret[0].(domain.House)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullByID indicates an expected call of GetFullByID.
func (mr *MockhouseRepoMockRecorder) GetFullByID(ctx, id, flatStatus any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullByID", reflect.TypeOf((*MockhouseRepo)(nil).GetFullByID), ctx, id, flatStatus)
}

// UpdateLastFlatAddedAt mocks base method.
func (m *MockhouseRepo) UpdateLastFlatAddedAt(ctx context.Context, id uint, lastFlatAddedAt time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastFlatAddedAt", ctx, id, lastFlatAddedAt)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastFlatAddedAt indicates an expected call of UpdateLastFlatAddedAt.
func (mr *MockhouseRepoMockRecorder) UpdateLastFlatAddedAt(ctx, id, lastFlatAddedAt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastFlatAddedAt", reflect.TypeOf((*MockhouseRepo)(nil).UpdateLastFlatAddedAt), ctx, id, lastFlatAddedAt)
}
