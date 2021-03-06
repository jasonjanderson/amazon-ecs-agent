// Copyright 2015-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/dockerclient/dockeriface (interfaces: Client)

// Package mock_dockeriface is a generated GoMock package.
package mock_dockeriface

import (
	context "context"
	reflect "reflect"

	go_dockerclient "github.com/fsouza/go-dockerclient"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// AddEventListener mocks base method
func (m *MockClient) AddEventListener(arg0 chan<- *go_dockerclient.APIEvents) error {
	ret := m.ctrl.Call(m, "AddEventListener", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEventListener indicates an expected call of AddEventListener
func (mr *MockClientMockRecorder) AddEventListener(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEventListener", reflect.TypeOf((*MockClient)(nil).AddEventListener), arg0)
}

// CreateContainer mocks base method
func (m *MockClient) CreateContainer(arg0 go_dockerclient.CreateContainerOptions) (*go_dockerclient.Container, error) {
	ret := m.ctrl.Call(m, "CreateContainer", arg0)
	ret0, _ := ret[0].(*go_dockerclient.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContainer indicates an expected call of CreateContainer
func (mr *MockClientMockRecorder) CreateContainer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockClient)(nil).CreateContainer), arg0)
}

// ImportImage mocks base method
func (m *MockClient) ImportImage(arg0 go_dockerclient.ImportImageOptions) error {
	ret := m.ctrl.Call(m, "ImportImage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportImage indicates an expected call of ImportImage
func (mr *MockClientMockRecorder) ImportImage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportImage", reflect.TypeOf((*MockClient)(nil).ImportImage), arg0)
}

// InspectContainer mocks base method
func (m *MockClient) InspectContainer(arg0 string) (*go_dockerclient.Container, error) {
	ret := m.ctrl.Call(m, "InspectContainer", arg0)
	ret0, _ := ret[0].(*go_dockerclient.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectContainer indicates an expected call of InspectContainer
func (mr *MockClientMockRecorder) InspectContainer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectContainer", reflect.TypeOf((*MockClient)(nil).InspectContainer), arg0)
}

// InspectContainerWithContext mocks base method
func (m *MockClient) InspectContainerWithContext(arg0 string, arg1 context.Context) (*go_dockerclient.Container, error) {
	ret := m.ctrl.Call(m, "InspectContainerWithContext", arg0, arg1)
	ret0, _ := ret[0].(*go_dockerclient.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectContainerWithContext indicates an expected call of InspectContainerWithContext
func (mr *MockClientMockRecorder) InspectContainerWithContext(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectContainerWithContext", reflect.TypeOf((*MockClient)(nil).InspectContainerWithContext), arg0, arg1)
}

// InspectImage mocks base method
func (m *MockClient) InspectImage(arg0 string) (*go_dockerclient.Image, error) {
	ret := m.ctrl.Call(m, "InspectImage", arg0)
	ret0, _ := ret[0].(*go_dockerclient.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectImage indicates an expected call of InspectImage
func (mr *MockClientMockRecorder) InspectImage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectImage", reflect.TypeOf((*MockClient)(nil).InspectImage), arg0)
}

// ListContainers mocks base method
func (m *MockClient) ListContainers(arg0 go_dockerclient.ListContainersOptions) ([]go_dockerclient.APIContainers, error) {
	ret := m.ctrl.Call(m, "ListContainers", arg0)
	ret0, _ := ret[0].([]go_dockerclient.APIContainers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContainers indicates an expected call of ListContainers
func (mr *MockClientMockRecorder) ListContainers(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*MockClient)(nil).ListContainers), arg0)
}

// LoadImage mocks base method
func (m *MockClient) LoadImage(arg0 go_dockerclient.LoadImageOptions) error {
	ret := m.ctrl.Call(m, "LoadImage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadImage indicates an expected call of LoadImage
func (mr *MockClientMockRecorder) LoadImage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadImage", reflect.TypeOf((*MockClient)(nil).LoadImage), arg0)
}

// Ping mocks base method
func (m *MockClient) Ping() error {
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockClientMockRecorder) Ping() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockClient)(nil).Ping))
}

// PullImage mocks base method
func (m *MockClient) PullImage(arg0 go_dockerclient.PullImageOptions, arg1 go_dockerclient.AuthConfiguration) error {
	ret := m.ctrl.Call(m, "PullImage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PullImage indicates an expected call of PullImage
func (mr *MockClientMockRecorder) PullImage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullImage", reflect.TypeOf((*MockClient)(nil).PullImage), arg0, arg1)
}

// RemoveContainer mocks base method
func (m *MockClient) RemoveContainer(arg0 go_dockerclient.RemoveContainerOptions) error {
	ret := m.ctrl.Call(m, "RemoveContainer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainer indicates an expected call of RemoveContainer
func (mr *MockClientMockRecorder) RemoveContainer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveContainer", reflect.TypeOf((*MockClient)(nil).RemoveContainer), arg0)
}

// RemoveEventListener mocks base method
func (m *MockClient) RemoveEventListener(arg0 chan *go_dockerclient.APIEvents) error {
	ret := m.ctrl.Call(m, "RemoveEventListener", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveEventListener indicates an expected call of RemoveEventListener
func (mr *MockClientMockRecorder) RemoveEventListener(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEventListener", reflect.TypeOf((*MockClient)(nil).RemoveEventListener), arg0)
}

// RemoveImage mocks base method
func (m *MockClient) RemoveImage(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveImage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveImage indicates an expected call of RemoveImage
func (mr *MockClientMockRecorder) RemoveImage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveImage", reflect.TypeOf((*MockClient)(nil).RemoveImage), arg0)
}

// StartContainer mocks base method
func (m *MockClient) StartContainer(arg0 string, arg1 *go_dockerclient.HostConfig) error {
	ret := m.ctrl.Call(m, "StartContainer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartContainer indicates an expected call of StartContainer
func (mr *MockClientMockRecorder) StartContainer(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartContainer", reflect.TypeOf((*MockClient)(nil).StartContainer), arg0, arg1)
}

// StartContainerWithContext mocks base method
func (m *MockClient) StartContainerWithContext(arg0 string, arg1 *go_dockerclient.HostConfig, arg2 context.Context) error {
	ret := m.ctrl.Call(m, "StartContainerWithContext", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartContainerWithContext indicates an expected call of StartContainerWithContext
func (mr *MockClientMockRecorder) StartContainerWithContext(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartContainerWithContext", reflect.TypeOf((*MockClient)(nil).StartContainerWithContext), arg0, arg1, arg2)
}

// Stats mocks base method
func (m *MockClient) Stats(arg0 go_dockerclient.StatsOptions) error {
	ret := m.ctrl.Call(m, "Stats", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stats indicates an expected call of Stats
func (mr *MockClientMockRecorder) Stats(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockClient)(nil).Stats), arg0)
}

// StopContainer mocks base method
func (m *MockClient) StopContainer(arg0 string, arg1 uint) error {
	ret := m.ctrl.Call(m, "StopContainer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopContainer indicates an expected call of StopContainer
func (mr *MockClientMockRecorder) StopContainer(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopContainer", reflect.TypeOf((*MockClient)(nil).StopContainer), arg0, arg1)
}

// StopContainerWithContext mocks base method
func (m *MockClient) StopContainerWithContext(arg0 string, arg1 uint, arg2 context.Context) error {
	ret := m.ctrl.Call(m, "StopContainerWithContext", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// StopContainerWithContext indicates an expected call of StopContainerWithContext
func (mr *MockClientMockRecorder) StopContainerWithContext(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopContainerWithContext", reflect.TypeOf((*MockClient)(nil).StopContainerWithContext), arg0, arg1, arg2)
}

// VersionWithContext mocks base method
func (m *MockClient) VersionWithContext(arg0 context.Context) (*go_dockerclient.Env, error) {
	ret := m.ctrl.Call(m, "VersionWithContext", arg0)
	ret0, _ := ret[0].(*go_dockerclient.Env)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VersionWithContext indicates an expected call of VersionWithContext
func (mr *MockClientMockRecorder) VersionWithContext(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VersionWithContext", reflect.TypeOf((*MockClient)(nil).VersionWithContext), arg0)
}
