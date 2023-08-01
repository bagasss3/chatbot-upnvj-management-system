// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: FacultyRepository)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFacultyRepository is a mock of FacultyRepository interface.
type MockFacultyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockFacultyRepositoryMockRecorder
}

// MockFacultyRepositoryMockRecorder is the mock recorder for MockFacultyRepository.
type MockFacultyRepositoryMockRecorder struct {
	mock *MockFacultyRepository
}

// NewMockFacultyRepository creates a new mock instance.
func NewMockFacultyRepository(ctrl *gomock.Controller) *MockFacultyRepository {
	mock := &MockFacultyRepository{ctrl: ctrl}
	mock.recorder = &MockFacultyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFacultyRepository) EXPECT() *MockFacultyRepositoryMockRecorder {
	return m.recorder
}

// FindAll mocks base method.
func (m *MockFacultyRepository) FindAll(arg0 context.Context) ([]*model.Faculty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0)
	ret0, _ := ret[0].([]*model.Faculty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockFacultyRepositoryMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockFacultyRepository)(nil).FindAll), arg0)
}

// FindByID mocks base method.
func (m *MockFacultyRepository) FindByID(arg0 context.Context, arg1 string) (*model.Faculty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Faculty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockFacultyRepositoryMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockFacultyRepository)(nil).FindByID), arg0, arg1)
}
