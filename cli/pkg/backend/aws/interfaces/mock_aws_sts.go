// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chanzuckerberg/happy/pkg/cli/backend/aws/interfaces (interfaces: STSAPI)

// Package interfaces is a generated GoMock package.
package interfaces

import (
	context "context"
	reflect "reflect"

	sts "github.com/aws/aws-sdk-go-v2/service/sts"
	gomock "github.com/golang/mock/gomock"
)

// MockSTSAPI is a mock of STSAPI interface.
type MockSTSAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSTSAPIMockRecorder
}

// MockSTSAPIMockRecorder is the mock recorder for MockSTSAPI.
type MockSTSAPIMockRecorder struct {
	mock *MockSTSAPI
}

// NewMockSTSAPI creates a new mock instance.
func NewMockSTSAPI(ctrl *gomock.Controller) *MockSTSAPI {
	mock := &MockSTSAPI{ctrl: ctrl}
	mock.recorder = &MockSTSAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSTSAPI) EXPECT() *MockSTSAPIMockRecorder {
	return m.recorder
}

// AssumeRole mocks base method.
func (m *MockSTSAPI) AssumeRole(arg0 context.Context, arg1 *sts.AssumeRoleInput, arg2 ...func(*sts.Options)) (*sts.AssumeRoleOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssumeRole", varargs...)
	ret0, _ := ret[0].(*sts.AssumeRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRole indicates an expected call of AssumeRole.
func (mr *MockSTSAPIMockRecorder) AssumeRole(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRole", reflect.TypeOf((*MockSTSAPI)(nil).AssumeRole), varargs...)
}

// AssumeRoleWithSAML mocks base method.
func (m *MockSTSAPI) AssumeRoleWithSAML(arg0 context.Context, arg1 *sts.AssumeRoleWithSAMLInput, arg2 ...func(*sts.Options)) (*sts.AssumeRoleWithSAMLOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssumeRoleWithSAML", varargs...)
	ret0, _ := ret[0].(*sts.AssumeRoleWithSAMLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleWithSAML indicates an expected call of AssumeRoleWithSAML.
func (mr *MockSTSAPIMockRecorder) AssumeRoleWithSAML(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithSAML", reflect.TypeOf((*MockSTSAPI)(nil).AssumeRoleWithSAML), varargs...)
}

// AssumeRoleWithWebIdentity mocks base method.
func (m *MockSTSAPI) AssumeRoleWithWebIdentity(arg0 context.Context, arg1 *sts.AssumeRoleWithWebIdentityInput, arg2 ...func(*sts.Options)) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AssumeRoleWithWebIdentity", varargs...)
	ret0, _ := ret[0].(*sts.AssumeRoleWithWebIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleWithWebIdentity indicates an expected call of AssumeRoleWithWebIdentity.
func (mr *MockSTSAPIMockRecorder) AssumeRoleWithWebIdentity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithWebIdentity", reflect.TypeOf((*MockSTSAPI)(nil).AssumeRoleWithWebIdentity), varargs...)
}

// DecodeAuthorizationMessage mocks base method.
func (m *MockSTSAPI) DecodeAuthorizationMessage(arg0 context.Context, arg1 *sts.DecodeAuthorizationMessageInput, arg2 ...func(*sts.Options)) (*sts.DecodeAuthorizationMessageOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DecodeAuthorizationMessage", varargs...)
	ret0, _ := ret[0].(*sts.DecodeAuthorizationMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeAuthorizationMessage indicates an expected call of DecodeAuthorizationMessage.
func (mr *MockSTSAPIMockRecorder) DecodeAuthorizationMessage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeAuthorizationMessage", reflect.TypeOf((*MockSTSAPI)(nil).DecodeAuthorizationMessage), varargs...)
}

// GetAccessKeyInfo mocks base method.
func (m *MockSTSAPI) GetAccessKeyInfo(arg0 context.Context, arg1 *sts.GetAccessKeyInfoInput, arg2 ...func(*sts.Options)) (*sts.GetAccessKeyInfoOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAccessKeyInfo", varargs...)
	ret0, _ := ret[0].(*sts.GetAccessKeyInfoOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessKeyInfo indicates an expected call of GetAccessKeyInfo.
func (mr *MockSTSAPIMockRecorder) GetAccessKeyInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessKeyInfo", reflect.TypeOf((*MockSTSAPI)(nil).GetAccessKeyInfo), varargs...)
}

// GetCallerIdentity mocks base method.
func (m *MockSTSAPI) GetCallerIdentity(arg0 context.Context, arg1 *sts.GetCallerIdentityInput, arg2 ...func(*sts.Options)) (*sts.GetCallerIdentityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCallerIdentity", varargs...)
	ret0, _ := ret[0].(*sts.GetCallerIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCallerIdentity indicates an expected call of GetCallerIdentity.
func (mr *MockSTSAPIMockRecorder) GetCallerIdentity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallerIdentity", reflect.TypeOf((*MockSTSAPI)(nil).GetCallerIdentity), varargs...)
}

// GetFederationToken mocks base method.
func (m *MockSTSAPI) GetFederationToken(arg0 context.Context, arg1 *sts.GetFederationTokenInput, arg2 ...func(*sts.Options)) (*sts.GetFederationTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFederationToken", varargs...)
	ret0, _ := ret[0].(*sts.GetFederationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFederationToken indicates an expected call of GetFederationToken.
func (mr *MockSTSAPIMockRecorder) GetFederationToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFederationToken", reflect.TypeOf((*MockSTSAPI)(nil).GetFederationToken), varargs...)
}

// GetSessionToken mocks base method.
func (m *MockSTSAPI) GetSessionToken(arg0 context.Context, arg1 *sts.GetSessionTokenInput, arg2 ...func(*sts.Options)) (*sts.GetSessionTokenOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSessionToken", varargs...)
	ret0, _ := ret[0].(*sts.GetSessionTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionToken indicates an expected call of GetSessionToken.
func (mr *MockSTSAPIMockRecorder) GetSessionToken(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionToken", reflect.TypeOf((*MockSTSAPI)(nil).GetSessionToken), varargs...)
}
