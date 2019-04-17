// Code generated by MockGen. DO NOT EDIT.
// Source: ./mux.go

// Package bimux is a generated GoMock package.
package bimux

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockMuxer is a mock of Muxer interface
type MockMuxer struct {
	ctrl     *gomock.Controller
	recorder *MockMuxerMockRecorder
}

// MockMuxerMockRecorder is the mock recorder for MockMuxer
type MockMuxerMockRecorder struct {
	mock *MockMuxer
}

// NewMockMuxer creates a new mock instance
func NewMockMuxer(ctrl *gomock.Controller) *MockMuxer {
	mock := &MockMuxer{ctrl: ctrl}
	mock.recorder = &MockMuxerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMuxer) EXPECT() *MockMuxerMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockMuxer) Send(route int32, data []byte) error {
	ret := m.ctrl.Call(m, "Send", route, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockMuxerMockRecorder) Send(route, data interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockMuxer)(nil).Send), route, data)
}

// Rpc mocks base method
func (m *MockMuxer) Rpc(route int32, req []byte, timeout time.Duration) ([]byte, error) {
	ret := m.ctrl.Call(m, "Rpc", route, req, timeout)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rpc indicates an expected call of Rpc
func (mr *MockMuxerMockRecorder) Rpc(route, req, timeout interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rpc", reflect.TypeOf((*MockMuxer)(nil).Rpc), route, req, timeout)
}

// Wait mocks base method
func (m *MockMuxer) Wait() error {
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait
func (mr *MockMuxerMockRecorder) Wait() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockMuxer)(nil).Wait))
}

// Close mocks base method
func (m *MockMuxer) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockMuxerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMuxer)(nil).Close))
}
