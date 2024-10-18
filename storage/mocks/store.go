// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/onflow/flow-emulator/storage (interfaces: Store)
//
// Generated by this command:
//
//	mockgen -destination=storage/mocks/store.go -package=mocks github.com/onflow/flow-emulator/storage Store
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "github.com/onflow/flow-emulator/types"
	snapshot "github.com/onflow/flow-go/fvm/storage/snapshot"
	flow "github.com/onflow/flow-go/model/flow"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
	isgomock struct{}
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// BlockByHeight mocks base method.
func (m *MockStore) BlockByHeight(ctx context.Context, height uint64) (*flow.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHeight", ctx, height)
	ret0, _ := ret[0].(*flow.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHeight indicates an expected call of BlockByHeight.
func (mr *MockStoreMockRecorder) BlockByHeight(ctx, height any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHeight", reflect.TypeOf((*MockStore)(nil).BlockByHeight), ctx, height)
}

// BlockByID mocks base method.
func (m *MockStore) BlockByID(ctx context.Context, blockID flow.Identifier) (*flow.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByID", ctx, blockID)
	ret0, _ := ret[0].(*flow.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByID indicates an expected call of BlockByID.
func (mr *MockStoreMockRecorder) BlockByID(ctx, blockID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByID", reflect.TypeOf((*MockStore)(nil).BlockByID), ctx, blockID)
}

// CollectionByID mocks base method.
func (m *MockStore) CollectionByID(ctx context.Context, collectionID flow.Identifier) (flow.LightCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CollectionByID", ctx, collectionID)
	ret0, _ := ret[0].(flow.LightCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CollectionByID indicates an expected call of CollectionByID.
func (mr *MockStoreMockRecorder) CollectionByID(ctx, collectionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollectionByID", reflect.TypeOf((*MockStore)(nil).CollectionByID), ctx, collectionID)
}

// CommitBlock mocks base method.
func (m *MockStore) CommitBlock(ctx context.Context, block flow.Block, collections []*flow.LightCollection, transactions map[flow.Identifier]*flow.TransactionBody, transactionResults map[flow.Identifier]*types.StorableTransactionResult, executionSnapshot *snapshot.ExecutionSnapshot, events []flow.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitBlock", ctx, block, collections, transactions, transactionResults, executionSnapshot, events)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitBlock indicates an expected call of CommitBlock.
func (mr *MockStoreMockRecorder) CommitBlock(ctx, block, collections, transactions, transactionResults, executionSnapshot, events any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitBlock", reflect.TypeOf((*MockStore)(nil).CommitBlock), ctx, block, collections, transactions, transactionResults, executionSnapshot, events)
}

// EventsByHeight mocks base method.
func (m *MockStore) EventsByHeight(ctx context.Context, blockHeight uint64, eventType string) ([]flow.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventsByHeight", ctx, blockHeight, eventType)
	ret0, _ := ret[0].([]flow.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EventsByHeight indicates an expected call of EventsByHeight.
func (mr *MockStoreMockRecorder) EventsByHeight(ctx, blockHeight, eventType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsByHeight", reflect.TypeOf((*MockStore)(nil).EventsByHeight), ctx, blockHeight, eventType)
}

// FullCollectionByID mocks base method.
func (m *MockStore) FullCollectionByID(ctx context.Context, collectionID flow.Identifier) (flow.Collection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullCollectionByID", ctx, collectionID)
	ret0, _ := ret[0].(flow.Collection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FullCollectionByID indicates an expected call of FullCollectionByID.
func (mr *MockStoreMockRecorder) FullCollectionByID(ctx, collectionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullCollectionByID", reflect.TypeOf((*MockStore)(nil).FullCollectionByID), ctx, collectionID)
}

// LatestBlock mocks base method.
func (m *MockStore) LatestBlock(ctx context.Context) (flow.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestBlock", ctx)
	ret0, _ := ret[0].(flow.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestBlock indicates an expected call of LatestBlock.
func (mr *MockStoreMockRecorder) LatestBlock(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestBlock", reflect.TypeOf((*MockStore)(nil).LatestBlock), ctx)
}

// LatestBlockHeight mocks base method.
func (m *MockStore) LatestBlockHeight(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestBlockHeight", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestBlockHeight indicates an expected call of LatestBlockHeight.
func (mr *MockStoreMockRecorder) LatestBlockHeight(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestBlockHeight", reflect.TypeOf((*MockStore)(nil).LatestBlockHeight), ctx)
}

// LedgerByHeight mocks base method.
func (m *MockStore) LedgerByHeight(ctx context.Context, blockHeight uint64) (snapshot.StorageSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LedgerByHeight", ctx, blockHeight)
	ret0, _ := ret[0].(snapshot.StorageSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LedgerByHeight indicates an expected call of LedgerByHeight.
func (mr *MockStoreMockRecorder) LedgerByHeight(ctx, blockHeight any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LedgerByHeight", reflect.TypeOf((*MockStore)(nil).LedgerByHeight), ctx, blockHeight)
}

// Start mocks base method.
func (m *MockStore) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockStoreMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockStore)(nil).Start))
}

// Stop mocks base method.
func (m *MockStore) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockStoreMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockStore)(nil).Stop))
}

// StoreBlock mocks base method.
func (m *MockStore) StoreBlock(ctx context.Context, block *flow.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreBlock", ctx, block)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreBlock indicates an expected call of StoreBlock.
func (mr *MockStoreMockRecorder) StoreBlock(ctx, block any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreBlock", reflect.TypeOf((*MockStore)(nil).StoreBlock), ctx, block)
}

// TransactionByID mocks base method.
func (m *MockStore) TransactionByID(ctx context.Context, transactionID flow.Identifier) (flow.TransactionBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByID", ctx, transactionID)
	ret0, _ := ret[0].(flow.TransactionBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionByID indicates an expected call of TransactionByID.
func (mr *MockStoreMockRecorder) TransactionByID(ctx, transactionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByID", reflect.TypeOf((*MockStore)(nil).TransactionByID), ctx, transactionID)
}

// TransactionResultByID mocks base method.
func (m *MockStore) TransactionResultByID(ctx context.Context, transactionID flow.Identifier) (types.StorableTransactionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionResultByID", ctx, transactionID)
	ret0, _ := ret[0].(types.StorableTransactionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionResultByID indicates an expected call of TransactionResultByID.
func (mr *MockStoreMockRecorder) TransactionResultByID(ctx, transactionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionResultByID", reflect.TypeOf((*MockStore)(nil).TransactionResultByID), ctx, transactionID)
}
