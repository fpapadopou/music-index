// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/elastic/go-elasticsearch/v7/esapi (interfaces: Transport)

// Package elastic is a generated GoMock package.
package elastic

import (
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockTransport is a mock of Transport interface
type MockTransport struct {
	ctrl     *gomock.Controller
	recorder *MockTransportMockRecorder
}

// MockTransportMockRecorder is the mock recorder for MockTransport
type MockTransportMockRecorder struct {
	mock *MockTransport
}

// NewMockTransport creates a new mock instance
func NewMockTransport(ctrl *gomock.Controller) *MockTransport {
	mock := &MockTransport{ctrl: ctrl}
	mock.recorder = &MockTransportMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransport) EXPECT() *MockTransportMockRecorder {
	return m.recorder
}

// Perform mocks base method
func (m *MockTransport) Perform(arg0 *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Perform", arg0)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Perform indicates an expected call of Perform
func (mr *MockTransportMockRecorder) Perform(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Perform", reflect.TypeOf((*MockTransport)(nil).Perform), arg0)
}
