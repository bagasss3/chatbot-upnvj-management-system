// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: FacultyService)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFacultyService is a mock of FacultyService interface.
type MockFacultyService struct {
	ctrl     *gomock.Controller
	recorder *MockFacultyServiceMockRecorder
}

// MockFacultyServiceMockRecorder is the mock recorder for MockFacultyService.
type MockFacultyServiceMockRecorder struct {
	mock *MockFacultyService
}

// NewMockFacultyService creates a new mock instance.
func NewMockFacultyService(ctrl *gomock.Controller) *MockFacultyService {
	mock := &MockFacultyService{ctrl: ctrl}
	mock.recorder = &MockFacultyServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFacultyService) EXPECT() *MockFacultyServiceMockRecorder {
	return m.recorder
}

// FindAllFaculty mocks base method.
func (m *MockFacultyService) FindAllFaculty(arg0 context.Context) ([]*model.Faculty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllFaculty", arg0)
	ret0, _ := ret[0].([]*model.Faculty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllFaculty indicates an expected call of FindAllFaculty.
func (mr *MockFacultyServiceMockRecorder) FindAllFaculty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllFaculty", reflect.TypeOf((*MockFacultyService)(nil).FindAllFaculty), arg0)
}

// FindByIDFaculty mocks base method.
func (m *MockFacultyService) FindByIDFaculty(arg0 context.Context, arg1 string) (*model.Faculty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByIDFaculty", arg0, arg1)
	ret0, _ := ret[0].(*model.Faculty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByIDFaculty indicates an expected call of FindByIDFaculty.
func (mr *MockFacultyServiceMockRecorder) FindByIDFaculty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByIDFaculty", reflect.TypeOf((*MockFacultyService)(nil).FindByIDFaculty), arg0, arg1)
}
