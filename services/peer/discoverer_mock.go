// Code generated by MockGen. DO NOT EDIT.
// Source: services/peer/service.go

// Package peer is a generated GoMock package.
package peer

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDiscoverer is a mock of Discoverer interface
type MockDiscoverer struct {
	ctrl     *gomock.Controller
	recorder *MockDiscovererMockRecorder
}

// MockDiscovererMockRecorder is the mock recorder for MockDiscoverer
type MockDiscovererMockRecorder struct {
	mock *MockDiscoverer
}

// NewMockDiscoverer creates a new mock instance
func NewMockDiscoverer(ctrl *gomock.Controller) *MockDiscoverer {
	mock := &MockDiscoverer{ctrl: ctrl}
	mock.recorder = &MockDiscovererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDiscoverer) EXPECT() *MockDiscovererMockRecorder {
	return m.recorder
}

// Discover mocks base method
func (m *MockDiscoverer) Discover(topic string, max, min int) error {
	ret := m.ctrl.Call(m, "Discover", topic, max, min)
	ret0, _ := ret[0].(error)
	return ret0
}

// Discover indicates an expected call of Discover
func (mr *MockDiscovererMockRecorder) Discover(topic, max, min interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Discover", reflect.TypeOf((*MockDiscoverer)(nil).Discover), topic, max, min)
}
