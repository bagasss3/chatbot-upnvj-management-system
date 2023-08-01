// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: ReqBodyService)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReqBodyService is a mock of ReqBodyService interface.
type MockReqBodyService struct {
	ctrl     *gomock.Controller
	recorder *MockReqBodyServiceMockRecorder
}

// MockReqBodyServiceMockRecorder is the mock recorder for MockReqBodyService.
type MockReqBodyServiceMockRecorder struct {
	mock *MockReqBodyService
}

// NewMockReqBodyService creates a new mock instance.
func NewMockReqBodyService(ctrl *gomock.Controller) *MockReqBodyService {
	mock := &MockReqBodyService{ctrl: ctrl}
	mock.recorder = &MockReqBodyServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReqBodyService) EXPECT() *MockReqBodyServiceMockRecorder {
	return m.recorder
}

// CreateReqBody mocks base method.
func (m *MockReqBodyService) CreateReqBody(arg0 context.Context, arg1 model.CreateReqBodyActionArrayRequest) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReqBody", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateReqBody indicates an expected call of CreateReqBody.
func (mr *MockReqBodyServiceMockRecorder) CreateReqBody(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReqBody", reflect.TypeOf((*MockReqBodyService)(nil).CreateReqBody), arg0, arg1)
}

// DeleteReqBody mocks base method.
func (m *MockReqBodyService) DeleteReqBody(arg0 context.Context, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteReqBody", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteReqBody indicates an expected call of DeleteReqBody.
func (mr *MockReqBodyServiceMockRecorder) DeleteReqBody(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteReqBody", reflect.TypeOf((*MockReqBodyService)(nil).DeleteReqBody), arg0, arg1, arg2)
}

// FindAllReqBodyByActionHttpID mocks base method.
func (m *MockReqBodyService) FindAllReqBodyByActionHttpID(arg0 context.Context, arg1, arg2 string) ([]*model.ReqBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllReqBodyByActionHttpID", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*model.ReqBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllReqBodyByActionHttpID indicates an expected call of FindAllReqBodyByActionHttpID.
func (mr *MockReqBodyServiceMockRecorder) FindAllReqBodyByActionHttpID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllReqBodyByActionHttpID", reflect.TypeOf((*MockReqBodyService)(nil).FindAllReqBodyByActionHttpID), arg0, arg1, arg2)
}

// FindByID mocks base method.
func (m *MockReqBodyService) FindByID(arg0 context.Context, arg1 string) (*model.ReqBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*model.ReqBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockReqBodyServiceMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockReqBodyService)(nil).FindByID), arg0, arg1)
}

// UpdateReqBody mocks base method.
func (m *MockReqBodyService) UpdateReqBody(arg0 context.Context, arg1, arg2 string, arg3 model.UpdateReqBodyRequest) (*model.ReqBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateReqBody", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*model.ReqBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateReqBody indicates an expected call of UpdateReqBody.
func (mr *MockReqBodyServiceMockRecorder) UpdateReqBody(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateReqBody", reflect.TypeOf((*MockReqBodyService)(nil).UpdateReqBody), arg0, arg1, arg2, arg3)
}