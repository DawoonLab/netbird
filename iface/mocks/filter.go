// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/netbirdio/netbird/iface (interfaces: PacketFilter)

// Package mocks is a generated GoMock package.
package mocks

import (
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPacketFilter is a mock of PacketFilter interface.
type MockPacketFilter struct {
	ctrl     *gomock.Controller
	recorder *MockPacketFilterMockRecorder
}

// MockPacketFilterMockRecorder is the mock recorder for MockPacketFilter.
type MockPacketFilterMockRecorder struct {
	mock *MockPacketFilter
}

// NewMockPacketFilter creates a new mock instance.
func NewMockPacketFilter(ctrl *gomock.Controller) *MockPacketFilter {
	mock := &MockPacketFilter{ctrl: ctrl}
	mock.recorder = &MockPacketFilterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacketFilter) EXPECT() *MockPacketFilterMockRecorder {
	return m.recorder
}

// DropInput mocks base method.
func (m *MockPacketFilter) DropOutgoing(arg0 []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropOutgoing", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DropInput indicates an expected call of DropInput.
func (mr *MockPacketFilterMockRecorder) DropInput(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropOutgoing", reflect.TypeOf((*MockPacketFilter)(nil).DropOutgoing), arg0)
}

// DropOutput mocks base method.
func (m *MockPacketFilter) DropIncoming(arg0 []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DropIncoming", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DropOutput indicates an expected call of DropOutput.
func (mr *MockPacketFilterMockRecorder) DropOutput(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropIncoming", reflect.TypeOf((*MockPacketFilter)(nil).DropIncoming), arg0)
}

// SetNetwork mocks base method.
func (m *MockPacketFilter) SetNetwork(arg0 *net.IPNet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetNetwork", arg0)
}

// SetNetwork indicates an expected call of SetNetwork.
func (mr *MockPacketFilterMockRecorder) SetNetwork(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetNetwork", reflect.TypeOf((*MockPacketFilter)(nil).SetNetwork), arg0)
}
