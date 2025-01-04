// Code generated by MockGen. DO NOT EDIT.
// Source: internal/modules/task/repository/type.go
//
// Generated by this command:
//
//	mockgen -package=mock_repository -source=internal/modules/task/repository/type.go -destination=mock/repository/task_repository_mock.go -typed
//

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	task_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/model"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockTaskRepository is a mock of TaskRepository interface.
type MockTaskRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTaskRepositoryMockRecorder
	isgomock struct{}
}

// MockTaskRepositoryMockRecorder is the mock recorder for MockTaskRepository.
type MockTaskRepositoryMockRecorder struct {
	mock *MockTaskRepository
}

// NewMockTaskRepository creates a new mock instance.
func NewMockTaskRepository(ctrl *gomock.Controller) *MockTaskRepository {
	mock := &MockTaskRepository{ctrl: ctrl}
	mock.recorder = &MockTaskRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskRepository) EXPECT() *MockTaskRepositoryMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockTaskRepository) CreateTask(ctx context.Context, userID uuid.UUID, title, description, status string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", ctx, userID, title, description, status)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskRepositoryMockRecorder) CreateTask(ctx, userID, title, description, status any) *MockTaskRepositoryCreateTaskCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTaskRepository)(nil).CreateTask), ctx, userID, title, description, status)
	return &MockTaskRepositoryCreateTaskCall{Call: call}
}

// MockTaskRepositoryCreateTaskCall wrap *gomock.Call
type MockTaskRepositoryCreateTaskCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTaskRepositoryCreateTaskCall) Return(id string, err error) *MockTaskRepositoryCreateTaskCall {
	c.Call = c.Call.Return(id, err)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTaskRepositoryCreateTaskCall) Do(f func(context.Context, uuid.UUID, string, string, string) (string, error)) *MockTaskRepositoryCreateTaskCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTaskRepositoryCreateTaskCall) DoAndReturn(f func(context.Context, uuid.UUID, string, string, string) (string, error)) *MockTaskRepositoryCreateTaskCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteTask mocks base method.
func (m *MockTaskRepository) DeleteTask(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockTaskRepositoryMockRecorder) DeleteTask(ctx, id any) *MockTaskRepositoryDeleteTaskCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockTaskRepository)(nil).DeleteTask), ctx, id)
	return &MockTaskRepositoryDeleteTaskCall{Call: call}
}

// MockTaskRepositoryDeleteTaskCall wrap *gomock.Call
type MockTaskRepositoryDeleteTaskCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTaskRepositoryDeleteTaskCall) Return(arg0 error) *MockTaskRepositoryDeleteTaskCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTaskRepositoryDeleteTaskCall) Do(f func(context.Context, uuid.UUID) error) *MockTaskRepositoryDeleteTaskCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTaskRepositoryDeleteTaskCall) DoAndReturn(f func(context.Context, uuid.UUID) error) *MockTaskRepositoryDeleteTaskCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetAllUserTasks mocks base method.
func (m *MockTaskRepository) GetAllUserTasks(ctx context.Context, userID uuid.UUID) ([]task_model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserTasks", ctx, userID)
	ret0, _ := ret[0].([]task_model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserTasks indicates an expected call of GetAllUserTasks.
func (mr *MockTaskRepositoryMockRecorder) GetAllUserTasks(ctx, userID any) *MockTaskRepositoryGetAllUserTasksCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserTasks", reflect.TypeOf((*MockTaskRepository)(nil).GetAllUserTasks), ctx, userID)
	return &MockTaskRepositoryGetAllUserTasksCall{Call: call}
}

// MockTaskRepositoryGetAllUserTasksCall wrap *gomock.Call
type MockTaskRepositoryGetAllUserTasksCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTaskRepositoryGetAllUserTasksCall) Return(arg0 []task_model.Task, arg1 error) *MockTaskRepositoryGetAllUserTasksCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTaskRepositoryGetAllUserTasksCall) Do(f func(context.Context, uuid.UUID) ([]task_model.Task, error)) *MockTaskRepositoryGetAllUserTasksCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTaskRepositoryGetAllUserTasksCall) DoAndReturn(f func(context.Context, uuid.UUID) ([]task_model.Task, error)) *MockTaskRepositoryGetAllUserTasksCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetTaskByID mocks base method.
func (m *MockTaskRepository) GetTaskByID(ctx context.Context, id uuid.UUID) (task_model.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskByID", ctx, id)
	ret0, _ := ret[0].(task_model.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskByID indicates an expected call of GetTaskByID.
func (mr *MockTaskRepositoryMockRecorder) GetTaskByID(ctx, id any) *MockTaskRepositoryGetTaskByIDCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByID", reflect.TypeOf((*MockTaskRepository)(nil).GetTaskByID), ctx, id)
	return &MockTaskRepositoryGetTaskByIDCall{Call: call}
}

// MockTaskRepositoryGetTaskByIDCall wrap *gomock.Call
type MockTaskRepositoryGetTaskByIDCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTaskRepositoryGetTaskByIDCall) Return(arg0 task_model.Task, arg1 error) *MockTaskRepositoryGetTaskByIDCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTaskRepositoryGetTaskByIDCall) Do(f func(context.Context, uuid.UUID) (task_model.Task, error)) *MockTaskRepositoryGetTaskByIDCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTaskRepositoryGetTaskByIDCall) DoAndReturn(f func(context.Context, uuid.UUID) (task_model.Task, error)) *MockTaskRepositoryGetTaskByIDCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateTask mocks base method.
func (m *MockTaskRepository) UpdateTask(ctx context.Context, id uuid.UUID, title, description, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", ctx, id, title, description, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockTaskRepositoryMockRecorder) UpdateTask(ctx, id, title, description, status any) *MockTaskRepositoryUpdateTaskCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockTaskRepository)(nil).UpdateTask), ctx, id, title, description, status)
	return &MockTaskRepositoryUpdateTaskCall{Call: call}
}

// MockTaskRepositoryUpdateTaskCall wrap *gomock.Call
type MockTaskRepositoryUpdateTaskCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockTaskRepositoryUpdateTaskCall) Return(arg0 error) *MockTaskRepositoryUpdateTaskCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockTaskRepositoryUpdateTaskCall) Do(f func(context.Context, uuid.UUID, string, string, string) error) *MockTaskRepositoryUpdateTaskCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockTaskRepositoryUpdateTaskCall) DoAndReturn(f func(context.Context, uuid.UUID, string, string, string) error) *MockTaskRepositoryUpdateTaskCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
