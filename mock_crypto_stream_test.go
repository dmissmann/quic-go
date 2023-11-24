// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quic-go/quic-go (interfaces: CryptoStream)
//
// Generated by this command:
//
//	mockgen -typed -build_flags=-tags=gomock -package quic -self_package github.com/quic-go/quic-go -destination mock_crypto_stream_test.go github.com/quic-go/quic-go CryptoStream
//
// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	protocol "github.com/dmissmann/quic-go/internal/protocol"
	wire "github.com/dmissmann/quic-go/internal/wire"
	gomock "go.uber.org/mock/gomock"
)

// MockCryptoStream is a mock of CryptoStream interface.
type MockCryptoStream struct {
	ctrl     *gomock.Controller
	recorder *MockCryptoStreamMockRecorder
}

// MockCryptoStreamMockRecorder is the mock recorder for MockCryptoStream.
type MockCryptoStreamMockRecorder struct {
	mock *MockCryptoStream
}

// NewMockCryptoStream creates a new mock instance.
func NewMockCryptoStream(ctrl *gomock.Controller) *MockCryptoStream {
	mock := &MockCryptoStream{ctrl: ctrl}
	mock.recorder = &MockCryptoStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptoStream) EXPECT() *MockCryptoStreamMockRecorder {
	return m.recorder
}

// Finish mocks base method.
func (m *MockCryptoStream) Finish() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finish")
	ret0, _ := ret[0].(error)
	return ret0
}

// Finish indicates an expected call of Finish.
func (mr *MockCryptoStreamMockRecorder) Finish() *CryptoStreamFinishCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finish", reflect.TypeOf((*MockCryptoStream)(nil).Finish))
	return &CryptoStreamFinishCall{Call: call}
}

// CryptoStreamFinishCall wrap *gomock.Call
type CryptoStreamFinishCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoStreamFinishCall) Return(arg0 error) *CryptoStreamFinishCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoStreamFinishCall) Do(f func() error) *CryptoStreamFinishCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoStreamFinishCall) DoAndReturn(f func() error) *CryptoStreamFinishCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetCryptoData mocks base method.
func (m *MockCryptoStream) GetCryptoData() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCryptoData")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetCryptoData indicates an expected call of GetCryptoData.
func (mr *MockCryptoStreamMockRecorder) GetCryptoData() *CryptoStreamGetCryptoDataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCryptoData", reflect.TypeOf((*MockCryptoStream)(nil).GetCryptoData))
	return &CryptoStreamGetCryptoDataCall{Call: call}
}

// CryptoStreamGetCryptoDataCall wrap *gomock.Call
type CryptoStreamGetCryptoDataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoStreamGetCryptoDataCall) Return(arg0 []byte) *CryptoStreamGetCryptoDataCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoStreamGetCryptoDataCall) Do(f func() []byte) *CryptoStreamGetCryptoDataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoStreamGetCryptoDataCall) DoAndReturn(f func() []byte) *CryptoStreamGetCryptoDataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HandleCryptoFrame mocks base method.
func (m *MockCryptoStream) HandleCryptoFrame(arg0 *wire.CryptoFrame) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleCryptoFrame", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleCryptoFrame indicates an expected call of HandleCryptoFrame.
func (mr *MockCryptoStreamMockRecorder) HandleCryptoFrame(arg0 any) *CryptoStreamHandleCryptoFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleCryptoFrame", reflect.TypeOf((*MockCryptoStream)(nil).HandleCryptoFrame), arg0)
	return &CryptoStreamHandleCryptoFrameCall{Call: call}
}

// CryptoStreamHandleCryptoFrameCall wrap *gomock.Call
type CryptoStreamHandleCryptoFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoStreamHandleCryptoFrameCall) Return(arg0 error) *CryptoStreamHandleCryptoFrameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoStreamHandleCryptoFrameCall) Do(f func(*wire.CryptoFrame) error) *CryptoStreamHandleCryptoFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoStreamHandleCryptoFrameCall) DoAndReturn(f func(*wire.CryptoFrame) error) *CryptoStreamHandleCryptoFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HasData mocks base method.
func (m *MockCryptoStream) HasData() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasData")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasData indicates an expected call of HasData.
func (mr *MockCryptoStreamMockRecorder) HasData() *CryptoStreamHasDataCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasData", reflect.TypeOf((*MockCryptoStream)(nil).HasData))
	return &CryptoStreamHasDataCall{Call: call}
}

// CryptoStreamHasDataCall wrap *gomock.Call
type CryptoStreamHasDataCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoStreamHasDataCall) Return(arg0 bool) *CryptoStreamHasDataCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoStreamHasDataCall) Do(f func() bool) *CryptoStreamHasDataCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoStreamHasDataCall) DoAndReturn(f func() bool) *CryptoStreamHasDataCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PopCryptoFrame mocks base method.
func (m *MockCryptoStream) PopCryptoFrame(arg0 protocol.ByteCount) *wire.CryptoFrame {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopCryptoFrame", arg0)
	ret0, _ := ret[0].(*wire.CryptoFrame)
	return ret0
}

// PopCryptoFrame indicates an expected call of PopCryptoFrame.
func (mr *MockCryptoStreamMockRecorder) PopCryptoFrame(arg0 any) *CryptoStreamPopCryptoFrameCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopCryptoFrame", reflect.TypeOf((*MockCryptoStream)(nil).PopCryptoFrame), arg0)
	return &CryptoStreamPopCryptoFrameCall{Call: call}
}

// CryptoStreamPopCryptoFrameCall wrap *gomock.Call
type CryptoStreamPopCryptoFrameCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoStreamPopCryptoFrameCall) Return(arg0 *wire.CryptoFrame) *CryptoStreamPopCryptoFrameCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoStreamPopCryptoFrameCall) Do(f func(protocol.ByteCount) *wire.CryptoFrame) *CryptoStreamPopCryptoFrameCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoStreamPopCryptoFrameCall) DoAndReturn(f func(protocol.ByteCount) *wire.CryptoFrame) *CryptoStreamPopCryptoFrameCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Write mocks base method.
func (m *MockCryptoStream) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockCryptoStreamMockRecorder) Write(arg0 any) *CryptoStreamWriteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockCryptoStream)(nil).Write), arg0)
	return &CryptoStreamWriteCall{Call: call}
}

// CryptoStreamWriteCall wrap *gomock.Call
type CryptoStreamWriteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *CryptoStreamWriteCall) Return(arg0 int, arg1 error) *CryptoStreamWriteCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *CryptoStreamWriteCall) Do(f func([]byte) (int, error)) *CryptoStreamWriteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *CryptoStreamWriteCall) DoAndReturn(f func([]byte) (int, error)) *CryptoStreamWriteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
