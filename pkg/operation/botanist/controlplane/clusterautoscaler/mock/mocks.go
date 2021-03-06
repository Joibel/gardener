// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/operation/botanist/controlplane/clusterautoscaler (interfaces: ClusterAutoscaler)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	clusterautoscaler "github.com/gardener/gardener/pkg/operation/botanist/controlplane/clusterautoscaler"
	gomock "github.com/golang/mock/gomock"
	types "k8s.io/apimachinery/pkg/types"
)

// MockClusterAutoscaler is a mock of ClusterAutoscaler interface.
type MockClusterAutoscaler struct {
	ctrl     *gomock.Controller
	recorder *MockClusterAutoscalerMockRecorder
}

// MockClusterAutoscalerMockRecorder is the mock recorder for MockClusterAutoscaler.
type MockClusterAutoscalerMockRecorder struct {
	mock *MockClusterAutoscaler
}

// NewMockClusterAutoscaler creates a new mock instance.
func NewMockClusterAutoscaler(ctrl *gomock.Controller) *MockClusterAutoscaler {
	mock := &MockClusterAutoscaler{ctrl: ctrl}
	mock.recorder = &MockClusterAutoscalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterAutoscaler) EXPECT() *MockClusterAutoscalerMockRecorder {
	return m.recorder
}

// AlertingRules mocks base method.
func (m *MockClusterAutoscaler) AlertingRules() (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlertingRules")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlertingRules indicates an expected call of AlertingRules.
func (mr *MockClusterAutoscalerMockRecorder) AlertingRules() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlertingRules", reflect.TypeOf((*MockClusterAutoscaler)(nil).AlertingRules))
}

// Deploy mocks base method.
func (m *MockClusterAutoscaler) Deploy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockClusterAutoscalerMockRecorder) Deploy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockClusterAutoscaler)(nil).Deploy), arg0)
}

// Destroy mocks base method.
func (m *MockClusterAutoscaler) Destroy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockClusterAutoscalerMockRecorder) Destroy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockClusterAutoscaler)(nil).Destroy), arg0)
}

// ScrapeConfigs mocks base method.
func (m *MockClusterAutoscaler) ScrapeConfigs() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScrapeConfigs")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScrapeConfigs indicates an expected call of ScrapeConfigs.
func (mr *MockClusterAutoscalerMockRecorder) ScrapeConfigs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScrapeConfigs", reflect.TypeOf((*MockClusterAutoscaler)(nil).ScrapeConfigs))
}

// SetMachineDeployments mocks base method.
func (m *MockClusterAutoscaler) SetMachineDeployments(arg0 []v1alpha1.MachineDeployment) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMachineDeployments", arg0)
}

// SetMachineDeployments indicates an expected call of SetMachineDeployments.
func (mr *MockClusterAutoscalerMockRecorder) SetMachineDeployments(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMachineDeployments", reflect.TypeOf((*MockClusterAutoscaler)(nil).SetMachineDeployments), arg0)
}

// SetNamespaceUID mocks base method.
func (m *MockClusterAutoscaler) SetNamespaceUID(arg0 types.UID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNamespaceUID", arg0)
}

// SetNamespaceUID indicates an expected call of SetNamespaceUID.
func (mr *MockClusterAutoscalerMockRecorder) SetNamespaceUID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNamespaceUID", reflect.TypeOf((*MockClusterAutoscaler)(nil).SetNamespaceUID), arg0)
}

// SetSecrets mocks base method.
func (m *MockClusterAutoscaler) SetSecrets(arg0 clusterautoscaler.Secrets) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSecrets", arg0)
}

// SetSecrets indicates an expected call of SetSecrets.
func (mr *MockClusterAutoscalerMockRecorder) SetSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecrets", reflect.TypeOf((*MockClusterAutoscaler)(nil).SetSecrets), arg0)
}

// Wait mocks base method.
func (m *MockClusterAutoscaler) Wait(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockClusterAutoscalerMockRecorder) Wait(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockClusterAutoscaler)(nil).Wait), arg0)
}

// WaitCleanup mocks base method.
func (m *MockClusterAutoscaler) WaitCleanup(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitCleanup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitCleanup indicates an expected call of WaitCleanup.
func (mr *MockClusterAutoscalerMockRecorder) WaitCleanup(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitCleanup", reflect.TypeOf((*MockClusterAutoscaler)(nil).WaitCleanup), arg0)
}
