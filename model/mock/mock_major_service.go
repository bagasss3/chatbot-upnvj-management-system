// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: MajorService)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMajorService is a mock of MajorService interface.
type MockMajorService struct {
	ctrl     *gomock.Controller
	recorder *MockMajorServiceMockRecorder
}

// MockMajorServiceMockRecorder is the mock recorder for MockMajorService.
type MockMajorServiceMockRecorder struct {
	mock *MockMajorService
}

// NewMockMajorService creates a new mock instance.
func NewMockMajorService(ctrl *gomock.Controller) *MockMajorService {
	mock := &MockMajorService{ctrl: ctrl}
	mock.recorder = &MockMajorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMajorService) EXPECT() *MockMajorServiceMockRecorder {
	return m.recorder
}

// FindAllMajor mocks base method.
func (m *MockMajorService) FindAllMajor(arg0 context.Context) ([]*model.Major, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllMajor", arg0)
	ret0, _ := ret[0].([]*model.Major)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllMajor indicates an expected call of FindAllMajor.
func (mr *MockMajorServiceMockRecorder) FindAllMajor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllMajor", reflect.TypeOf((*MockMajorService)(nil).FindAllMajor), arg0)
}

// FindByIDMajor mocks base method.
func (m *MockMajorService) FindByIDMajor(arg0 context.Context, arg1 string) (*model.Major, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDMajor", arg0, arg1)
	ret0, _ := ret[0].(*model.Major)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDMajor indicates an expected call of FindByIDMajor.
func (mr *MockMajorServiceMockRecorder) FindByIDMajor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDMajor", reflect.TypeOf((*MockMajorService)(nil).FindByIDMajor), arg0, arg1)
}
