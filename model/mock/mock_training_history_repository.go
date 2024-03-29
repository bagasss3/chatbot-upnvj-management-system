// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: TrainingHistoryRepository)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTrainingHistoryRepository is a mock of TrainingHistoryRepository interface.
type MockTrainingHistoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTrainingHistoryRepositoryMockRecorder
}

// MockTrainingHistoryRepositoryMockRecorder is the mock recorder for MockTrainingHistoryRepository.
type MockTrainingHistoryRepositoryMockRecorder struct {
	mock *MockTrainingHistoryRepository
}

// NewMockTrainingHistoryRepository creates a new mock instance.
func NewMockTrainingHistoryRepository(ctrl *gomock.Controller) *MockTrainingHistoryRepository {
	mock := &MockTrainingHistoryRepository{ctrl: ctrl}
	mock.recorder = &MockTrainingHistoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrainingHistoryRepository) EXPECT() *MockTrainingHistoryRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTrainingHistoryRepository) Create(arg0 context.Context, arg1 *model.TrainingHistory) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTrainingHistoryRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTrainingHistoryRepository)(nil).Create), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockTrainingHistoryRepository) FindAll(arg0 context.Context) ([]*model.TrainingHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].([]*model.TrainingHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockTrainingHistoryRepositoryMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockTrainingHistoryRepository)(nil).FindAll), arg0)
}
