// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: ConfigurationRepository)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockConfigurationRepository is a mock of ConfigurationRepository interface.
type MockConfigurationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationRepositoryMockRecorder
}

// MockConfigurationRepositoryMockRecorder is the mock recorder for MockConfigurationRepository.
type MockConfigurationRepositoryMockRecorder struct {
	mock *MockConfigurationRepository
}

// NewMockConfigurationRepository creates a new mock instance.
func NewMockConfigurationRepository(ctrl *gomock.Controller) *MockConfigurationRepository {
	mock := &MockConfigurationRepository{ctrl: ctrl}
	mock.recorder = &MockConfigurationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigurationRepository) EXPECT() *MockConfigurationRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockConfigurationRepository) Create(arg0 context.Context, arg1 *model.Configuration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockConfigurationRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockConfigurationRepository)(nil).Create), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockConfigurationRepository) FindAll(arg0 context.Context) ([]*model.Configuration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].([]*model.Configuration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockConfigurationRepositoryMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockConfigurationRepository)(nil).FindAll), arg0)
}

// FindByID mocks base method.
func (m *MockConfigurationRepository) FindByID(arg0 context.Context, arg1 string) (*model.Configuration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Configuration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockConfigurationRepositoryMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockConfigurationRepository)(nil).FindByID), arg0, arg1)
}

// Update mocks base method.
func (m *MockConfigurationRepository) Update(arg0 context.Context, arg1 string, arg2 *model.Configuration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockConfigurationRepositoryMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockConfigurationRepository)(nil).Update), arg0, arg1, arg2)
}