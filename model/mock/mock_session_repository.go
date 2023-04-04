// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: SessionRepository)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSessionRepository is a mock of SessionRepository interface.
type MockSessionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSessionRepositoryMockRecorder
}

// MockSessionRepositoryMockRecorder is the mock recorder for MockSessionRepository.
type MockSessionRepositoryMockRecorder struct {
	mock *MockSessionRepository
}

// NewMockSessionRepository creates a new mock instance.
func NewMockSessionRepository(ctrl *gomock.Controller) *MockSessionRepository {
	mock := &MockSessionRepository{ctrl: ctrl}
	mock.recorder = &MockSessionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionRepository) EXPECT() *MockSessionRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSessionRepository) Create(arg0 context.Context, arg1 *model.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSessionRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSessionRepository)(nil).Create), arg0, arg1)
}

// FindByID mocks base method.
func (m *MockSessionRepository) FindByID(arg0 context.Context, arg1 int64) (*model.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockSessionRepositoryMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockSessionRepository)(nil).FindByID), arg0, arg1)
}

// FindByRefreshToken mocks base method.
func (m *MockSessionRepository) FindByRefreshToken(arg0 context.Context, arg1 string) (*model.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByRefreshToken", arg0, arg1)
	ret0, _ := ret[0].(*model.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByRefreshToken indicates an expected call of FindByRefreshToken.
func (mr *MockSessionRepositoryMockRecorder) FindByRefreshToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByRefreshToken", reflect.TypeOf((*MockSessionRepository)(nil).FindByRefreshToken), arg0, arg1)
}

// RefreshToken mocks base method.
func (m *MockSessionRepository) RefreshToken(arg0 context.Context, arg1 *model.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockSessionRepositoryMockRecorder) RefreshToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockSessionRepository)(nil).RefreshToken), arg0, arg1)
}
