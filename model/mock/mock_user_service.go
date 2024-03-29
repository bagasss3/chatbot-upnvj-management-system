// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: UserService)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CreateAdmin mocks base method.
func (m *MockUserService) CreateAdmin(arg0 context.Context, arg1 model.CreateAdminRequest) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdmin", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAdmin indicates an expected call of CreateAdmin.
func (mr *MockUserServiceMockRecorder) CreateAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockUserService)(nil).CreateAdmin), arg0, arg1)
}

// DeleteAdminByID mocks base method.
func (m *MockUserService) DeleteAdminByID(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAdminByID", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAdminByID indicates an expected call of DeleteAdminByID.
func (mr *MockUserServiceMockRecorder) DeleteAdminByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdminByID", reflect.TypeOf((*MockUserService)(nil).DeleteAdminByID), arg0, arg1)
}

// FindAdminByID mocks base method.
func (m *MockUserService) FindAdminByID(arg0 context.Context, arg1 string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAdminByID", arg0, arg1)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAdminByID indicates an expected call of FindAdminByID.
func (mr *MockUserServiceMockRecorder) FindAdminByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAdminByID", reflect.TypeOf((*MockUserService)(nil).FindAdminByID), arg0, arg1)
}

// FindAllAdmin mocks base method.
func (m *MockUserService) FindAllAdmin(arg0 context.Context) ([]*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllAdmin", arg0)
	ret0, _ := ret[0].([]*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllAdmin indicates an expected call of FindAllAdmin.
func (mr *MockUserServiceMockRecorder) FindAllAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllAdmin", reflect.TypeOf((*MockUserService)(nil).FindAllAdmin), arg0)
}

// UpdateAdmin mocks base method.
func (m *MockUserService) UpdateAdmin(arg0 context.Context, arg1 string, arg2 model.UpdateAdminRequest) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdmin", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAdmin indicates an expected call of UpdateAdmin.
func (mr *MockUserServiceMockRecorder) UpdateAdmin(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdmin", reflect.TypeOf((*MockUserService)(nil).UpdateAdmin), arg0, arg1, arg2)
}

// UpdateProfile mocks base method.
func (m *MockUserService) UpdateProfile(arg0 context.Context, arg1 string, arg2 model.UpdateUserPasswordRequest) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockUserServiceMockRecorder) UpdateProfile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockUserService)(nil).UpdateProfile), arg0, arg1, arg2)
}
