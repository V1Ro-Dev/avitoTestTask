// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecase\pvz-usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	forms "pvz/internal/delivery/forms"
	models "pvz/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPvzRepository is a mock of PvzRepository interface.
type MockPvzRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPvzRepositoryMockRecorder
}

// MockPvzRepositoryMockRecorder is the mock recorder for MockPvzRepository.
type MockPvzRepositoryMockRecorder struct {
	mock *MockPvzRepository
}

// NewMockPvzRepository creates a new mock instance.
func NewMockPvzRepository(ctrl *gomock.Controller) *MockPvzRepository {
	mock := &MockPvzRepository{ctrl: ctrl}
	mock.recorder = &MockPvzRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPvzRepository) EXPECT() *MockPvzRepositoryMockRecorder {
	return m.recorder
}

// CreatePvz mocks base method.
func (m *MockPvzRepository) CreatePvz(ctx context.Context, pvzData models.Pvz) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePvz", ctx, pvzData)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePvz indicates an expected call of CreatePvz.
func (mr *MockPvzRepositoryMockRecorder) CreatePvz(ctx, pvzData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePvz", reflect.TypeOf((*MockPvzRepository)(nil).CreatePvz), ctx, pvzData)
}

// GetPvzInfo mocks base method.
func (m *MockPvzRepository) GetPvzInfo(ctx context.Context, form forms.GetPvzInfoForm) ([]models.PvzInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPvzInfo", ctx, form)
	ret0, _ := ret[0].([]models.PvzInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPvzInfo indicates an expected call of GetPvzInfo.
func (mr *MockPvzRepositoryMockRecorder) GetPvzInfo(ctx, form interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPvzInfo", reflect.TypeOf((*MockPvzRepository)(nil).GetPvzInfo), ctx, form)
}

// GetPvzList mocks base method.
func (m *MockPvzRepository) GetPvzList(ctx context.Context) ([]models.Pvz, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPvzList", ctx)
	ret0, _ := ret[0].([]models.Pvz)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPvzList indicates an expected call of GetPvzList.
func (mr *MockPvzRepositoryMockRecorder) GetPvzList(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPvzList", reflect.TypeOf((*MockPvzRepository)(nil).GetPvzList), ctx)
}
