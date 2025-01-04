// Code generated by MockGen. DO NOT EDIT.
// Source: internal/modules/user/service/type.go
//
// Generated by this command:
//
//	mockgen -package=mock_service -source=internal/modules/user/service/type.go -destination=mock/service/user_service_mock.go -typed
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	gomock "go.uber.org/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
	isgomock struct{}
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockUserService) ChangePassword(ctx context.Context, id, oldPassword, newPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", ctx, id, oldPassword, newPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockUserServiceMockRecorder) ChangePassword(ctx, id, oldPassword, newPassword any) *MockUserServiceChangePasswordCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockUserService)(nil).ChangePassword), ctx, id, oldPassword, newPassword)
	return &MockUserServiceChangePasswordCall{Call: call}
}

// MockUserServiceChangePasswordCall wrap *gomock.Call
type MockUserServiceChangePasswordCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceChangePasswordCall) Return(arg0 error) *MockUserServiceChangePasswordCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceChangePasswordCall) Do(f func(context.Context, string, string, string) error) *MockUserServiceChangePasswordCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceChangePasswordCall) DoAndReturn(f func(context.Context, string, string, string) error) *MockUserServiceChangePasswordCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUserByEmail mocks base method.
func (m *MockUserService) GetUserByEmail(ctx context.Context, email string) (user_model.GetUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(user_model.GetUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserServiceMockRecorder) GetUserByEmail(ctx, email any) *MockUserServiceGetUserByEmailCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserService)(nil).GetUserByEmail), ctx, email)
	return &MockUserServiceGetUserByEmailCall{Call: call}
}

// MockUserServiceGetUserByEmailCall wrap *gomock.Call
type MockUserServiceGetUserByEmailCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceGetUserByEmailCall) Return(res user_model.GetUserResponse, err error) *MockUserServiceGetUserByEmailCall {
	c.Call = c.Call.Return(res, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceGetUserByEmailCall) Do(f func(context.Context, string) (user_model.GetUserResponse, error)) *MockUserServiceGetUserByEmailCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceGetUserByEmailCall) DoAndReturn(f func(context.Context, string) (user_model.GetUserResponse, error)) *MockUserServiceGetUserByEmailCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetUserByID mocks base method.
func (m *MockUserService) GetUserByID(ctx context.Context, id string) (user_model.GetUserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(user_model.GetUserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserServiceMockRecorder) GetUserByID(ctx, id any) *MockUserServiceGetUserByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserService)(nil).GetUserByID), ctx, id)
	return &MockUserServiceGetUserByIDCall{Call: call}
}

// MockUserServiceGetUserByIDCall wrap *gomock.Call
type MockUserServiceGetUserByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceGetUserByIDCall) Return(res user_model.GetUserResponse, err error) *MockUserServiceGetUserByIDCall {
	c.Call = c.Call.Return(res, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceGetUserByIDCall) Do(f func(context.Context, string) (user_model.GetUserResponse, error)) *MockUserServiceGetUserByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceGetUserByIDCall) DoAndReturn(f func(context.Context, string) (user_model.GetUserResponse, error)) *MockUserServiceGetUserByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Login mocks base method.
func (m *MockUserService) Login(ctx context.Context, email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockUserServiceMockRecorder) Login(ctx, email, password any) *MockUserServiceLoginCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserService)(nil).Login), ctx, email, password)
	return &MockUserServiceLoginCall{Call: call}
}

// MockUserServiceLoginCall wrap *gomock.Call
type MockUserServiceLoginCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceLoginCall) Return(token string, err error) *MockUserServiceLoginCall {
	c.Call = c.Call.Return(token, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceLoginCall) Do(f func(context.Context, string, string) (string, error)) *MockUserServiceLoginCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceLoginCall) DoAndReturn(f func(context.Context, string, string) (string, error)) *MockUserServiceLoginCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Register mocks base method.
func (m *MockUserService) Register(ctx context.Context, email, password, name, address string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, email, password, name, address)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockUserServiceMockRecorder) Register(ctx, email, password, name, address any) *MockUserServiceRegisterCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockUserService)(nil).Register), ctx, email, password, name, address)
	return &MockUserServiceRegisterCall{Call: call}
}

// MockUserServiceRegisterCall wrap *gomock.Call
type MockUserServiceRegisterCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceRegisterCall) Return(id string, err error) *MockUserServiceRegisterCall {
	c.Call = c.Call.Return(id, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceRegisterCall) Do(f func(context.Context, string, string, string, string) (string, error)) *MockUserServiceRegisterCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceRegisterCall) DoAndReturn(f func(context.Context, string, string, string, string) (string, error)) *MockUserServiceRegisterCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateUser mocks base method.
func (m *MockUserService) UpdateUser(ctx context.Context, id, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, id, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserServiceMockRecorder) UpdateUser(ctx, id, name any) *MockUserServiceUpdateUserCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserService)(nil).UpdateUser), ctx, id, name)
	return &MockUserServiceUpdateUserCall{Call: call}
}

// MockUserServiceUpdateUserCall wrap *gomock.Call
type MockUserServiceUpdateUserCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockUserServiceUpdateUserCall) Return(arg0 error) *MockUserServiceUpdateUserCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockUserServiceUpdateUserCall) Do(f func(context.Context, string, string) error) *MockUserServiceUpdateUserCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockUserServiceUpdateUserCall) DoAndReturn(f func(context.Context, string, string) error) *MockUserServiceUpdateUserCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
