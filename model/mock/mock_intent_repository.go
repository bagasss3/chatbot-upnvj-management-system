// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: IntentRepository)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockIntentRepository is a mock of IntentRepository interface.
type MockIntentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIntentRepositoryMockRecorder
}

// MockIntentRepositoryMockRecorder is the mock recorder for MockIntentRepository.
type MockIntentRepositoryMockRecorder struct {
	mock *MockIntentRepository
}

// NewMockIntentRepository creates a new mock instance.
func NewMockIntentRepository(ctrl *gomock.Controller) *MockIntentRepository {
	mock := &MockIntentRepository{ctrl: ctrl}
	mock.recorder = &MockIntentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntentRepository) EXPECT() *MockIntentRepositoryMockRecorder {
	return m.recorder
}

// CountAll mocks base method.
func (m *MockIntentRepository) CountAll(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAll", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountAll indicates an expected call of CountAll.
func (mr *MockIntentRepositoryMockRecorder) CountAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAll", reflect.TypeOf((*MockIntentRepository)(nil).CountAll), arg0)
}

// Create mocks base method.
func (m *MockIntentRepository) Create(arg0 context.Context, arg1 *model.Intent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockIntentRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIntentRepository)(nil).Create), arg0, arg1)
}

// DeleteWithTx mocks base method.
func (m *MockIntentRepository) DeleteWithTx(arg0 context.Context, arg1 string, arg2 *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWithTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteWithTx indicates an expected call of DeleteWithTx.
func (mr *MockIntentRepositoryMockRecorder) DeleteWithTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWithTx", reflect.TypeOf((*MockIntentRepository)(nil).DeleteWithTx), arg0, arg1, arg2)
}

// FindAll mocks base method.
func (m *MockIntentRepository) FindAll(arg0 context.Context, arg1 string) ([]*model.Intent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", arg0, arg1)
	ret0, _ := ret[0].([]*model.Intent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockIntentRepositoryMockRecorder) FindAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIntentRepository)(nil).FindAll), arg0, arg1)
}

// FindAllInformationAcademics mocks base method.
func (m *MockIntentRepository) FindAllInformationAcademics(arg0 context.Context) ([]*model.Intent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllInformationAcademics", arg0)
	ret0, _ := ret[0].([]*model.Intent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllInformationAcademics indicates an expected call of FindAllInformationAcademics.
func (mr *MockIntentRepositoryMockRecorder) FindAllInformationAcademics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllInformationAcademics", reflect.TypeOf((*MockIntentRepository)(nil).FindAllInformationAcademics), arg0)
}

// FindAllWithExamples mocks base method.
func (m *MockIntentRepository) FindAllWithExamples(arg0 context.Context) ([]*model.Intent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllWithExamples", arg0)
	ret0, _ := ret[0].([]*model.Intent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllWithExamples indicates an expected call of FindAllWithExamples.
func (mr *MockIntentRepositoryMockRecorder) FindAllWithExamples(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllWithExamples", reflect.TypeOf((*MockIntentRepository)(nil).FindAllWithExamples), arg0)
}

// FindByID mocks base method.
func (m *MockIntentRepository) FindByID(arg0 context.Context, arg1 string) (*model.Intent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Intent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockIntentRepositoryMockRecorder) FindByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockIntentRepository)(nil).FindByID), arg0, arg1)
}

// FindByName mocks base method.
func (m *MockIntentRepository) FindByName(arg0 context.Context, arg1 string) (*model.Intent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0, arg1)
	ret0, _ := ret[0].(*model.Intent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName.
func (mr *MockIntentRepositoryMockRecorder) FindByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockIntentRepository)(nil).FindByName), arg0, arg1)
}

// Update mocks base method.
func (m *MockIntentRepository) Update(arg0 context.Context, arg1 string, arg2 *model.Intent) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIntentRepositoryMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIntentRepository)(nil).Update), arg0, arg1, arg2)
}
