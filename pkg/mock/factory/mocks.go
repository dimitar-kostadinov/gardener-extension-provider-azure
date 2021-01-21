// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extension-provider-azure/pkg/azure/client (interfaces: Factory)

// Package factory is a generated GoMock package.
package factory

import (
	context "context"
	reflect "reflect"

	client "github.com/gardener/gardener-extension-provider-azure/pkg/azure/client"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// Group mocks base method.
func (m *MockFactory) Group(arg0 context.Context, arg1 v1.SecretReference) (client.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group", arg0, arg1)
	ret0, _ := ret[0].(client.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Group indicates an expected call of Group.
func (mr *MockFactoryMockRecorder) Group(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockFactory)(nil).Group), arg0, arg1)
}

// Storage mocks base method.
func (m *MockFactory) Storage(arg0 context.Context, arg1 v1.SecretReference) (client.Storage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Storage", arg0, arg1)
	ret0, _ := ret[0].(client.Storage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Storage indicates an expected call of Storage.
func (mr *MockFactoryMockRecorder) Storage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockFactory)(nil).Storage), arg0, arg1)
}

// StorageAccount mocks base method.
func (m *MockFactory) StorageAccount(arg0 context.Context, arg1 v1.SecretReference) (client.StorageAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageAccount", arg0, arg1)
	ret0, _ := ret[0].(client.StorageAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageAccount indicates an expected call of StorageAccount.
func (mr *MockFactoryMockRecorder) StorageAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageAccount", reflect.TypeOf((*MockFactory)(nil).StorageAccount), arg0, arg1)
}

// Vmss mocks base method.
func (m *MockFactory) Vmss(arg0 context.Context, arg1 v1.SecretReference) (client.Vmss, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vmss", arg0, arg1)
	ret0, _ := ret[0].(client.Vmss)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Vmss indicates an expected call of Vmss.
func (mr *MockFactoryMockRecorder) Vmss(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vmss", reflect.TypeOf((*MockFactory)(nil).Vmss), arg0, arg1)
}
