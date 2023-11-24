// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: SealingManager)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/quic-go/quic-go -destination mock_sealing_manager_test.go github.com/quic-go/quic-go SealingManager
//
// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	handshake "github.com/dmissmann/quic-go/internal/handshake"
	gomock "go.uber.org/mock/gomock"
)

// MockSealingManager is a mock of SealingManager interface.
type MockSealingManager struct {
	ctrl     *gomock.Controller
	recorder *MockSealingManagerMockRecorder
}

// MockSealingManagerMockRecorder is the mock recorder for MockSealingManager.
type MockSealingManagerMockRecorder struct {
	mock *MockSealingManager
}

// NewMockSealingManager creates a new mock instance.
func NewMockSealingManager(ctrl *gomock.Controller) *MockSealingManager {
	mock := &MockSealingManager{ctrl: ctrl}
	mock.recorder = &MockSealingManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSealingManager) EXPECT() *MockSealingManagerMockRecorder {
	return m.recorder
}

// Get0RTTSealer mocks base method.
func (m *MockSealingManager) Get0RTTSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get0RTTSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get0RTTSealer indicates an expected call of Get0RTTSealer.
func (mr *MockSealingManagerMockRecorder) Get0RTTSealer() *SealingManagerGet0RTTSealerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get0RTTSealer", reflect.TypeOf((*MockSealingManager)(nil).Get0RTTSealer))
	return &SealingManagerGet0RTTSealerCall{Call: call}
}

// SealingManagerGet0RTTSealerCall wrap *gomock.Call
type SealingManagerGet0RTTSealerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SealingManagerGet0RTTSealerCall) Return(arg0 handshake.LongHeaderSealer, arg1 error) *SealingManagerGet0RTTSealerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SealingManagerGet0RTTSealerCall) Do(f func() (handshake.LongHeaderSealer, error)) *SealingManagerGet0RTTSealerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SealingManagerGet0RTTSealerCall) DoAndReturn(f func() (handshake.LongHeaderSealer, error)) *SealingManagerGet0RTTSealerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get1RTTSealer mocks base method.
func (m *MockSealingManager) Get1RTTSealer() (handshake.ShortHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get1RTTSealer")
	ret0, _ := ret[0].(handshake.ShortHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get1RTTSealer indicates an expected call of Get1RTTSealer.
func (mr *MockSealingManagerMockRecorder) Get1RTTSealer() *SealingManagerGet1RTTSealerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get1RTTSealer", reflect.TypeOf((*MockSealingManager)(nil).Get1RTTSealer))
	return &SealingManagerGet1RTTSealerCall{Call: call}
}

// SealingManagerGet1RTTSealerCall wrap *gomock.Call
type SealingManagerGet1RTTSealerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SealingManagerGet1RTTSealerCall) Return(arg0 handshake.ShortHeaderSealer, arg1 error) *SealingManagerGet1RTTSealerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SealingManagerGet1RTTSealerCall) Do(f func() (handshake.ShortHeaderSealer, error)) *SealingManagerGet1RTTSealerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SealingManagerGet1RTTSealerCall) DoAndReturn(f func() (handshake.ShortHeaderSealer, error)) *SealingManagerGet1RTTSealerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetHandshakeSealer mocks base method.
func (m *MockSealingManager) GetHandshakeSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHandshakeSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHandshakeSealer indicates an expected call of GetHandshakeSealer.
func (mr *MockSealingManagerMockRecorder) GetHandshakeSealer() *SealingManagerGetHandshakeSealerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHandshakeSealer", reflect.TypeOf((*MockSealingManager)(nil).GetHandshakeSealer))
	return &SealingManagerGetHandshakeSealerCall{Call: call}
}

// SealingManagerGetHandshakeSealerCall wrap *gomock.Call
type SealingManagerGetHandshakeSealerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SealingManagerGetHandshakeSealerCall) Return(arg0 handshake.LongHeaderSealer, arg1 error) *SealingManagerGetHandshakeSealerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SealingManagerGetHandshakeSealerCall) Do(f func() (handshake.LongHeaderSealer, error)) *SealingManagerGetHandshakeSealerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SealingManagerGetHandshakeSealerCall) DoAndReturn(f func() (handshake.LongHeaderSealer, error)) *SealingManagerGetHandshakeSealerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetInitialSealer mocks base method.
func (m *MockSealingManager) GetInitialSealer() (handshake.LongHeaderSealer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInitialSealer")
	ret0, _ := ret[0].(handshake.LongHeaderSealer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInitialSealer indicates an expected call of GetInitialSealer.
func (mr *MockSealingManagerMockRecorder) GetInitialSealer() *SealingManagerGetInitialSealerCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInitialSealer", reflect.TypeOf((*MockSealingManager)(nil).GetInitialSealer))
	return &SealingManagerGetInitialSealerCall{Call: call}
}

// SealingManagerGetInitialSealerCall wrap *gomock.Call
type SealingManagerGetInitialSealerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SealingManagerGetInitialSealerCall) Return(arg0 handshake.LongHeaderSealer, arg1 error) *SealingManagerGetInitialSealerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SealingManagerGetInitialSealerCall) Do(f func() (handshake.LongHeaderSealer, error)) *SealingManagerGetInitialSealerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SealingManagerGetInitialSealerCall) DoAndReturn(f func() (handshake.LongHeaderSealer, error)) *SealingManagerGetInitialSealerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
