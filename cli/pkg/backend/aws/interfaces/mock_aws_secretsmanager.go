// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chanzuckerberg/happy/cli/pkg/backend/aws/interfaces (interfaces: SecretsManagerAPI)

// Package interfaces is a generated GoMock package.
package interfaces

import (
	context "context"
	reflect "reflect"

	secretsmanager "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	gomock "github.com/golang/mock/gomock"
)

// MockSecretsManagerAPI is a mock of SecretsManagerAPI interface.
type MockSecretsManagerAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsManagerAPIMockRecorder
}

// MockSecretsManagerAPIMockRecorder is the mock recorder for MockSecretsManagerAPI.
type MockSecretsManagerAPIMockRecorder struct {
	mock *MockSecretsManagerAPI
}

// NewMockSecretsManagerAPI creates a new mock instance.
func NewMockSecretsManagerAPI(ctrl *gomock.Controller) *MockSecretsManagerAPI {
	mock := &MockSecretsManagerAPI{ctrl: ctrl}
	mock.recorder = &MockSecretsManagerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsManagerAPI) EXPECT() *MockSecretsManagerAPIMockRecorder {
	return m.recorder
}

// GetSecretValue mocks base method.
func (m *MockSecretsManagerAPI) GetSecretValue(arg0 context.Context, arg1 *secretsmanager.GetSecretValueInput, arg2 ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSecretValue", varargs...)
	ret0, _ := ret[0].(*secretsmanager.GetSecretValueOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretValue indicates an expected call of GetSecretValue.
func (mr *MockSecretsManagerAPIMockRecorder) GetSecretValue(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretValue", reflect.TypeOf((*MockSecretsManagerAPI)(nil).GetSecretValue), varargs...)
}
