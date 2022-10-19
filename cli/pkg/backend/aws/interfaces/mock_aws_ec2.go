// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chanzuckerberg/happy/pkg/cli/backend/aws/interfaces (interfaces: EC2API)

// Package interfaces is a generated GoMock package.
package interfaces

import (
	context "context"
	reflect "reflect"

	ec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	gomock "github.com/golang/mock/gomock"
)

// MockEC2API is a mock of EC2API interface.
type MockEC2API struct {
	ctrl     *gomock.Controller
	recorder *MockEC2APIMockRecorder
}

// MockEC2APIMockRecorder is the mock recorder for MockEC2API.
type MockEC2APIMockRecorder struct {
	mock *MockEC2API
}

// NewMockEC2API creates a new mock instance.
func NewMockEC2API(ctrl *gomock.Controller) *MockEC2API {
	mock := &MockEC2API{ctrl: ctrl}
	mock.recorder = &MockEC2APIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEC2API) EXPECT() *MockEC2APIMockRecorder {
	return m.recorder
}

// DescribeInstances mocks base method.
func (m *MockEC2API) DescribeInstances(arg0 context.Context, arg1 *ec2.DescribeInstancesInput, arg2 ...func(*ec2.Options)) (*ec2.DescribeInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeInstances", varargs...)
	ret0, _ := ret[0].(*ec2.DescribeInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstances indicates an expected call of DescribeInstances.
func (mr *MockEC2APIMockRecorder) DescribeInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstances", reflect.TypeOf((*MockEC2API)(nil).DescribeInstances), varargs...)
}
