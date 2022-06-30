// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chanzuckerberg/happy/pkg/workspace_repo (interfaces: Workspace)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	options "github.com/chanzuckerberg/happy/pkg/options"
	util "github.com/chanzuckerberg/happy/pkg/util"
	gomock "github.com/golang/mock/gomock"
	tfe "github.com/hashicorp/go-tfe"
)

// MockWorkspace is a mock of Workspace interface.
type MockWorkspace struct {
	ctrl     *gomock.Controller
	recorder *MockWorkspaceMockRecorder
}

// MockWorkspaceMockRecorder is the mock recorder for MockWorkspace.
type MockWorkspaceMockRecorder struct {
	mock *MockWorkspace
}

// NewMockWorkspace creates a new mock instance.
func NewMockWorkspace(ctrl *gomock.Controller) *MockWorkspace {
	mock := &MockWorkspace{ctrl: ctrl}
	mock.recorder = &MockWorkspaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkspace) EXPECT() *MockWorkspaceMockRecorder {
	return m.recorder
}

// DiscardRun mocks base method.
func (m *MockWorkspace) DiscardRun(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DiscardRun", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DiscardRun indicates an expected call of DiscardRun.
func (mr *MockWorkspaceMockRecorder) DiscardRun(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DiscardRun", reflect.TypeOf((*MockWorkspace)(nil).DiscardRun), arg0, arg1)
}

// GetCurrentRunID mocks base method.
func (m *MockWorkspace) GetCurrentRunID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentRunID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCurrentRunID indicates an expected call of GetCurrentRunID.
func (mr *MockWorkspaceMockRecorder) GetCurrentRunID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentRunID", reflect.TypeOf((*MockWorkspace)(nil).GetCurrentRunID))
}

// GetCurrentRunStatus mocks base method.
func (m *MockWorkspace) GetCurrentRunStatus() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentRunStatus")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetCurrentRunStatus indicates an expected call of GetCurrentRunStatus.
func (mr *MockWorkspaceMockRecorder) GetCurrentRunStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentRunStatus", reflect.TypeOf((*MockWorkspace)(nil).GetCurrentRunStatus))
}

// GetLatestConfigVersionID mocks base method.
func (m *MockWorkspace) GetLatestConfigVersionID() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestConfigVersionID")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestConfigVersionID indicates an expected call of GetLatestConfigVersionID.
func (mr *MockWorkspaceMockRecorder) GetLatestConfigVersionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestConfigVersionID", reflect.TypeOf((*MockWorkspace)(nil).GetLatestConfigVersionID))
}

// GetOutputs mocks base method.
func (m *MockWorkspace) GetOutputs() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOutputs")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOutputs indicates an expected call of GetOutputs.
func (mr *MockWorkspaceMockRecorder) GetOutputs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOutputs", reflect.TypeOf((*MockWorkspace)(nil).GetOutputs))
}

// GetTags mocks base method.
func (m *MockWorkspace) GetTags() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTags")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTags indicates an expected call of GetTags.
func (mr *MockWorkspaceMockRecorder) GetTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTags", reflect.TypeOf((*MockWorkspace)(nil).GetTags))
}

// GetWorkspaceID mocks base method.
func (m *MockWorkspace) GetWorkspaceID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaceID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWorkspaceID indicates an expected call of GetWorkspaceID.
func (mr *MockWorkspaceMockRecorder) GetWorkspaceID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaceID", reflect.TypeOf((*MockWorkspace)(nil).GetWorkspaceID))
}

// GetWorkspaceId mocks base method.
func (m *MockWorkspace) GetWorkspaceId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWorkspaceId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetWorkspaceId indicates an expected call of GetWorkspaceId.
func (mr *MockWorkspaceMockRecorder) GetWorkspaceId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWorkspaceId", reflect.TypeOf((*MockWorkspace)(nil).GetWorkspaceId))
}

// HasState mocks base method.
func (m *MockWorkspace) HasState(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasState", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasState indicates an expected call of HasState.
func (mr *MockWorkspaceMockRecorder) HasState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasState", reflect.TypeOf((*MockWorkspace)(nil).HasState), arg0)
}

// ResetCache mocks base method.
func (m *MockWorkspace) ResetCache() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ResetCache")
}

// ResetCache indicates an expected call of ResetCache.
func (mr *MockWorkspaceMockRecorder) ResetCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetCache", reflect.TypeOf((*MockWorkspace)(nil).ResetCache))
}

// Run mocks base method.
func (m *MockWorkspace) Run(arg0 bool, arg1 util.DryRunType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockWorkspaceMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockWorkspace)(nil).Run), arg0, arg1)
}

// RunConfigVersion mocks base method.
func (m *MockWorkspace) RunConfigVersion(arg0 string, arg1 bool, arg2 util.DryRunType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunConfigVersion", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunConfigVersion indicates an expected call of RunConfigVersion.
func (mr *MockWorkspaceMockRecorder) RunConfigVersion(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunConfigVersion", reflect.TypeOf((*MockWorkspace)(nil).RunConfigVersion), arg0, arg1, arg2)
}

// SetClient mocks base method.
func (m *MockWorkspace) SetClient(arg0 *tfe.Client) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClient", arg0)
}

// SetClient indicates an expected call of SetClient.
func (mr *MockWorkspaceMockRecorder) SetClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClient", reflect.TypeOf((*MockWorkspace)(nil).SetClient), arg0)
}

// SetOutputs mocks base method.
func (m *MockWorkspace) SetOutputs(arg0 map[string]string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetOutputs", arg0)
}

// SetOutputs indicates an expected call of SetOutputs.
func (mr *MockWorkspaceMockRecorder) SetOutputs(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetOutputs", reflect.TypeOf((*MockWorkspace)(nil).SetOutputs), arg0)
}

// SetVars mocks base method.
func (m *MockWorkspace) SetVars(arg0, arg1, arg2 string, arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVars", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVars indicates an expected call of SetVars.
func (mr *MockWorkspaceMockRecorder) SetVars(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVars", reflect.TypeOf((*MockWorkspace)(nil).SetVars), arg0, arg1, arg2, arg3)
}

// SetWorkspace mocks base method.
func (m *MockWorkspace) SetWorkspace(arg0 *tfe.Workspace) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetWorkspace", arg0)
}

// SetWorkspace indicates an expected call of SetWorkspace.
func (mr *MockWorkspaceMockRecorder) SetWorkspace(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWorkspace", reflect.TypeOf((*MockWorkspace)(nil).SetWorkspace), arg0)
}

// UploadVersion mocks base method.
func (m *MockWorkspace) UploadVersion(arg0 string, arg1 util.DryRunType) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadVersion", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadVersion indicates an expected call of UploadVersion.
func (mr *MockWorkspaceMockRecorder) UploadVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadVersion", reflect.TypeOf((*MockWorkspace)(nil).UploadVersion), arg0, arg1)
}

// Wait mocks base method.
func (m *MockWorkspace) Wait(arg0 context.Context, arg1 util.DryRunType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockWorkspaceMockRecorder) Wait(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockWorkspace)(nil).Wait), arg0, arg1)
}

// WaitWithOptions mocks base method.
func (m *MockWorkspace) WaitWithOptions(arg0 context.Context, arg1 options.WaitOptions, arg2 util.DryRunType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitWithOptions", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitWithOptions indicates an expected call of WaitWithOptions.
func (mr *MockWorkspaceMockRecorder) WaitWithOptions(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitWithOptions", reflect.TypeOf((*MockWorkspace)(nil).WaitWithOptions), arg0, arg1, arg2)
}

// WorkspaceName mocks base method.
func (m *MockWorkspace) WorkspaceName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkspaceName")
	ret0, _ := ret[0].(string)
	return ret0
}

// WorkspaceName indicates an expected call of WorkspaceName.
func (mr *MockWorkspaceMockRecorder) WorkspaceName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkspaceName", reflect.TypeOf((*MockWorkspace)(nil).WorkspaceName))
}
