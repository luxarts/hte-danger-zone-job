// Code generated by mockery v2.23.1. DO NOT EDIT.

package repository

import mock "github.com/stretchr/testify/mock"

// MockAlarmRepository is an autogenerated mock type for the AlarmRepository type
type MockAlarmRepository struct {
	mock.Mock
}

// Send provides a mock function with given fields: deviceID, message
func (_m *MockAlarmRepository) Send(deviceID string, message string) error {
	ret := _m.Called(deviceID, message)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(deviceID, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockAlarmRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAlarmRepository creates a new instance of MockAlarmRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAlarmRepository(t mockConstructorTestingTNewMockAlarmRepository) *MockAlarmRepository {
	mock := &MockAlarmRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
