// Code generated by MockGen. DO NOT EDIT.
// Source: routes/unmarshaller.go

// Package mocks is a generated GoMock package.
package mocks

import (
	routes "github.com/gessnerfl/awsroutes/routes"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUnmarshaller is a mock of Unmarshaller interface
type MockUnmarshaller struct {
	ctrl     *gomock.Controller
	recorder *MockUnmarshallerMockRecorder
}

// MockUnmarshallerMockRecorder is the mock recorder for MockUnmarshaller
type MockUnmarshallerMockRecorder struct {
	mock *MockUnmarshaller
}

// NewMockUnmarshaller creates a new mock instance
func NewMockUnmarshaller(ctrl *gomock.Controller) *MockUnmarshaller {
	mock := &MockUnmarshaller{ctrl: ctrl}
	mock.recorder = &MockUnmarshallerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUnmarshaller) EXPECT() *MockUnmarshallerMockRecorder {
	return m.recorder
}

// Unmarshal mocks base method
func (m *MockUnmarshaller) Unmarshal(data []byte) (*routes.IPAddressRanges, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unmarshal", data)
	ret0, _ := ret[0].(*routes.IPAddressRanges)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unmarshal indicates an expected call of Unmarshal
func (mr *MockUnmarshallerMockRecorder) Unmarshal(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unmarshal", reflect.TypeOf((*MockUnmarshaller)(nil).Unmarshal), data)
}
