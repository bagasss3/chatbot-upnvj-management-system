// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: ExampleService)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExampleService is a mock of ExampleService interface.
type MockExampleService struct {
	ctrl     *gomock.Controller
	recorder *MockExampleServiceMockRecorder
}

// MockExampleServiceMockRecorder is the mock recorder for MockExampleService.
type MockExampleServiceMockRecorder struct {
	mock *MockExampleService
}

// NewMockExampleService creates a new mock instance.
func NewMockExampleService(ctrl *gomock.Controller) *MockExampleService {
	mock := &MockExampleService{ctrl: ctrl}
	mock.recorder = &MockExampleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExampleService) EXPECT() *MockExampleServiceMockRecorder {
	return m.recorder
}

// CreateExample mocks base method.
func (m *MockExampleService) CreateExample(arg0 context.Context, arg1 model.CreateExampleRequest) (*model.Example, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateExample", arg0, arg1)
	ret0, _ := ret[0].(*model.Example)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateExample indicates an expected call of CreateExample.
func (mr *MockExampleServiceMockRecorder) CreateExample(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateExample", reflect.TypeOf((*MockExampleService)(nil).CreateExample), arg0, arg1)
}

// DeleteExample mocks base method.
func (m *MockExampleService) DeleteExample(arg0 context.Context, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteExample", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteExample indicates an expected call of DeleteExample.
func (mr *MockExampleServiceMockRecorder) DeleteExample(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteExample", reflect.TypeOf((*MockExampleService)(nil).DeleteExample), arg0, arg1, arg2)
}

// FindAllExampleByIntentID mocks base method.
func (m *MockExampleService) FindAllExampleByIntentID(arg0 context.Context, arg1 string) ([]*model.Example, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllExampleByIntentID", arg0, arg1)
	ret0, _ := ret[0].([]*model.Example)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllExampleByIntentID indicates an expected call of FindAllExampleByIntentID.
func (mr *MockExampleServiceMockRecorder) FindAllExampleByIntentID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllExampleByIntentID", reflect.TypeOf((*MockExampleService)(nil).FindAllExampleByIntentID), arg0, arg1)
}

// FindExampleByIntentID mocks base method.
func (m *MockExampleService) FindExampleByIntentID(arg0 context.Context, arg1, arg2 string) (*model.Example, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindExampleByIntentID", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.Example)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindExampleByIntentID indicates an expected call of FindExampleByIntentID.
func (mr *MockExampleServiceMockRecorder) FindExampleByIntentID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindExampleByIntentID", reflect.TypeOf((*MockExampleService)(nil).FindExampleByIntentID), arg0, arg1, arg2)
}

// UpdateExample mocks base method.
func (m *MockExampleService) UpdateExample(arg0 context.Context, arg1, arg2 string, arg3 model.UpdateExampleRequest) (*model.Example, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateExample", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*model.Example)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateExample indicates an expected call of UpdateExample.
func (mr *MockExampleServiceMockRecorder) UpdateExample(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateExample", reflect.TypeOf((*MockExampleService)(nil).UpdateExample), arg0, arg1, arg2, arg3)
}
