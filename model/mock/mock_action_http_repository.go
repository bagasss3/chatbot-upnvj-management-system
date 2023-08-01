// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: ActionHttpRepository)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockActionHttpRepository is a mock of ActionHttpRepository interface.
type MockActionHttpRepository struct {
	ctrl     *gomock.Controller
	recorder *MockActionHttpRepositoryMockRecorder
}

// MockActionHttpRepositoryMockRecorder is the mock recorder for MockActionHttpRepository.
type MockActionHttpRepositoryMockRecorder struct {
	mock *MockActionHttpRepository
}

// NewMockActionHttpRepository creates a new mock instance.
func NewMockActionHttpRepository(ctrl *gomock.Controller) *MockActionHttpRepository {
	mock := &MockActionHttpRepository{ctrl: ctrl}
	mock.recorder = &MockActionHttpRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActionHttpRepository) EXPECT() *MockActionHttpRepositoryMockRecorder {
	return m.recorder
}

// CountAll mocks base method.
func (m *MockActionHttpRepository) CountAll(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAll", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountAll indicates an expected call of CountAll.
func (mr *MockActionHttpRepositoryMockRecorder) CountAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAll", reflect.TypeOf((*MockActionHttpRepository)(nil).CountAll), arg0)
}

// Create mocks base method.
func (m *MockActionHttpRepository) Create(arg0 context.Context, arg1 *model.ActionHttp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockActionHttpRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockActionHttpRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockActionHttpRepository) Delete(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockActionHttpRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockActionHttpRepository)(nil).Delete), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockActionHttpRepository) FindAll(arg0 context.Context, arg1 string) ([]*model.ActionHttp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0, arg1)
	ret0, _ := ret[0].([]*model.ActionHttp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockActionHttpRepositoryMockRecorder) FindAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockActionHttpRepository)(nil).FindAll), arg0, arg1)
}

// FindAllWithReqBodies mocks base method.
func (m *MockActionHttpRepository) FindAllWithReqBodies(arg0 context.Context) ([]*model.ActionHttp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllWithReqBodies", arg0)
	ret0, _ := ret[0].([]*model.ActionHttp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllWithReqBodies indicates an expected call of FindAllWithReqBodies.
func (mr *MockActionHttpRepositoryMockRecorder) FindAllWithReqBodies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllWithReqBodies", reflect.TypeOf((*MockActionHttpRepository)(nil).FindAllWithReqBodies), arg0)
}

// FindByID mocks base method.
func (m *MockActionHttpRepository) FindByID(arg0 context.Context, arg1 string) (*model.ActionHttp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*model.ActionHttp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockActionHttpRepositoryMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockActionHttpRepository)(nil).FindByID), arg0, arg1)
}

// Update mocks base method.
func (m *MockActionHttpRepository) Update(arg0 context.Context, arg1 *model.ActionHttp) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockActionHttpRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockActionHttpRepository)(nil).Update), arg0, arg1)
}