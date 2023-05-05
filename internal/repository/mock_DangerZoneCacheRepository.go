// Code generated by mockery v2.23.1. DO NOT EDIT.

package repository

import (
	domain "hte-danger-zone-job/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockDangerZoneCacheRepository is an autogenerated mock type for the DangerZoneCacheRepository type
type MockDangerZoneCacheRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: z
func (_m *MockDangerZoneCacheRepository) Create(z *domain.DangerZone) error {
	ret := _m.Called(z)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.DangerZone) error); ok {
		r0 = rf(z)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByDeviceID provides a mock function with given fields: deviceID
func (_m *MockDangerZoneCacheRepository) DeleteByDeviceID(deviceID string) error {
	ret := _m.Called(deviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(deviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByDeviceID provides a mock function with given fields: deviceID
func (_m *MockDangerZoneCacheRepository) GetByDeviceID(deviceID string) (*domain.DangerZone, error) {
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

type mockConstructorTestingTNewMockDangerZoneCacheRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockDangerZoneCacheRepository creates a new instance of MockDangerZoneCacheRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDangerZoneCacheRepository(t mockConstructorTestingTNewMockDangerZoneCacheRepository) *MockDangerZoneCacheRepository {
	mock := &MockDangerZoneCacheRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
