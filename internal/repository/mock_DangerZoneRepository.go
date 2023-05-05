// Code generated by mockery v2.23.1. DO NOT EDIT.

package repository

import (
	domain "hte-danger-zone-job/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// MockDangerZoneRepository is an autogenerated mock type for the DangerZoneRepository type
type MockDangerZoneRepository struct {
	mock.Mock
}

// DeleteByDeviceID provides a mock function with given fields: deviceID
func (_m *MockDangerZoneRepository) DeleteByDeviceID(deviceID string) error {
	ret := _m.Called(deviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(deviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllActive provides a mock function with given fields:
func (_m *MockDangerZoneRepository) GetAllActive() (*[]domain.DangerZone, error) {
	ret := _m.Called()

	var r0 *[]domain.DangerZone
	var r1 error
	if rf, ok := ret.Get(0).(func() (*[]domain.DangerZone, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *[]domain.DangerZone); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.DangerZone)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockDangerZoneRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockDangerZoneRepository creates a new instance of MockDangerZoneRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockDangerZoneRepository(t mockConstructorTestingTNewMockDangerZoneRepository) *MockDangerZoneRepository {
	mock := &MockDangerZoneRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
