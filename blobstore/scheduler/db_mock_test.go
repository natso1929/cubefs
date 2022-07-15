// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubefs/cubefs/blobstore/scheduler/db (interfaces: IKafkaOffsetTable,IOrphanShardTable,IInspectCheckPointTable)

// Package scheduler is a generated GoMock package.
package scheduler

import (
	context "context"
	reflect "reflect"

	proto "github.com/cubefs/cubefs/blobstore/common/proto"
	db "github.com/cubefs/cubefs/blobstore/scheduler/db"
	gomock "github.com/golang/mock/gomock"
)

// MockKafkaOffsetTable is a mock of IKafkaOffsetTable interface.
type MockKafkaOffsetTable struct {
	ctrl     *gomock.Controller
	recorder *MockKafkaOffsetTableMockRecorder
}

// MockKafkaOffsetTableMockRecorder is the mock recorder for MockKafkaOffsetTable.
type MockKafkaOffsetTableMockRecorder struct {
	mock *MockKafkaOffsetTable
}

// NewMockKafkaOffsetTable creates a new mock instance.
func NewMockKafkaOffsetTable(ctrl *gomock.Controller) *MockKafkaOffsetTable {
	mock := &MockKafkaOffsetTable{ctrl: ctrl}
	mock.recorder = &MockKafkaOffsetTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKafkaOffsetTable) EXPECT() *MockKafkaOffsetTableMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockKafkaOffsetTable) Get(arg0 string, arg1 int32) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKafkaOffsetTableMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKafkaOffsetTable)(nil).Get), arg0, arg1)
}

// Set mocks base method.
func (m *MockKafkaOffsetTable) Set(arg0 string, arg1 int32, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockKafkaOffsetTableMockRecorder) Set(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockKafkaOffsetTable)(nil).Set), arg0, arg1, arg2)
}

// MockOrphanShardTable is a mock of IOrphanShardTable interface.
type MockOrphanShardTable struct {
	ctrl     *gomock.Controller
	recorder *MockOrphanShardTableMockRecorder
}

// MockOrphanShardTableMockRecorder is the mock recorder for MockOrphanShardTable.
type MockOrphanShardTableMockRecorder struct {
	mock *MockOrphanShardTable
}

// NewMockOrphanShardTable creates a new mock instance.
func NewMockOrphanShardTable(ctrl *gomock.Controller) *MockOrphanShardTable {
	mock := &MockOrphanShardTable{ctrl: ctrl}
	mock.recorder = &MockOrphanShardTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrphanShardTable) EXPECT() *MockOrphanShardTableMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockOrphanShardTable) Save(arg0 db.OrphanShard) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockOrphanShardTableMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockOrphanShardTable)(nil).Save), arg0)
}

// MockInspectCheckPointTable is a mock of IInspectCheckPointTable interface.
type MockInspectCheckPointTable struct {
	ctrl     *gomock.Controller
	recorder *MockInspectCheckPointTableMockRecorder
}

// MockInspectCheckPointTableMockRecorder is the mock recorder for MockInspectCheckPointTable.
type MockInspectCheckPointTableMockRecorder struct {
	mock *MockInspectCheckPointTable
}

// NewMockInspectCheckPointTable creates a new mock instance.
func NewMockInspectCheckPointTable(ctrl *gomock.Controller) *MockInspectCheckPointTable {
	mock := &MockInspectCheckPointTable{ctrl: ctrl}
	mock.recorder = &MockInspectCheckPointTableMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInspectCheckPointTable) EXPECT() *MockInspectCheckPointTableMockRecorder {
	return m.recorder
}

// GetCheckPoint mocks base method.
func (m *MockInspectCheckPointTable) GetCheckPoint(arg0 context.Context) (*proto.InspectCheckPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCheckPoint", arg0)
	ret0, _ := ret[0].(*proto.InspectCheckPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCheckPoint indicates an expected call of GetCheckPoint.
func (mr *MockInspectCheckPointTableMockRecorder) GetCheckPoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCheckPoint", reflect.TypeOf((*MockInspectCheckPointTable)(nil).GetCheckPoint), arg0)
}

// SaveCheckPoint mocks base method.
func (m *MockInspectCheckPointTable) SaveCheckPoint(arg0 context.Context, arg1 proto.Vid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCheckPoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCheckPoint indicates an expected call of SaveCheckPoint.
func (mr *MockInspectCheckPointTableMockRecorder) SaveCheckPoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCheckPoint", reflect.TypeOf((*MockInspectCheckPointTable)(nil).SaveCheckPoint), arg0, arg1)
}