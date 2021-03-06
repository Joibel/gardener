// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/operation/botanist/controlplane/kubecontrollermanager (interfaces: KubeControllerManager)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	kubecontrollermanager "github.com/gardener/gardener/pkg/operation/botanist/controlplane/kubecontrollermanager"
	gomock "github.com/golang/mock/gomock"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockKubeControllerManager is a mock of KubeControllerManager interface.
type MockKubeControllerManager struct {
	ctrl     *gomock.Controller
	recorder *MockKubeControllerManagerMockRecorder
}

// MockKubeControllerManagerMockRecorder is the mock recorder for MockKubeControllerManager.
type MockKubeControllerManagerMockRecorder struct {
	mock *MockKubeControllerManager
}

// NewMockKubeControllerManager creates a new mock instance.
func NewMockKubeControllerManager(ctrl *gomock.Controller) *MockKubeControllerManager {
	mock := &MockKubeControllerManager{ctrl: ctrl}
	mock.recorder = &MockKubeControllerManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubeControllerManager) EXPECT() *MockKubeControllerManagerMockRecorder {
	return m.recorder
}

// AlertingRules mocks base method.
func (m *MockKubeControllerManager) AlertingRules() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertingRules")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertingRules indicates an expected call of AlertingRules.
func (mr *MockKubeControllerManagerMockRecorder) AlertingRules() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertingRules", reflect.TypeOf((*MockKubeControllerManager)(nil).AlertingRules))
}

// Deploy mocks base method.
func (m *MockKubeControllerManager) Deploy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockKubeControllerManagerMockRecorder) Deploy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockKubeControllerManager)(nil).Deploy), arg0)
}

// Destroy mocks base method.
func (m *MockKubeControllerManager) Destroy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockKubeControllerManagerMockRecorder) Destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockKubeControllerManager)(nil).Destroy), arg0)
}

// ScrapeConfigs mocks base method.
func (m *MockKubeControllerManager) ScrapeConfigs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScrapeConfigs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScrapeConfigs indicates an expected call of ScrapeConfigs.
func (mr *MockKubeControllerManagerMockRecorder) ScrapeConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScrapeConfigs", reflect.TypeOf((*MockKubeControllerManager)(nil).ScrapeConfigs))
}

// SetReplicaCount mocks base method.
func (m *MockKubeControllerManager) SetReplicaCount(arg0 int32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReplicaCount", arg0)
}

// SetReplicaCount indicates an expected call of SetReplicaCount.
func (mr *MockKubeControllerManagerMockRecorder) SetReplicaCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReplicaCount", reflect.TypeOf((*MockKubeControllerManager)(nil).SetReplicaCount), arg0)
}

// SetSecrets mocks base method.
func (m *MockKubeControllerManager) SetSecrets(arg0 kubecontrollermanager.Secrets) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSecrets", arg0)
}

// SetSecrets indicates an expected call of SetSecrets.
func (mr *MockKubeControllerManagerMockRecorder) SetSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecrets", reflect.TypeOf((*MockKubeControllerManager)(nil).SetSecrets), arg0)
}

// SetShootClient mocks base method.
func (m *MockKubeControllerManager) SetShootClient(arg0 client.Client) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetShootClient", arg0)
}

// SetShootClient indicates an expected call of SetShootClient.
func (mr *MockKubeControllerManagerMockRecorder) SetShootClient(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetShootClient", reflect.TypeOf((*MockKubeControllerManager)(nil).SetShootClient), arg0)
}

// Wait mocks base method.
func (m *MockKubeControllerManager) Wait(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockKubeControllerManagerMockRecorder) Wait(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockKubeControllerManager)(nil).Wait), arg0)
}

// WaitCleanup mocks base method.
func (m *MockKubeControllerManager) WaitCleanup(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitCleanup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitCleanup indicates an expected call of WaitCleanup.
func (mr *MockKubeControllerManagerMockRecorder) WaitCleanup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitCleanup", reflect.TypeOf((*MockKubeControllerManager)(nil).WaitCleanup), arg0)
}

// WaitForControllerToBeActive mocks base method.
func (m *MockKubeControllerManager) WaitForControllerToBeActive(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForControllerToBeActive", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitForControllerToBeActive indicates an expected call of WaitForControllerToBeActive.
func (mr *MockKubeControllerManagerMockRecorder) WaitForControllerToBeActive(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForControllerToBeActive", reflect.TypeOf((*MockKubeControllerManager)(nil).WaitForControllerToBeActive), arg0)
}
