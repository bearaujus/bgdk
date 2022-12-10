// Code generated by MockGen. DO NOT EDIT.
// Source: worker/init.go

// Package mockWorker is a generated GoMock package.
package mockWorker

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWorker is a mock of Worker interface.
type MockWorker struct {
	ctrl     *gomock.Controller
	recorder *MockWorkerMockRecorder
}

// MockWorkerMockRecorder is the mock recorder for MockWorker.
type MockWorkerMockRecorder struct {
	mock *MockWorker
}

// NewMockWorker creates a new mock instance.
func NewMockWorker(ctrl *gomock.Controller) *MockWorker {
	mock := &MockWorker{ctrl: ctrl}
	mock.recorder = &MockWorkerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorker) EXPECT() *MockWorkerMockRecorder {
	return m.recorder
}

// PushJob mocks base method.
func (m *MockWorker) PushJob(job func() error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PushJob", job)
}

// PushJob indicates an expected call of PushJob.
func (mr *MockWorkerMockRecorder) PushJob(job interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PushJob", reflect.TypeOf((*MockWorker)(nil).PushJob), job)
}

// Shutdown mocks base method.
func (m *MockWorker) Shutdown() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Shutdown")
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockWorkerMockRecorder) Shutdown() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockWorker)(nil).Shutdown))
}

// Wait mocks base method.
func (m *MockWorker) Wait() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Wait")
}

// Wait indicates an expected call of Wait.
func (mr *MockWorkerMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockWorker)(nil).Wait))
}