// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fidelity/kconnect/pkg/history/loader (interfaces: Loader)

// Package mock_loader is a generated GoMock package.
package mock_loader

import (
	v1alpha1 "github.com/fidelity/kconnect/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockLoader is a mock of Loader interface
type MockLoader struct {
	ctrl     *gomock.Controller
	recorder *MockLoaderMockRecorder
}

// MockLoaderMockRecorder is the mock recorder for MockLoader
type MockLoaderMockRecorder struct {
	mock *MockLoader
}

// NewMockLoader creates a new mock instance
func NewMockLoader(ctrl *gomock.Controller) *MockLoader {
	mock := &MockLoader{ctrl: ctrl}
	mock.recorder = &MockLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLoader) EXPECT() *MockLoaderMockRecorder {
	return m.recorder
}

// Load mocks base method
func (m *MockLoader) Load() (*v1alpha1.HistoryEntryList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].(*v1alpha1.HistoryEntryList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load
func (mr *MockLoaderMockRecorder) Load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockLoader)(nil).Load))
}

// Save mocks base method
func (m *MockLoader) Save(arg0 *v1alpha1.HistoryEntryList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockLoaderMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockLoader)(nil).Save), arg0)
}