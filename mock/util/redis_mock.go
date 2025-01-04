// Code generated by MockGen. DO NOT EDIT.
// Source: util/redis/type.go
//
// Generated by this command:
//
//	mockgen -package=mock_util -source=util/redis/type.go -destination=mock/util/redis_mock.go -typed
//

// Package mock_util is a generated GoMock package.
package mock_util

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockRedisManager is a mock of RedisManager interface.
type MockRedisManager struct {
	ctrl     *gomock.Controller
	recorder *MockRedisManagerMockRecorder
	isgomock struct{}
}

// MockRedisManagerMockRecorder is the mock recorder for MockRedisManager.
type MockRedisManagerMockRecorder struct {
	mock *MockRedisManager
}

// NewMockRedisManager creates a new mock instance.
func NewMockRedisManager(ctrl *gomock.Controller) *MockRedisManager {
	mock := &MockRedisManager{ctrl: ctrl}
	mock.recorder = &MockRedisManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisManager) EXPECT() *MockRedisManagerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRedisManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockRedisManagerMockRecorder) Close() *MockRedisManagerCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRedisManager)(nil).Close))
	return &MockRedisManagerCloseCall{Call: call}
}

// MockRedisManagerCloseCall wrap *gomock.Call
type MockRedisManagerCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRedisManagerCloseCall) Return(arg0 error) *MockRedisManagerCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRedisManagerCloseCall) Do(f func() error) *MockRedisManagerCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRedisManagerCloseCall) DoAndReturn(f func() error) *MockRedisManagerCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Delete mocks base method.
func (m *MockRedisManager) Delete(ctx context.Context, key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockRedisManagerMockRecorder) Delete(ctx, key any) *MockRedisManagerDeleteCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockRedisManager)(nil).Delete), ctx, key)
	return &MockRedisManagerDeleteCall{Call: call}
}

// MockRedisManagerDeleteCall wrap *gomock.Call
type MockRedisManagerDeleteCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRedisManagerDeleteCall) Return(arg0 error) *MockRedisManagerDeleteCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRedisManagerDeleteCall) Do(f func(context.Context, string) error) *MockRedisManagerDeleteCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRedisManagerDeleteCall) DoAndReturn(f func(context.Context, string) error) *MockRedisManagerDeleteCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Get mocks base method.
func (m *MockRedisManager) Get(ctx context.Context, key string, dest any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, key, dest)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockRedisManagerMockRecorder) Get(ctx, key, dest any) *MockRedisManagerGetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRedisManager)(nil).Get), ctx, key, dest)
	return &MockRedisManagerGetCall{Call: call}
}

// MockRedisManagerGetCall wrap *gomock.Call
type MockRedisManagerGetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRedisManagerGetCall) Return(arg0 error) *MockRedisManagerGetCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRedisManagerGetCall) Do(f func(context.Context, string, any) error) *MockRedisManagerGetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRedisManagerGetCall) DoAndReturn(f func(context.Context, string, any) error) *MockRedisManagerGetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Set mocks base method.
func (m *MockRedisManager) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, key, value, expiration)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockRedisManagerMockRecorder) Set(ctx, key, value, expiration any) *MockRedisManagerSetCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockRedisManager)(nil).Set), ctx, key, value, expiration)
	return &MockRedisManagerSetCall{Call: call}
}

// MockRedisManagerSetCall wrap *gomock.Call
type MockRedisManagerSetCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockRedisManagerSetCall) Return(arg0 error) *MockRedisManagerSetCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockRedisManagerSetCall) Do(f func(context.Context, string, any, time.Duration) error) *MockRedisManagerSetCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockRedisManagerSetCall) DoAndReturn(f func(context.Context, string, any, time.Duration) error) *MockRedisManagerSetCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}