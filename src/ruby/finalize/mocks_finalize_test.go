// Code generated by MockGen. DO NOT EDIT.
// Source: finalize.go

// Package finalize_test is a generated GoMock package.
package finalize_test

import (
	gomock "github.com/golang/mock/gomock"
	exec "os/exec"
	reflect "reflect"
)

// MockStager is a mock of Stager interface
type MockStager struct {
	ctrl     *gomock.Controller
	recorder *MockStagerMockRecorder
}

// MockStagerMockRecorder is the mock recorder for MockStager
type MockStagerMockRecorder struct {
	mock *MockStager
}

// NewMockStager creates a new mock instance
func NewMockStager(ctrl *gomock.Controller) *MockStager {
	mock := &MockStager{ctrl: ctrl}
	mock.recorder = &MockStagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStager) EXPECT() *MockStagerMockRecorder {
	return m.recorder
}

// BuildDir mocks base method
func (m *MockStager) BuildDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// BuildDir indicates an expected call of BuildDir
func (mr *MockStagerMockRecorder) BuildDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildDir", reflect.TypeOf((*MockStager)(nil).BuildDir))
}

// DepsIdx mocks base method
func (m *MockStager) DepsIdx() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DepsIdx")
	ret0, _ := ret[0].(string)
	return ret0
}

// DepsIdx indicates an expected call of DepsIdx
func (mr *MockStagerMockRecorder) DepsIdx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepsIdx", reflect.TypeOf((*MockStager)(nil).DepsIdx))
}

// DepDir mocks base method
func (m *MockStager) DepDir() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DepDir")
	ret0, _ := ret[0].(string)
	return ret0
}

// DepDir indicates an expected call of DepDir
func (mr *MockStagerMockRecorder) DepDir() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepDir", reflect.TypeOf((*MockStager)(nil).DepDir))
}

// MockCommand is a mock of Command interface
type MockCommand struct {
	ctrl     *gomock.Controller
	recorder *MockCommandMockRecorder
}

// MockCommandMockRecorder is the mock recorder for MockCommand
type MockCommandMockRecorder struct {
	mock *MockCommand
}

// NewMockCommand creates a new mock instance
func NewMockCommand(ctrl *gomock.Controller) *MockCommand {
	mock := &MockCommand{ctrl: ctrl}
	mock.recorder = &MockCommandMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommand) EXPECT() *MockCommandMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockCommand) Run(arg0 *exec.Cmd) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockCommandMockRecorder) Run(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockCommand)(nil).Run), arg0)
}
