// Code generated by mockery v2.23.1. DO NOT EDIT.

package controller

import mock "github.com/stretchr/testify/mock"

// MockGeolocController is an autogenerated mock type for the GeolocController type
type MockGeolocController struct {
	mock.Mock
}

// Process provides a mock function with given fields: deviceID, body
func (_m *MockGeolocController) Process(deviceID string, body string) error {
	ret := _m.Called(deviceID, body)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(deviceID, body)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockGeolocController interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockGeolocController creates a new instance of MockGeolocController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockGeolocController(t mockConstructorTestingTNewMockGeolocController) *MockGeolocController {
	mock := &MockGeolocController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
