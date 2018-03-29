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
//
// Source: exec.go in package sysctl
// Code generated by MockGen. DO NOT EDIT.

// Package sysctl is a generated GoMock package.
package sysctl

import (
	reflect "reflect"

	cmd "github.com/aws/amazon-ecs-init/ecs-init/cmd"
	gomock "github.com/golang/mock/gomock"
)

// MockExec is a mock of Exec interface
type MockExec struct {
	ctrl     *gomock.Controller
	recorder *MockExecMockRecorder
}

// MockExecMockRecorder is the mock recorder for MockExec
type MockExecMockRecorder struct {
	mock *MockExec
}

// NewMockExec creates a new mock instance
func NewMockExec(ctrl *gomock.Controller) *MockExec {
	mock := &MockExec{ctrl: ctrl}
	mock.recorder = &MockExecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExec) EXPECT() *MockExecMockRecorder {
	return m.recorder
}

// LookPath mocks base method
func (m *MockExec) LookPath(file string) (string, error) {
	ret := m.ctrl.Call(m, "LookPath", file)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LookPath indicates an expected call of LookPath
func (mr *MockExecMockRecorder) LookPath(file interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LookPath", reflect.TypeOf((*MockExec)(nil).LookPath), file)
}

// Command mocks base method
func (m *MockExec) Command(name string, arg ...string) cmd.Cmd {
	varargs := []interface{}{name}
	for _, a := range arg {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Command", varargs...)
	ret0, _ := ret[0].(cmd.Cmd)
	return ret0
}

// Command indicates an expected call of Command
func (mr *MockExecMockRecorder) Command(name interface{}, arg ...interface{}) *gomock.Call {
	varargs := append([]interface{}{name}, arg...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockExec)(nil).Command), varargs...)
}
