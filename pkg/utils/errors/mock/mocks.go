// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/utils/errors (interfaces: TaskFunc)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	errors "github.com/gardener/gardener/pkg/utils/errors"
	gomock "github.com/golang/mock/gomock"
)

// MockTaskFunc is a mock of TaskFunc interface.
type MockTaskFunc struct {
	ctrl     *gomock.Controller
	recorder *MockTaskFuncMockRecorder
}

// MockTaskFuncMockRecorder is the mock recorder for MockTaskFunc.
type MockTaskFuncMockRecorder struct {
	mock *MockTaskFunc
}

// NewMockTaskFunc creates a new mock instance.
func NewMockTaskFunc(ctrl *gomock.Controller) *MockTaskFunc {
	mock := &MockTaskFunc{ctrl: ctrl}
	mock.recorder = &MockTaskFuncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskFunc) EXPECT() *MockTaskFuncMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockTaskFunc) Do(arg0 *errors.ErrorContext) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockTaskFuncMockRecorder) Do(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockTaskFunc)(nil).Do), arg0)
}