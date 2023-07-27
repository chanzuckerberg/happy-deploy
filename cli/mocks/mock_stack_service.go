// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chanzuckerberg/happy/cli/pkg/stack_mgr (interfaces: StackServiceIface)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	stack_mgr "github.com/chanzuckerberg/happy/shared/stack"
	config "github.com/chanzuckerberg/happy/shared/config"
	workspace_repo "github.com/chanzuckerberg/happy/shared/workspace_repo"
	gomock "github.com/golang/mock/gomock"
)

// MockStackServiceIface is a mock of StackServiceIface interface.
type MockStackServiceIface struct {
	ctrl     *gomock.Controller
	recorder *MockStackServiceIfaceMockRecorder
}

// MockStackServiceIfaceMockRecorder is the mock recorder for MockStackServiceIface.
type MockStackServiceIfaceMockRecorder struct {
	mock *MockStackServiceIface
}

// NewMockStackServiceIface creates a new mock instance.
func NewMockStackServiceIface(ctrl *gomock.Controller) *MockStackServiceIface {
	mock := &MockStackServiceIface{ctrl: ctrl}
	mock.recorder = &MockStackServiceIfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStackServiceIface) EXPECT() *MockStackServiceIfaceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockStackServiceIface) Add(arg0 context.Context, arg1 string, arg2 ...workspace_repo.TFERunOption) (*stack_mgr.Stack, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Add", varargs...)
	ret0, _ := ret[0].(*stack_mgr.Stack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockStackServiceIfaceMockRecorder) Add(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStackServiceIface)(nil).Add), varargs...)
}

// GetConfig mocks base method.
func (m *MockStackServiceIface) GetConfig() *config.HappyConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*config.HappyConfig)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockStackServiceIfaceMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockStackServiceIface)(nil).GetConfig))
}

// GetStackWorkspace mocks base method.
func (m *MockStackServiceIface) GetStackWorkspace(arg0 context.Context, arg1 string) (workspace_repo.Workspace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStackWorkspace", arg0, arg1)
	ret0, _ := ret[0].(workspace_repo.Workspace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStackWorkspace indicates an expected call of GetStackWorkspace.
func (mr *MockStackServiceIfaceMockRecorder) GetStackWorkspace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStackWorkspace", reflect.TypeOf((*MockStackServiceIface)(nil).GetStackWorkspace), arg0, arg1)
}

// GetStacks mocks base method.
func (m *MockStackServiceIface) GetStacks(arg0 context.Context) (map[string]*stack_mgr.Stack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStacks", arg0)
	ret0, _ := ret[0].(map[string]*stack_mgr.Stack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStacks indicates an expected call of GetStacks.
func (mr *MockStackServiceIfaceMockRecorder) GetStacks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStacks", reflect.TypeOf((*MockStackServiceIface)(nil).GetStacks), arg0)
}

// NewStackMeta mocks base method.
func (m *MockStackServiceIface) NewStackMeta(arg0 string) *stack_mgr.StackMeta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewStackMeta", arg0)
	ret0, _ := ret[0].(*stack_mgr.StackMeta)
	return ret0
}

// NewStackMeta indicates an expected call of NewStackMeta.
func (mr *MockStackServiceIfaceMockRecorder) NewStackMeta(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewStackMeta", reflect.TypeOf((*MockStackServiceIface)(nil).NewStackMeta), arg0)
}

// Remove mocks base method.
func (m *MockStackServiceIface) Remove(arg0 context.Context, arg1 string, arg2 ...workspace_repo.TFERunOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Remove", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockStackServiceIfaceMockRecorder) Remove(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockStackServiceIface)(nil).Remove), varargs...)
}
