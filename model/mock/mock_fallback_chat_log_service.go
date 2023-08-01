// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: FallbackChatLogService)

// Package mock is a generated GoMock package.
package mock

import (
	model "cbupnvj/model"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFallbackChatLogService is a mock of FallbackChatLogService interface.
type MockFallbackChatLogService struct {
	ctrl     *gomock.Controller
	recorder *MockFallbackChatLogServiceMockRecorder
}

// MockFallbackChatLogServiceMockRecorder is the mock recorder for MockFallbackChatLogService.
type MockFallbackChatLogServiceMockRecorder struct {
	mock *MockFallbackChatLogService
}

// NewMockFallbackChatLogService creates a new mock instance.
func NewMockFallbackChatLogService(ctrl *gomock.Controller) *MockFallbackChatLogService {
	mock := &MockFallbackChatLogService{ctrl: ctrl}
	mock.recorder = &MockFallbackChatLogServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFallbackChatLogService) EXPECT() *MockFallbackChatLogServiceMockRecorder {
	return m.recorder
}

// CreateFallbackChatLog mocks base method.
func (m *MockFallbackChatLogService) CreateFallbackChatLog(arg0 context.Context, arg1 model.CreateFallbackChatLogRequest) (*model.FallbackChatLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFallbackChatLog", arg0, arg1)
	ret0, _ := ret[0].(*model.FallbackChatLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFallbackChatLog indicates an expected call of CreateFallbackChatLog.
func (mr *MockFallbackChatLogServiceMockRecorder) CreateFallbackChatLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFallbackChatLog", reflect.TypeOf((*MockFallbackChatLogService)(nil).CreateFallbackChatLog), arg0, arg1)
}

// FindAllFallbackChatLog mocks base method.
func (m *MockFallbackChatLogService) FindAllFallbackChatLog(arg0 context.Context, arg1 string) ([]*model.FallbackChatLog, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllFallbackChatLog", arg0, arg1)
	ret0, _ := ret[0].([]*model.FallbackChatLog)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllFallbackChatLog indicates an expected call of FindAllFallbackChatLog.
func (mr *MockFallbackChatLogServiceMockRecorder) FindAllFallbackChatLog(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllFallbackChatLog", reflect.TypeOf((*MockFallbackChatLogService)(nil).FindAllFallbackChatLog), arg0, arg1)
}

// FindAllFallbackChatLogGroupCluster mocks base method.
func (m *MockFallbackChatLogService) FindAllFallbackChatLogGroupCluster(arg0 context.Context) ([]*model.ClusterData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllFallbackChatLogGroupCluster", arg0)
	ret0, _ := ret[0].([]*model.ClusterData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllFallbackChatLogGroupCluster indicates an expected call of FindAllFallbackChatLogGroupCluster.
func (mr *MockFallbackChatLogServiceMockRecorder) FindAllFallbackChatLogGroupCluster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllFallbackChatLogGroupCluster", reflect.TypeOf((*MockFallbackChatLogService)(nil).FindAllFallbackChatLogGroupCluster), arg0)
}

// FindAllFallbackChatLogOldAndNew mocks base method.
func (m *MockFallbackChatLogService) FindAllFallbackChatLogOldAndNew(arg0 context.Context) (*model.ResponseFallback, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllFallbackChatLogOldAndNew", arg0)
	ret0, _ := ret[0].(*model.ResponseFallback)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllFallbackChatLogOldAndNew indicates an expected call of FindAllFallbackChatLogOldAndNew.
func (mr *MockFallbackChatLogServiceMockRecorder) FindAllFallbackChatLogOldAndNew(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllFallbackChatLogOldAndNew", reflect.TypeOf((*MockFallbackChatLogService)(nil).FindAllFallbackChatLogOldAndNew), arg0)
}

// UpdateGroupCluster mocks base method.
func (m *MockFallbackChatLogService) UpdateGroupCluster(arg0 context.Context) ([]*model.ClusterData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroupCluster", arg0)
	ret0, _ := ret[0].([]*model.ClusterData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGroupCluster indicates an expected call of UpdateGroupCluster.
func (mr *MockFallbackChatLogServiceMockRecorder) UpdateGroupCluster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroupCluster", reflect.TypeOf((*MockFallbackChatLogService)(nil).UpdateGroupCluster), arg0)
}
