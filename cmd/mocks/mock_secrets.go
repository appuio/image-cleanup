// Code generated by MockGen. DO NOT EDIT.
// Source: secrets_test.go

// Package mocks is a generated GoMock package.
package mocks

import (
	cfg "github.com/appuio/seiso/cfg"
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	reflect "reflect"
	time "time"
)

// MockTestSecretResources is a mock of TestSecretResources interface
type MockTestSecretResources struct {
	ctrl     *gomock.Controller
	recorder *MockTestSecretResourcesMockRecorder
}

// MockTestSecretResourcesMockRecorder is the mock recorder for MockTestSecretResources
type MockTestSecretResourcesMockRecorder struct {
	mock *MockTestSecretResources
}

// NewMockTestSecretResources creates a new mock instance
func NewMockTestSecretResources(ctrl *gomock.Controller) *MockTestSecretResources {
	mock := &MockTestSecretResources{ctrl: ctrl}
	mock.recorder = &MockTestSecretResourcesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTestSecretResources) EXPECT() *MockTestSecretResourcesMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockTestSecretResources) Delete(resourceSelectorFunc cfg.ResourceNamespaceSelector) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", resourceSelectorFunc)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTestSecretResourcesMockRecorder) Delete(resourceSelectorFunc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTestSecretResources)(nil).Delete), resourceSelectorFunc)
}

// Print mocks base method
func (m *MockTestSecretResources) Print(batch bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Print", batch)
}

// Print indicates an expected call of Print
func (mr *MockTestSecretResourcesMockRecorder) Print(batch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockTestSecretResources)(nil).Print), batch)
}

// GetNamespace mocks base method
func (m *MockTestSecretResources) GetNamespace() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetNamespace indicates an expected call of GetNamespace
func (mr *MockTestSecretResourcesMockRecorder) GetNamespace() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNamespace", reflect.TypeOf((*MockTestSecretResources)(nil).GetNamespace))
}

// FilterByTime mocks base method
func (m *MockTestSecretResources) FilterByTime(olderThan time.Time) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FilterByTime", olderThan)
}

// FilterByTime indicates an expected call of FilterByTime
func (mr *MockTestSecretResourcesMockRecorder) FilterByTime(olderThan interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterByTime", reflect.TypeOf((*MockTestSecretResources)(nil).FilterByTime), olderThan)
}

// FilterUsed mocks base method
func (m *MockTestSecretResources) FilterUsed() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterUsed")
	ret0, _ := ret[0].(error)
	return ret0
}

// FilterUsed indicates an expected call of FilterUsed
func (mr *MockTestSecretResourcesMockRecorder) FilterUsed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterUsed", reflect.TypeOf((*MockTestSecretResources)(nil).FilterUsed))
}

// FilterByMaxCount mocks base method
func (m *MockTestSecretResources) FilterByMaxCount(keep int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FilterByMaxCount", keep)
}

// FilterByMaxCount indicates an expected call of FilterByMaxCount
func (mr *MockTestSecretResourcesMockRecorder) FilterByMaxCount(keep interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterByMaxCount", reflect.TypeOf((*MockTestSecretResources)(nil).FilterByMaxCount), keep)
}

// LoadSecrets mocks base method
func (m *MockTestSecretResources) LoadSecrets(listOptions v1.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadSecrets", listOptions)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadSecrets indicates an expected call of LoadSecrets
func (mr *MockTestSecretResourcesMockRecorder) LoadSecrets(listOptions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadSecrets", reflect.TypeOf((*MockTestSecretResources)(nil).LoadSecrets), listOptions)
}

// ListSecrets mocks base method
func (m *MockTestSecretResources) ListSecrets() ([]string, []string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].([]string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListSecrets indicates an expected call of ListSecrets
func (mr *MockTestSecretResourcesMockRecorder) ListSecrets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockTestSecretResources)(nil).ListSecrets))
}