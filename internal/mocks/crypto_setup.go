// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/fkwhite/Quic_GO/internal/handshake (interfaces: CryptoSetup)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	handshake "github.com/fkwhite/Quic_GO/internal/handshake"
	protocol "github.com/fkwhite/Quic_GO/internal/protocol"
	qtls "github.com/fkwhite/Quic_GO/internal/qtls"
)

// MockCryptoSetup is a mock of CryptoSetup interface.
type MockCryptoSetup struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoSetupMockRecorder
}

// MockCryptoSetupMockRecorder is the mock recorder for MockCryptoSetup.
type MockCryptoSetupMockRecorder struct {
	mock *MockCryptoSetup
}

// NewMockCryptoSetup creates a new mock instance.
func NewMockCryptoSetup(ctrl *gomock.Controller) *MockCryptoSetup {
	mock := &MockCryptoSetup{ctrl: ctrl}
	mock.recorder = &MockCryptoSetupMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoSetup) EXPECT() *MockCryptoSetupMockRecorder {
	return m.recorder
}

// ChangeConnectionID mocks base method.
func (m *MockCryptoSetup) ChangeConnectionID(arg0 protocol.ConnectionID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ChangeConnectionID", arg0)
}

// ChangeConnectionID indicates an expected call of ChangeConnectionID.
func (mr *MockCryptoSetupMockRecorder) ChangeConnectionID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeConnectionID", reflect.TypeOf((*MockCryptoSetup)(nil).ChangeConnectionID), arg0)
}

// Close mocks base method.
func (m *MockCryptoSetup) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCryptoSetupMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCryptoSetup)(nil).Close))
}

// ConnectionState mocks base method.
func (m *MockCryptoSetup) ConnectionState() qtls.ConnectionState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectionState")
	ret0, _ := ret[0].(qtls.ConnectionState)
	return ret0
}

// ConnectionState indicates an expected call of ConnectionState.
func (mr *MockCryptoSetupMockRecorder) ConnectionState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionState", reflect.TypeOf((*MockCryptoSetup)(nil).ConnectionState))
}

// Get0RTTOpener mocks base method.
func (m *MockCryptoSetup) Get0RTTOpener() (handshake.LongHeaderOpener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get0RTTOpener")
	ret0, _ := ret[0].(handshake.LongHeaderOpener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get0RTTOpener indicates an expected call of Get0RTTOpener.
func (mr *MockCryptoSetupMockRecorder) Get0RTTOpener() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get0RTTOpener", reflect.TypeOf((*MockCryptoSetup)(nil).Get0RTTOpener))
}

// Get0RTTSealer mocks base method.
func (m *MockCryptoSetup) Get0RTTSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get0RTTSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get0RTTSealer indicates an expected call of Get0RTTSealer.
func (mr *MockCryptoSetupMockRecorder) Get0RTTSealer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get0RTTSealer", reflect.TypeOf((*MockCryptoSetup)(nil).Get0RTTSealer))
}

// Get1RTTOpener mocks base method.
func (m *MockCryptoSetup) Get1RTTOpener() (handshake.ShortHeaderOpener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get1RTTOpener")
	ret0, _ := ret[0].(handshake.ShortHeaderOpener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get1RTTOpener indicates an expected call of Get1RTTOpener.
func (mr *MockCryptoSetupMockRecorder) Get1RTTOpener() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get1RTTOpener", reflect.TypeOf((*MockCryptoSetup)(nil).Get1RTTOpener))
}

// Get1RTTSealer mocks base method.
func (m *MockCryptoSetup) Get1RTTSealer() (handshake.ShortHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get1RTTSealer")
	ret0, _ := ret[0].(handshake.ShortHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get1RTTSealer indicates an expected call of Get1RTTSealer.
func (mr *MockCryptoSetupMockRecorder) Get1RTTSealer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get1RTTSealer", reflect.TypeOf((*MockCryptoSetup)(nil).Get1RTTSealer))
}

// GetHandshakeOpener mocks base method.
func (m *MockCryptoSetup) GetHandshakeOpener() (handshake.LongHeaderOpener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandshakeOpener")
	ret0, _ := ret[0].(handshake.LongHeaderOpener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHandshakeOpener indicates an expected call of GetHandshakeOpener.
func (mr *MockCryptoSetupMockRecorder) GetHandshakeOpener() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandshakeOpener", reflect.TypeOf((*MockCryptoSetup)(nil).GetHandshakeOpener))
}

// GetHandshakeSealer mocks base method.
func (m *MockCryptoSetup) GetHandshakeSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandshakeSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHandshakeSealer indicates an expected call of GetHandshakeSealer.
func (mr *MockCryptoSetupMockRecorder) GetHandshakeSealer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandshakeSealer", reflect.TypeOf((*MockCryptoSetup)(nil).GetHandshakeSealer))
}

// GetInitialOpener mocks base method.
func (m *MockCryptoSetup) GetInitialOpener() (handshake.LongHeaderOpener, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInitialOpener")
	ret0, _ := ret[0].(handshake.LongHeaderOpener)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInitialOpener indicates an expected call of GetInitialOpener.
func (mr *MockCryptoSetupMockRecorder) GetInitialOpener() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInitialOpener", reflect.TypeOf((*MockCryptoSetup)(nil).GetInitialOpener))
}

// GetInitialSealer mocks base method.
func (m *MockCryptoSetup) GetInitialSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInitialSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInitialSealer indicates an expected call of GetInitialSealer.
func (mr *MockCryptoSetupMockRecorder) GetInitialSealer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInitialSealer", reflect.TypeOf((*MockCryptoSetup)(nil).GetInitialSealer))
}

// GetSessionTicket mocks base method.
func (m *MockCryptoSetup) GetSessionTicket() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionTicket")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSessionTicket indicates an expected call of GetSessionTicket.
func (mr *MockCryptoSetupMockRecorder) GetSessionTicket() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionTicket", reflect.TypeOf((*MockCryptoSetup)(nil).GetSessionTicket))
}

// HandleMessage mocks base method.
func (m *MockCryptoSetup) HandleMessage(arg0 []byte, arg1 protocol.EncryptionLevel) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleMessage", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HandleMessage indicates an expected call of HandleMessage.
func (mr *MockCryptoSetupMockRecorder) HandleMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleMessage", reflect.TypeOf((*MockCryptoSetup)(nil).HandleMessage), arg0, arg1)
}

// RunHandshake mocks base method.
func (m *MockCryptoSetup) RunHandshake() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RunHandshake")
}

// RunHandshake indicates an expected call of RunHandshake.
func (mr *MockCryptoSetupMockRecorder) RunHandshake() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunHandshake", reflect.TypeOf((*MockCryptoSetup)(nil).RunHandshake))
}

// SetHandshakeConfirmed mocks base method.
func (m *MockCryptoSetup) SetHandshakeConfirmed() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHandshakeConfirmed")
}

// SetHandshakeConfirmed indicates an expected call of SetHandshakeConfirmed.
func (mr *MockCryptoSetupMockRecorder) SetHandshakeConfirmed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHandshakeConfirmed", reflect.TypeOf((*MockCryptoSetup)(nil).SetHandshakeConfirmed))
}

// SetLargest1RTTAcked mocks base method.
func (m *MockCryptoSetup) SetLargest1RTTAcked(arg0 protocol.PacketNumber) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLargest1RTTAcked", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLargest1RTTAcked indicates an expected call of SetLargest1RTTAcked.
func (mr *MockCryptoSetupMockRecorder) SetLargest1RTTAcked(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLargest1RTTAcked", reflect.TypeOf((*MockCryptoSetup)(nil).SetLargest1RTTAcked), arg0)
}
