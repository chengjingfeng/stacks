// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/docker/stacks/pkg/interfaces (interfaces: BackendClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	time "time"

	types "github.com/docker/docker/api/types"
	events "github.com/docker/docker/api/types/events"
	filters "github.com/docker/docker/api/types/filters"
	swarm "github.com/docker/docker/api/types/swarm"
	types0 "github.com/docker/stacks/pkg/types"
	gomock "github.com/golang/mock/gomock"
)

// MockBackendClient is a mock of BackendClient interface
type MockBackendClient struct {
	ctrl     *gomock.Controller
	recorder *MockBackendClientMockRecorder
}

// MockBackendClientMockRecorder is the mock recorder for MockBackendClient
type MockBackendClientMockRecorder struct {
	mock *MockBackendClient
}

// NewMockBackendClient creates a new mock instance
func NewMockBackendClient(ctrl *gomock.Controller) *MockBackendClient {
	mock := &MockBackendClient{ctrl: ctrl}
	mock.recorder = &MockBackendClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBackendClient) EXPECT() *MockBackendClientMockRecorder {
	return m.recorder
}

// CreateConfig mocks base method
func (m *MockBackendClient) CreateConfig(arg0 swarm.ConfigSpec) (string, error) {
	ret := m.ctrl.Call(m, "CreateConfig", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateConfig indicates an expected call of CreateConfig
func (mr *MockBackendClientMockRecorder) CreateConfig(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateConfig", reflect.TypeOf((*MockBackendClient)(nil).CreateConfig), arg0)
}

// CreateNetwork mocks base method
func (m *MockBackendClient) CreateNetwork(arg0 types.NetworkCreateRequest) (string, error) {
	ret := m.ctrl.Call(m, "CreateNetwork", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNetwork indicates an expected call of CreateNetwork
func (mr *MockBackendClientMockRecorder) CreateNetwork(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetwork", reflect.TypeOf((*MockBackendClient)(nil).CreateNetwork), arg0)
}

// CreateSecret mocks base method
func (m *MockBackendClient) CreateSecret(arg0 swarm.SecretSpec) (string, error) {
	ret := m.ctrl.Call(m, "CreateSecret", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecret indicates an expected call of CreateSecret
func (mr *MockBackendClientMockRecorder) CreateSecret(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockBackendClient)(nil).CreateSecret), arg0)
}

// CreateService mocks base method
func (m *MockBackendClient) CreateService(arg0 swarm.ServiceSpec, arg1 string, arg2 bool) (*types.ServiceCreateResponse, error) {
	ret := m.ctrl.Call(m, "CreateService", arg0, arg1, arg2)
	ret0, _ := ret[0].(*types.ServiceCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateService indicates an expected call of CreateService
func (mr *MockBackendClientMockRecorder) CreateService(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateService", reflect.TypeOf((*MockBackendClient)(nil).CreateService), arg0, arg1, arg2)
}

// CreateStack mocks base method
func (m *MockBackendClient) CreateStack(arg0 types0.StackSpec) (types0.StackCreateResponse, error) {
	ret := m.ctrl.Call(m, "CreateStack", arg0)
	ret0, _ := ret[0].(types0.StackCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStack indicates an expected call of CreateStack
func (mr *MockBackendClientMockRecorder) CreateStack(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStack", reflect.TypeOf((*MockBackendClient)(nil).CreateStack), arg0)
}

// DeleteStack mocks base method
func (m *MockBackendClient) DeleteStack(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteStack", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStack indicates an expected call of DeleteStack
func (mr *MockBackendClientMockRecorder) DeleteStack(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStack", reflect.TypeOf((*MockBackendClient)(nil).DeleteStack), arg0)
}

// GetConfig mocks base method
func (m *MockBackendClient) GetConfig(arg0 string) (swarm.Config, error) {
	ret := m.ctrl.Call(m, "GetConfig", arg0)
	ret0, _ := ret[0].(swarm.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig
func (mr *MockBackendClientMockRecorder) GetConfig(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockBackendClient)(nil).GetConfig), arg0)
}

// GetConfigs mocks base method
func (m *MockBackendClient) GetConfigs(arg0 types.ConfigListOptions) ([]swarm.Config, error) {
	ret := m.ctrl.Call(m, "GetConfigs", arg0)
	ret0, _ := ret[0].([]swarm.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigs indicates an expected call of GetConfigs
func (mr *MockBackendClientMockRecorder) GetConfigs(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigs", reflect.TypeOf((*MockBackendClient)(nil).GetConfigs), arg0)
}

// GetNetwork mocks base method
func (m *MockBackendClient) GetNetwork(arg0 string) (types.NetworkResource, error) {
	ret := m.ctrl.Call(m, "GetNetwork", arg0)
	ret0, _ := ret[0].(types.NetworkResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetwork indicates an expected call of GetNetwork
func (mr *MockBackendClientMockRecorder) GetNetwork(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetwork", reflect.TypeOf((*MockBackendClient)(nil).GetNetwork), arg0)
}

// GetNetworks mocks base method
func (m *MockBackendClient) GetNetworks(arg0 filters.Args) ([]types.NetworkResource, error) {
	ret := m.ctrl.Call(m, "GetNetworks", arg0)
	ret0, _ := ret[0].([]types.NetworkResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworks indicates an expected call of GetNetworks
func (mr *MockBackendClientMockRecorder) GetNetworks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworks", reflect.TypeOf((*MockBackendClient)(nil).GetNetworks), arg0)
}

// GetNetworksByName mocks base method
func (m *MockBackendClient) GetNetworksByName(arg0 string) ([]types.NetworkResource, error) {
	ret := m.ctrl.Call(m, "GetNetworksByName", arg0)
	ret0, _ := ret[0].([]types.NetworkResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworksByName indicates an expected call of GetNetworksByName
func (mr *MockBackendClientMockRecorder) GetNetworksByName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworksByName", reflect.TypeOf((*MockBackendClient)(nil).GetNetworksByName), arg0)
}

// GetNode mocks base method
func (m *MockBackendClient) GetNode(arg0 string) (swarm.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNode", arg0)
	ret0, _ := ret[0].(swarm.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNode indicates an expected call of GetNode
func (mr *MockBackendClientMockRecorder) GetNode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNode", reflect.TypeOf((*MockBackendClient)(nil).GetNode), arg0)
}

// GetSecret mocks base method
func (m *MockBackendClient) GetSecret(arg0 string) (swarm.Secret, error) {
	ret := m.ctrl.Call(m, "GetSecret", arg0)
	ret0, _ := ret[0].(swarm.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret
func (mr *MockBackendClientMockRecorder) GetSecret(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockBackendClient)(nil).GetSecret), arg0)
}

// GetSecrets mocks base method
func (m *MockBackendClient) GetSecrets(arg0 types.SecretListOptions) ([]swarm.Secret, error) {
	ret := m.ctrl.Call(m, "GetSecrets", arg0)
	ret0, _ := ret[0].([]swarm.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecrets indicates an expected call of GetSecrets
func (mr *MockBackendClientMockRecorder) GetSecrets(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecrets", reflect.TypeOf((*MockBackendClient)(nil).GetSecrets), arg0)
}

// GetService mocks base method
func (m *MockBackendClient) GetService(arg0 string, arg1 bool) (swarm.Service, error) {
	ret := m.ctrl.Call(m, "GetService", arg0, arg1)
	ret0, _ := ret[0].(swarm.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService
func (mr *MockBackendClientMockRecorder) GetService(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockBackendClient)(nil).GetService), arg0, arg1)
}

// GetServices mocks base method
func (m *MockBackendClient) GetServices(arg0 types.ServiceListOptions) ([]swarm.Service, error) {
	ret := m.ctrl.Call(m, "GetServices", arg0)
	ret0, _ := ret[0].([]swarm.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServices indicates an expected call of GetServices
func (mr *MockBackendClientMockRecorder) GetServices(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServices", reflect.TypeOf((*MockBackendClient)(nil).GetServices), arg0)
}

// GetStack mocks base method
func (m *MockBackendClient) GetStack(arg0 string) (types0.Stack, error) {
	ret := m.ctrl.Call(m, "GetStack", arg0)
	ret0, _ := ret[0].(types0.Stack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStack indicates an expected call of GetStack
func (mr *MockBackendClientMockRecorder) GetStack(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStack", reflect.TypeOf((*MockBackendClient)(nil).GetStack), arg0)
}

// GetTask mocks base method
func (m *MockBackendClient) GetTask(arg0 string) (swarm.Task, error) {
	ret := m.ctrl.Call(m, "GetTask", arg0)
	ret0, _ := ret[0].(swarm.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask
func (mr *MockBackendClientMockRecorder) GetTask(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockBackendClient)(nil).GetTask), arg0)
}

// GetTasks mocks base method
func (m *MockBackendClient) GetTasks(arg0 types.TaskListOptions) ([]swarm.Task, error) {
	ret := m.ctrl.Call(m, "GetTasks", arg0)
	ret0, _ := ret[0].([]swarm.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasks indicates an expected call of GetTasks
func (mr *MockBackendClientMockRecorder) GetTasks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasks", reflect.TypeOf((*MockBackendClient)(nil).GetTasks), arg0)
}

// Info mocks base method
func (m *MockBackendClient) Info() swarm.Info {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info")
	ret0, _ := ret[0].(swarm.Info)
	return ret0
}

// Info indicates an expected call of Info
func (mr *MockBackendClientMockRecorder) Info() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockBackendClient)(nil).Info))
}

// ListStacks mocks base method
func (m *MockBackendClient) ListStacks() ([]types0.Stack, error) {
	ret := m.ctrl.Call(m, "ListStacks")
	ret0, _ := ret[0].([]types0.Stack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStacks indicates an expected call of ListStacks
func (mr *MockBackendClientMockRecorder) ListStacks() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStacks", reflect.TypeOf((*MockBackendClient)(nil).ListStacks))
}

// RemoveConfig mocks base method
func (m *MockBackendClient) RemoveConfig(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveConfig", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveConfig indicates an expected call of RemoveConfig
func (mr *MockBackendClientMockRecorder) RemoveConfig(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveConfig", reflect.TypeOf((*MockBackendClient)(nil).RemoveConfig), arg0)
}

// RemoveNetwork mocks base method
func (m *MockBackendClient) RemoveNetwork(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveNetwork", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNetwork indicates an expected call of RemoveNetwork
func (mr *MockBackendClientMockRecorder) RemoveNetwork(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNetwork", reflect.TypeOf((*MockBackendClient)(nil).RemoveNetwork), arg0)
}

// RemoveSecret mocks base method
func (m *MockBackendClient) RemoveSecret(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveSecret", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecret indicates an expected call of RemoveSecret
func (mr *MockBackendClientMockRecorder) RemoveSecret(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSecret", reflect.TypeOf((*MockBackendClient)(nil).RemoveSecret), arg0)
}

// RemoveService mocks base method
func (m *MockBackendClient) RemoveService(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveService", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveService indicates an expected call of RemoveService
func (mr *MockBackendClientMockRecorder) RemoveService(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveService", reflect.TypeOf((*MockBackendClient)(nil).RemoveService), arg0)
}

// SubscribeToEvents mocks base method
func (m *MockBackendClient) SubscribeToEvents(arg0, arg1 time.Time, arg2 filters.Args) ([]events.Message, chan interface{}) {
	ret := m.ctrl.Call(m, "SubscribeToEvents", arg0, arg1, arg2)
	ret0, _ := ret[0].([]events.Message)
	ret1, _ := ret[1].(chan interface{})
	return ret0, ret1
}

// SubscribeToEvents indicates an expected call of SubscribeToEvents
func (mr *MockBackendClientMockRecorder) SubscribeToEvents(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToEvents", reflect.TypeOf((*MockBackendClient)(nil).SubscribeToEvents), arg0, arg1, arg2)
}

// UnsubscribeFromEvents mocks base method
func (m *MockBackendClient) UnsubscribeFromEvents(arg0 chan interface{}) {
	m.ctrl.Call(m, "UnsubscribeFromEvents", arg0)
}

// UnsubscribeFromEvents indicates an expected call of UnsubscribeFromEvents
func (mr *MockBackendClientMockRecorder) UnsubscribeFromEvents(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnsubscribeFromEvents", reflect.TypeOf((*MockBackendClient)(nil).UnsubscribeFromEvents), arg0)
}

// UpdateConfig mocks base method
func (m *MockBackendClient) UpdateConfig(arg0 string, arg1 uint64, arg2 swarm.ConfigSpec) error {
	ret := m.ctrl.Call(m, "UpdateConfig", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConfig indicates an expected call of UpdateConfig
func (mr *MockBackendClientMockRecorder) UpdateConfig(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConfig", reflect.TypeOf((*MockBackendClient)(nil).UpdateConfig), arg0, arg1, arg2)
}

// UpdateSecret mocks base method
func (m *MockBackendClient) UpdateSecret(arg0 string, arg1 uint64, arg2 swarm.SecretSpec) error {
	ret := m.ctrl.Call(m, "UpdateSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSecret indicates an expected call of UpdateSecret
func (mr *MockBackendClientMockRecorder) UpdateSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecret", reflect.TypeOf((*MockBackendClient)(nil).UpdateSecret), arg0, arg1, arg2)
}

// UpdateService mocks base method
func (m *MockBackendClient) UpdateService(arg0 string, arg1 uint64, arg2 swarm.ServiceSpec, arg3 types.ServiceUpdateOptions, arg4 bool) (*types.ServiceUpdateResponse, error) {
	ret := m.ctrl.Call(m, "UpdateService", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*types.ServiceUpdateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateService indicates an expected call of UpdateService
func (mr *MockBackendClientMockRecorder) UpdateService(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateService", reflect.TypeOf((*MockBackendClient)(nil).UpdateService), arg0, arg1, arg2, arg3, arg4)
}

// UpdateStack mocks base method
func (m *MockBackendClient) UpdateStack(arg0 string, arg1 types0.StackSpec, arg2 uint64) error {
	ret := m.ctrl.Call(m, "UpdateStack", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStack indicates an expected call of UpdateStack
func (mr *MockBackendClientMockRecorder) UpdateStack(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStack", reflect.TypeOf((*MockBackendClient)(nil).UpdateStack), arg0, arg1, arg2)
}
