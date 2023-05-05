// Code generated by mockery v2.23.1. DO NOT EDIT.

package controller

import (
	domain "hte-danger-zone-job/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockAlarmController is an autogenerated mock type for the AlarmController type
type MockAlarmController struct {
	mock.Mock
}

// Send provides a mock function with given fields: body
func (_m *MockAlarmController) Send(body *domain.SendAlarmReq) error {
	ret := _m.Called(body)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.SendAlarmReq) error); ok {
		r0 = rf(body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockAlarmController interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAlarmController creates a new instance of MockAlarmController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAlarmController(t mockConstructorTestingTNewMockAlarmController) *MockAlarmController {
	mock := &MockAlarmController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}