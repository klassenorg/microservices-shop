// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCart is a mock of Cart interface.
type MockCart struct {
	ctrl     *gomock.Controller
	recorder *MockCartMockRecorder
}

// MockCartMockRecorder is the mock recorder for MockCart.
type MockCartMockRecorder struct {
	mock *MockCart
}

// NewMockCart creates a new mock instance.
func NewMockCart(ctrl *gomock.Controller) *MockCart {
	mock := &MockCart{ctrl: ctrl}
	mock.recorder = &MockCartMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCart) EXPECT() *MockCartMockRecorder {
	return m.recorder
}

// AddByID mocks base method.
func (m *MockCart) AddByID(ctx context.Context, id, product string, count int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddByID", ctx, id, product, count)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddByID indicates an expected call of AddByID.
func (mr *MockCartMockRecorder) AddByID(ctx, id, product, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddByID", reflect.TypeOf((*MockCart)(nil).AddByID), ctx, id, product, count)
}

// GetByID mocks base method.
func (m *MockCart) GetByID(ctx context.Context, id string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockCartMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockCart)(nil).GetByID), ctx, id)
}

// RemoveAllByID mocks base method.
func (m *MockCart) RemoveAllByID(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllByID", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAllByID indicates an expected call of RemoveAllByID.
func (mr *MockCartMockRecorder) RemoveAllByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllByID", reflect.TypeOf((*MockCart)(nil).RemoveAllByID), ctx, id)
}

// RemoveByID mocks base method.
func (m *MockCart) RemoveByID(ctx context.Context, id, product string, count int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveByID", ctx, id, product, count)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveByID indicates an expected call of RemoveByID.
func (mr *MockCartMockRecorder) RemoveByID(ctx, id, product, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveByID", reflect.TypeOf((*MockCart)(nil).RemoveByID), ctx, id, product, count)
}