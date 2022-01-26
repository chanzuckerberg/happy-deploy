// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chanzuckerberg/happy/pkg/stack_mgr (interfaces: StackIface)

// Package stack_mgr is a generated GoMock package.
package stack_mgr

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStackIface is a mock of StackIface interface.
type MockStackIface struct {
	ctrl     *gomock.Controller
	recorder *MockStackIfaceMockRecorder
}

// MockStackIfaceMockRecorder is the mock recorder for MockStackIface.
type MockStackIfaceMockRecorder struct {
	mock *MockStackIface
}

// NewMockStackIface creates a new mock instance.
func NewMockStackIface(ctrl *gomock.Controller) *MockStackIface {
	mock := &MockStackIface{ctrl: ctrl}
	mock.recorder = &MockStackIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStackIface) EXPECT() *MockStackIfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockStackIface) Apply() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply")
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockStackIfaceMockRecorder) Apply() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockStackIface)(nil).Apply))
}

// Destroy mocks base method.
func (m *MockStackIface) Destroy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockStackIfaceMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockStackIface)(nil).Destroy))
}

// GetMeta mocks base method.
func (m *MockStackIface) GetMeta() (*StackMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeta")
	ret0, _ := ret[0].(*StackMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeta indicates an expected call of GetMeta.
func (mr *MockStackIfaceMockRecorder) GetMeta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeta", reflect.TypeOf((*MockStackIface)(nil).GetMeta))
}

// GetName mocks base method.
func (m *MockStackIface) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockStackIfaceMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockStackIface)(nil).GetName))
}

// GetOutputs mocks base method.
func (m *MockStackIface) GetOutputs() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutputs")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutputs indicates an expected call of GetOutputs.
func (mr *MockStackIfaceMockRecorder) GetOutputs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutputs", reflect.TypeOf((*MockStackIface)(nil).GetOutputs))
}

// GetStatus mocks base method.
func (m *MockStackIface) GetStatus() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockStackIfaceMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockStackIface)(nil).GetStatus))
}

// PrintOutputs mocks base method.
func (m *MockStackIface) PrintOutputs() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrintOutputs")
	ret0, _ := ret[0].(string)
	return ret0
}

// PrintOutputs indicates an expected call of PrintOutputs.
func (mr *MockStackIfaceMockRecorder) PrintOutputs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintOutputs", reflect.TypeOf((*MockStackIface)(nil).PrintOutputs))
}

// SetMeta mocks base method.
func (m *MockStackIface) SetMeta(arg0 *StackMeta) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMeta", arg0)
}

// SetMeta indicates an expected call of SetMeta.
func (mr *MockStackIfaceMockRecorder) SetMeta(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMeta", reflect.TypeOf((*MockStackIface)(nil).SetMeta), arg0)
}
