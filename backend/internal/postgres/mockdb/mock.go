// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/flexGURU/zeiba-glam/backend/internal/postgres/generated (interfaces: Querier)
//
// Generated by this command:
//
//	mockgen -package mockdb -destination ./internal/postgres/mockdb/mock.go github.com/flexGURU/zeiba-glam/backend/internal/postgres/generated Querier
//

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	generated "github.com/flexGURU/zeiba-glam/backend/internal/postgres/generated"
	gomock "go.uber.org/mock/gomock"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
	isgomock struct{}
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// GetUserByEmail mocks base method.
func (m *MockQuerier) GetUserByEmail(ctx context.Context, email string) (generated.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", ctx, email)
	ret0, _ := ret[0].(generated.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockQuerierMockRecorder) GetUserByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockQuerier)(nil).GetUserByEmail), ctx, email)
}
