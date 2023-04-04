// Code generated by MockGen. DO NOT EDIT.
// Source: cbupnvj/model (interfaces: IntentController)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockIntentController is a mock of IntentController interface.
type MockIntentController struct {
	ctrl     *gomock.Controller
	recorder *MockIntentControllerMockRecorder
}

// MockIntentControllerMockRecorder is the mock recorder for MockIntentController.
type MockIntentControllerMockRecorder struct {
	mock *MockIntentController
}

// NewMockIntentController creates a new mock instance.
func NewMockIntentController(ctrl *gomock.Controller) *MockIntentController {
	mock := &MockIntentController{ctrl: ctrl}
	mock.recorder = &MockIntentControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIntentController) EXPECT() *MockIntentControllerMockRecorder {
	return m.recorder
}

// HandleCreateIntent mocks base method.
func (m *MockIntentController) HandleCreateIntent() echo.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleCreateIntent")
	ret0, _ := ret[0].(echo.HandlerFunc)
	return ret0
}

// HandleCreateIntent indicates an expected call of HandleCreateIntent.
func (mr *MockIntentControllerMockRecorder) HandleCreateIntent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleCreateIntent", reflect.TypeOf((*MockIntentController)(nil).HandleCreateIntent))
}

// HandleDeleteIntent mocks base method.
func (m *MockIntentController) HandleDeleteIntent() echo.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleDeleteIntent")
	ret0, _ := ret[0].(echo.HandlerFunc)
	return ret0
}

// HandleDeleteIntent indicates an expected call of HandleDeleteIntent.
func (mr *MockIntentControllerMockRecorder) HandleDeleteIntent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleDeleteIntent", reflect.TypeOf((*MockIntentController)(nil).HandleDeleteIntent))
}

// HandleFindAllIntent mocks base method.
func (m *MockIntentController) HandleFindAllIntent() echo.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleFindAllIntent")
	ret0, _ := ret[0].(echo.HandlerFunc)
	return ret0
}

// HandleFindAllIntent indicates an expected call of HandleFindAllIntent.
func (mr *MockIntentControllerMockRecorder) HandleFindAllIntent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleFindAllIntent", reflect.TypeOf((*MockIntentController)(nil).HandleFindAllIntent))
}

// HandleFindIntentByID mocks base method.
func (m *MockIntentController) HandleFindIntentByID() echo.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleFindIntentByID")
	ret0, _ := ret[0].(echo.HandlerFunc)
	return ret0
}

// HandleFindIntentByID indicates an expected call of HandleFindIntentByID.
func (mr *MockIntentControllerMockRecorder) HandleFindIntentByID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleFindIntentByID", reflect.TypeOf((*MockIntentController)(nil).HandleFindIntentByID))
}

// HandleUpdateIntent mocks base method.
func (m *MockIntentController) HandleUpdateIntent() echo.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleUpdateIntent")
	ret0, _ := ret[0].(echo.HandlerFunc)
	return ret0
}

// HandleUpdateIntent indicates an expected call of HandleUpdateIntent.
func (mr *MockIntentControllerMockRecorder) HandleUpdateIntent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleUpdateIntent", reflect.TypeOf((*MockIntentController)(nil).HandleUpdateIntent))
}
