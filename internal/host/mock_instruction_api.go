// Code generated by MockGen. DO NOT EDIT.
// Source: instructionmanager.go

// Package host is a generated GoMock package.
package host

import (
	context "context"
	reflect "reflect"

	models "github.com/filanov/bm-inventory/models"
	gomock "github.com/golang/mock/gomock"
)

// MockInstructionApi is a mock of InstructionApi interface.
type MockInstructionApi struct {
	ctrl     *gomock.Controller
	recorder *MockInstructionApiMockRecorder
}

// MockInstructionApiMockRecorder is the mock recorder for MockInstructionApi.
type MockInstructionApiMockRecorder struct {
	mock *MockInstructionApi
}

// NewMockInstructionApi creates a new mock instance.
func NewMockInstructionApi(ctrl *gomock.Controller) *MockInstructionApi {
	mock := &MockInstructionApi{ctrl: ctrl}
	mock.recorder = &MockInstructionApiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInstructionApi) EXPECT() *MockInstructionApiMockRecorder {
	return m.recorder
}

// GetNextSteps mocks base method.
func (m *MockInstructionApi) GetNextSteps(ctx context.Context, host *models.Host) (models.Steps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextSteps", ctx, host)
	ret0, _ := ret[0].(models.Steps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextSteps indicates an expected call of GetNextSteps.
func (mr *MockInstructionApiMockRecorder) GetNextSteps(ctx, host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextSteps", reflect.TypeOf((*MockInstructionApi)(nil).GetNextSteps), ctx, host)
}
