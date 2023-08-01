// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: ActionHttpService)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockActionHttpService is a mock of ActionHttpService interface.
type MockActionHttpService struct {
	ctrl     *gomock.Controller
	recorder *MockActionHttpServiceMockRecorder
}

// MockActionHttpServiceMockRecorder is the mock recorder for MockActionHttpService.
type MockActionHttpServiceMockRecorder struct {
	mock *MockActionHttpService
}

// NewMockActionHttpService creates a new mock instance.
func NewMockActionHttpService(ctrl *gomock.Controller) *MockActionHttpService {
	mock := &MockActionHttpService{ctrl: ctrl}
	mock.recorder = &MockActionHttpServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActionHttpService) EXPECT() *MockActionHttpServiceMockRecorder {
	return m.recorder
}

// CountAllActionHttp mocks base method.
func (m *MockActionHttpService) CountAllActionHttp(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAllActionHttp", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountAllActionHttp indicates an expected call of CountAllActionHttp.
func (mr *MockActionHttpServiceMockRecorder) CountAllActionHttp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAllActionHttp", reflect.TypeOf((*MockActionHttpService)(nil).CountAllActionHttp), arg0)
}

// CreateActionHttp mocks base method.
func (m *MockActionHttpService) CreateActionHttp(arg0 context.Context, arg1 model.CreateUpdateActionHttpRequest) (*model.ActionHttp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateActionHttp", arg0, arg1)
	ret0, _ := ret[0].(*model.ActionHttp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateActionHttp indicates an expected call of CreateActionHttp.
func (mr *MockActionHttpServiceMockRecorder) CreateActionHttp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActionHttp", reflect.TypeOf((*MockActionHttpService)(nil).CreateActionHttp), arg0, arg1)
}

// DeleteActionHttp mocks base method.
func (m *MockActionHttpService) DeleteActionHttp(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteActionHttp", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteActionHttp indicates an expected call of DeleteActionHttp.
func (mr *MockActionHttpServiceMockRecorder) DeleteActionHttp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteActionHttp", reflect.TypeOf((*MockActionHttpService)(nil).DeleteActionHttp), arg0, arg1)
}

// FindActionHttpByID mocks base method.
func (m *MockActionHttpService) FindActionHttpByID(arg0 context.Context, arg1 string) (*model.ActionHttp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindActionHttpByID", arg0, arg1)
	ret0, _ := ret[0].(*model.ActionHttp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindActionHttpByID indicates an expected call of FindActionHttpByID.
func (mr *MockActionHttpServiceMockRecorder) FindActionHttpByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindActionHttpByID", reflect.TypeOf((*MockActionHttpService)(nil).FindActionHttpByID), arg0, arg1)
}

// FindAllActionHttp mocks base method.
func (m *MockActionHttpService) FindAllActionHttp(arg0 context.Context, arg1 string) ([]*model.ActionHttp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllActionHttp", arg0, arg1)
	ret0, _ := ret[0].([]*model.ActionHttp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllActionHttp indicates an expected call of FindAllActionHttp.
func (mr *MockActionHttpServiceMockRecorder) FindAllActionHttp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllActionHttp", reflect.TypeOf((*MockActionHttpService)(nil).FindAllActionHttp), arg0, arg1)
}

// FindAllWithReqBodies mocks base method.
func (m *MockActionHttpService) FindAllWithReqBodies(arg0 context.Context) ([]*model.ActionHttp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllWithReqBodies", arg0)
	ret0, _ := ret[0].([]*model.ActionHttp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllWithReqBodies indicates an expected call of FindAllWithReqBodies.
func (mr *MockActionHttpServiceMockRecorder) FindAllWithReqBodies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllWithReqBodies", reflect.TypeOf((*MockActionHttpService)(nil).FindAllWithReqBodies), arg0)
}

// UpdateActionHttp mocks base method.
func (m *MockActionHttpService) UpdateActionHttp(arg0 context.Context, arg1 string, arg2 model.CreateUpdateActionHttpRequest) (*model.ActionHttp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateActionHttp", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.ActionHttp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateActionHttp indicates an expected call of UpdateActionHttp.
func (mr *MockActionHttpServiceMockRecorder) UpdateActionHttp(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActionHttp", reflect.TypeOf((*MockActionHttpService)(nil).UpdateActionHttp), arg0, arg1, arg2)
}