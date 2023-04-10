// Code generated by mockery v2.23.1. DO NOT EDIT.

package service

import (
	domain "hte-danger-zone-job/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockDangerZoneCacheService is an autogenerated mock type for the DangerZoneCacheService type
type MockDangerZoneCacheService struct {
	mock.Mock
}

// Create provides a mock function with given fields: z
func (_m *MockDangerZoneCacheService) Create(z *domain.DangerZone) error {
	ret := _m.Called(z)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.DangerZone) error); ok {
		r0 = rf(z)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByDeviceID provides a mock function with given fields: deviceID
func (_m *MockDangerZoneCacheService) GetByDeviceID(deviceID string) (*domain.DangerZone, error) {
	ret := _m.Called(deviceID)

	var r0 *domain.DangerZone
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.DangerZone, error)); ok {
		return rf(deviceID)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.DangerZone); ok {
		r0 = rf(deviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.DangerZone)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockDangerZoneCacheService interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockDangerZoneCacheService creates a new instance of MockDangerZoneCacheService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDangerZoneCacheService(t mockConstructorTestingTNewMockDangerZoneCacheService) *MockDangerZoneCacheService {
	mock := &MockDangerZoneCacheService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}