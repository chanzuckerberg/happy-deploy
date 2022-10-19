// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/chanzuckerberg/happy/cli/pkg/backend/aws/interfaces (interfaces: DynamoDB)

// Package interfaces is a generated GoMock package.
package interfaces

import (
	context "context"
	reflect "reflect"

	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	gomock "github.com/golang/mock/gomock"
)

// MockDynamoDB is a mock of DynamoDB interface.
type MockDynamoDB struct {
	ctrl     *gomock.Controller
	recorder *MockDynamoDBMockRecorder
}

// MockDynamoDBMockRecorder is the mock recorder for MockDynamoDB.
type MockDynamoDBMockRecorder struct {
	mock *MockDynamoDB
}

// NewMockDynamoDB creates a new mock instance.
func NewMockDynamoDB(ctrl *gomock.Controller) *MockDynamoDB {
	mock := &MockDynamoDB{ctrl: ctrl}
	mock.recorder = &MockDynamoDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDynamoDB) EXPECT() *MockDynamoDBMockRecorder {
	return m.recorder
}

// CreateTable mocks base method.
func (m *MockDynamoDB) CreateTable(arg0 context.Context, arg1 *dynamodb.CreateTableInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTable", varargs...)
	ret0, _ := ret[0].(*dynamodb.CreateTableOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTable indicates an expected call of CreateTable.
func (mr *MockDynamoDBMockRecorder) CreateTable(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTable", reflect.TypeOf((*MockDynamoDB)(nil).CreateTable), varargs...)
}

// DeleteItem mocks base method.
func (m *MockDynamoDB) DeleteItem(arg0 context.Context, arg1 *dynamodb.DeleteItemInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteItem", varargs...)
	ret0, _ := ret[0].(*dynamodb.DeleteItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteItem indicates an expected call of DeleteItem.
func (mr *MockDynamoDBMockRecorder) DeleteItem(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItem", reflect.TypeOf((*MockDynamoDB)(nil).DeleteItem), varargs...)
}

// GetItem mocks base method.
func (m *MockDynamoDB) GetItem(arg0 context.Context, arg1 *dynamodb.GetItemInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetItem", varargs...)
	ret0, _ := ret[0].(*dynamodb.GetItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetItem indicates an expected call of GetItem.
func (mr *MockDynamoDBMockRecorder) GetItem(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetItem", reflect.TypeOf((*MockDynamoDB)(nil).GetItem), varargs...)
}

// PutItem mocks base method.
func (m *MockDynamoDB) PutItem(arg0 context.Context, arg1 *dynamodb.PutItemInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutItem", varargs...)
	ret0, _ := ret[0].(*dynamodb.PutItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutItem indicates an expected call of PutItem.
func (mr *MockDynamoDBMockRecorder) PutItem(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutItem", reflect.TypeOf((*MockDynamoDB)(nil).PutItem), varargs...)
}

// UpdateItem mocks base method.
func (m *MockDynamoDB) UpdateItem(arg0 context.Context, arg1 *dynamodb.UpdateItemInput, arg2 ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateItem", varargs...)
	ret0, _ := ret[0].(*dynamodb.UpdateItemOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateItem indicates an expected call of UpdateItem.
func (mr *MockDynamoDBMockRecorder) UpdateItem(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateItem", reflect.TypeOf((*MockDynamoDB)(nil).UpdateItem), varargs...)
}
