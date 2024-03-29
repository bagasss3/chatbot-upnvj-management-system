// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: RuleService)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRuleService is a mock of RuleService interface.
type MockRuleService struct {
	ctrl     *gomock.Controller
	recorder *MockRuleServiceMockRecorder
}

// MockRuleServiceMockRecorder is the mock recorder for MockRuleService.
type MockRuleServiceMockRecorder struct {
	mock *MockRuleService
}

// NewMockRuleService creates a new mock instance.
func NewMockRuleService(ctrl *gomock.Controller) *MockRuleService {
	mock := &MockRuleService{ctrl: ctrl}
	mock.recorder = &MockRuleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRuleService) EXPECT() *MockRuleServiceMockRecorder {
	return m.recorder
}

// CreateRule mocks base method.
func (m *MockRuleService) CreateRule(arg0 context.Context, arg1 model.CreateUpdateRuleRequest) (*model.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRule", arg0, arg1)
	ret0, _ := ret[0].(*model.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRule indicates an expected call of CreateRule.
func (mr *MockRuleServiceMockRecorder) CreateRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRule", reflect.TypeOf((*MockRuleService)(nil).CreateRule), arg0, arg1)
}

// DeleteRule mocks base method.
func (m *MockRuleService) DeleteRule(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRule", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRule indicates an expected call of DeleteRule.
func (mr *MockRuleServiceMockRecorder) DeleteRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRule", reflect.TypeOf((*MockRuleService)(nil).DeleteRule), arg0, arg1)
}

// FindAllRule mocks base method.
func (m *MockRuleService) FindAllRule(arg0 context.Context, arg1 string) ([]*model.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllRule", arg0, arg1)
	ret0, _ := ret[0].([]*model.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllRule indicates an expected call of FindAllRule.
func (mr *MockRuleServiceMockRecorder) FindAllRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllRule", reflect.TypeOf((*MockRuleService)(nil).FindAllRule), arg0, arg1)
}

// FindRuleByID mocks base method.
func (m *MockRuleService) FindRuleByID(arg0 context.Context, arg1 string) (*model.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRuleByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRuleByID indicates an expected call of FindRuleByID.
func (mr *MockRuleServiceMockRecorder) FindRuleByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRuleByID", reflect.TypeOf((*MockRuleService)(nil).FindRuleByID), arg0, arg1)
}

// UpdateRule mocks base method.
func (m *MockRuleService) UpdateRule(arg0 context.Context, arg1 string, arg2 model.CreateUpdateRuleRequest) (*model.Rule, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRule", arg0, arg1, arg2)
	ret0, _ := ret[0].(*model.Rule)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRule indicates an expected call of UpdateRule.
func (mr *MockRuleServiceMockRecorder) UpdateRule(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRule", reflect.TypeOf((*MockRuleService)(nil).UpdateRule), arg0, arg1, arg2)
}
