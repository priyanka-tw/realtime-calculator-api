// Code generated by MockGen. DO NOT EDIT.
// Source: event_handler_generator.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	"realtime-calculator-api/socket/interface"
	reflect "reflect"
)

// MockEventHandlerGenerator is a mock of EventHandlerGenerator interface
type MockEventHandlerGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockEventHandlerGeneratorMockRecorder
}

// MockEventHandlerGeneratorMockRecorder is the mock recorder for MockEventHandlerGenerator
type MockEventHandlerGeneratorMockRecorder struct {
	mock *MockEventHandlerGenerator
}

// NewMockEventHandlerGenerator creates a new mock instance
func NewMockEventHandlerGenerator(ctrl *gomock.Controller) *MockEventHandlerGenerator {
	mock := &MockEventHandlerGenerator{ctrl: ctrl}
	mock.recorder = &MockEventHandlerGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventHandlerGenerator) EXPECT() *MockEventHandlerGeneratorMockRecorder {
	return m.recorder
}

// GetHandler mocks base method
func (m *MockEventHandlerGenerator) GetHandler(event string) (_interface.EventHandler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandler", event)
	ret0, _ := ret[0].(_interface.EventHandler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHandler indicates an expected call of GetHandler
func (mr *MockEventHandlerGeneratorMockRecorder) GetHandler(event interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandler", reflect.TypeOf((*MockEventHandlerGenerator)(nil).GetHandler), event)
}
