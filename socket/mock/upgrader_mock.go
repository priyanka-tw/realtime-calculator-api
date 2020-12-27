// Code generated by MockGen. DO NOT EDIT.
// Source: socket/upgrader.go

// Package mocks is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	websocket "github.com/gorilla/websocket"
	http "net/http"
	reflect "reflect"
)

// MockUpWrapper is a mock of UpgraderWrapper interface
type MockUpWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockUpWrapperMockRecorder
}

// MockUpWrapperMockRecorder is the mock recorder for MockUpWrapper
type MockUpWrapperMockRecorder struct {
	mock *MockUpWrapper
}

// NewMockUpWrapper creates a new mock instance
func NewMockUpWrapper(ctrl *gomock.Controller) *MockUpWrapper {
	mock := &MockUpWrapper{ctrl: ctrl}
	mock.recorder = &MockUpWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUpWrapper) EXPECT() *MockUpWrapperMockRecorder {
	return m.recorder
}

// Upgrade mocks base method
func (m *MockUpWrapper) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upgrade", w, r, responseHeader)
	ret0, _ := ret[0].(*websocket.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upgrade indicates an expected call of Upgrade
func (mr *MockUpWrapperMockRecorder) Upgrade(w, r, responseHeader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upgrade", reflect.TypeOf((*MockUpWrapper)(nil).Upgrade), w, r, responseHeader)
}