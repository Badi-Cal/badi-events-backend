// Code generated by MockGen. DO NOT EDIT.
// Source: models/notifications.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	models "badi-cal/badi-events-backend/models"
	orm "badi-cal/badi-events-backend/orm"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockINotifications is a mock of INotifications interface.
type MockINotifications struct {
	ctrl     *gomock.Controller
	recorder *MockINotificationsMockRecorder
}

// MockINotificationsMockRecorder is the mock recorder for MockINotifications.
type MockINotificationsMockRecorder struct {
	mock *MockINotifications
}

// NewMockINotifications creates a new mock instance.
func NewMockINotifications(ctrl *gomock.Controller) *MockINotifications {
	mock := &MockINotifications{ctrl: ctrl}
	mock.recorder = &MockINotificationsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockINotifications) EXPECT() *MockINotificationsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockINotifications) Create(payload models.NotificationCreatePayload) (*orm.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", payload)
	ret0, _ := ret[0].(*orm.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockINotificationsMockRecorder) Create(payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockINotifications)(nil).Create), payload)
}

// GetAll mocks base method.
func (m *MockINotifications) GetAll() ([]*orm.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*orm.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockINotificationsMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockINotifications)(nil).GetAll))
}