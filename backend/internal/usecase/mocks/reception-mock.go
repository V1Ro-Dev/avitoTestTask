// Code generated by MockGen. DO NOT EDIT.
// Source: internal/usecase\reception-usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	models "pvz/internal/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockReceptionRepository is a mock of ReceptionRepository interface.
type MockReceptionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockReceptionRepositoryMockRecorder
}

// MockReceptionRepositoryMockRecorder is the mock recorder for MockReceptionRepository.
type MockReceptionRepositoryMockRecorder struct {
	mock *MockReceptionRepository
}

// NewMockReceptionRepository creates a new mock instance.
func NewMockReceptionRepository(ctrl *gomock.Controller) *MockReceptionRepository {
	mock := &MockReceptionRepository{ctrl: ctrl}
	mock.recorder = &MockReceptionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReceptionRepository) EXPECT() *MockReceptionRepositoryMockRecorder {
	return m.recorder
}

// AddProduct mocks base method.
func (m *MockReceptionRepository) AddProduct(ctx context.Context, product models.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProduct", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProduct indicates an expected call of AddProduct.
func (mr *MockReceptionRepositoryMockRecorder) AddProduct(ctx, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProduct", reflect.TypeOf((*MockReceptionRepository)(nil).AddProduct), ctx, product)
}

// CloseReception mocks base method.
func (m *MockReceptionRepository) CloseReception(ctx context.Context, receptionData models.Reception) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseReception", ctx, receptionData)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseReception indicates an expected call of CloseReception.
func (mr *MockReceptionRepositoryMockRecorder) CloseReception(ctx, receptionData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseReception", reflect.TypeOf((*MockReceptionRepository)(nil).CloseReception), ctx, receptionData)
}

// CreateReception mocks base method.
func (m *MockReceptionRepository) CreateReception(ctx context.Context, receptionData models.Reception) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReception", ctx, receptionData)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateReception indicates an expected call of CreateReception.
func (mr *MockReceptionRepositoryMockRecorder) CreateReception(ctx, receptionData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReception", reflect.TypeOf((*MockReceptionRepository)(nil).CreateReception), ctx, receptionData)
}

// GetOpenReception mocks base method.
func (m *MockReceptionRepository) GetOpenReception(ctx context.Context, pvzId uuid.UUID) (models.Reception, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenReception", ctx, pvzId)
	ret0, _ := ret[0].(models.Reception)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOpenReception indicates an expected call of GetOpenReception.
func (mr *MockReceptionRepositoryMockRecorder) GetOpenReception(ctx, pvzId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenReception", reflect.TypeOf((*MockReceptionRepository)(nil).GetOpenReception), ctx, pvzId)
}

// RemoveProduct mocks base method.
func (m *MockReceptionRepository) RemoveProduct(ctx context.Context, receptionId uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProduct", ctx, receptionId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProduct indicates an expected call of RemoveProduct.
func (mr *MockReceptionRepositoryMockRecorder) RemoveProduct(ctx, receptionId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProduct", reflect.TypeOf((*MockReceptionRepository)(nil).RemoveProduct), ctx, receptionId)
}
