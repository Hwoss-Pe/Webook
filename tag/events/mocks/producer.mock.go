// Code generated by MockGen. DO NOT EDIT.
// Source: ./producer.go

// Package evtmocks is a generated GoMock package.
package evtmocks

import (
	events "Webook/tag/events"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProducer is a mock of Producer interface.
type MockProducer struct {
	ctrl     *gomock.Controller
	recorder *MockProducerMockRecorder
}

// MockProducerMockRecorder is the mock recorder for MockProducer.
type MockProducerMockRecorder struct {
	mock *MockProducer
}

// NewMockProducer creates a new mock instance.
func NewMockProducer(ctrl *gomock.Controller) *MockProducer {
	mock := &MockProducer{ctrl: ctrl}
	mock.recorder = &MockProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducer) EXPECT() *MockProducerMockRecorder {
	return m.recorder
}

// ProduceSyncEvent mocks base method.
func (m *MockProducer) ProduceSyncEvent(ctx context.Context, data events.BizTags) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProduceSyncEvent", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProduceSyncEvent indicates an expected call of ProduceSyncEvent.
func (mr *MockProducerMockRecorder) ProduceSyncEvent(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProduceSyncEvent", reflect.TypeOf((*MockProducer)(nil).ProduceSyncEvent), ctx, data)
}