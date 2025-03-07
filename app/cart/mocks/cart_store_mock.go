// Code generated by MockGen. DO NOT EDIT.
// Source: D:/GoCode/projects/douyin-mall-microservice/app/cart/biz/model/cart.go
//
// Generated by this command:
//
//	mockgen.exe -source=D:/GoCode/projects/douyin-mall-microservice/app/cart/biz/model/cart.go -destination=D:/GoCode/projects/douyin-mall-microservice/app/cart/mocks/cart_store_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	model "github.com/All-Done-Right/douyin-mall-microservice/app/cart/biz/model"
	gomock "go.uber.org/mock/gomock"
)

// MockCartStore is a mock of CartStore interface.
type MockCartStore struct {
	ctrl     *gomock.Controller
	recorder *MockCartStoreMockRecorder
}

// MockCartStoreMockRecorder is the mock recorder for MockCartStore.
type MockCartStoreMockRecorder struct {
	mock *MockCartStore
}

// NewMockCartStore creates a new mock instance.
func NewMockCartStore(ctrl *gomock.Controller) *MockCartStore {
	mock := &MockCartStore{ctrl: ctrl}
	mock.recorder = &MockCartStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCartStore) EXPECT() *MockCartStoreMockRecorder {
	return m.recorder
}

// AddItem mocks base method.
func (m *MockCartStore) AddItem(ctx context.Context, db *sql.DB, item *model.Cart) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddItem", ctx, db, item)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddItem indicates an expected call of AddItem.
func (mr *MockCartStoreMockRecorder) AddItem(ctx, db, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddItem", reflect.TypeOf((*MockCartStore)(nil).AddItem), ctx, db, item)
}

// EmptyCart mocks base method.
func (m *MockCartStore) EmptyCart(ctx context.Context, db *sql.DB, userID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EmptyCart", ctx, db, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EmptyCart indicates an expected call of EmptyCart.
func (mr *MockCartStoreMockRecorder) EmptyCart(ctx, db, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmptyCart", reflect.TypeOf((*MockCartStore)(nil).EmptyCart), ctx, db, userID)
}

// GetCartByUserId mocks base method.
func (m *MockCartStore) GetCartByUserId(ctx context.Context, db *sql.DB, userID int64) ([]*model.Cart, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCartByUserId", ctx, db, userID)
	ret0, _ := ret[0].([]*model.Cart)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCartByUserId indicates an expected call of GetCartByUserId.
func (mr *MockCartStoreMockRecorder) GetCartByUserId(ctx, db, userID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCartByUserId", reflect.TypeOf((*MockCartStore)(nil).GetCartByUserId), ctx, db, userID)
}
