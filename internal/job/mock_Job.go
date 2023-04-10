// Code generated by mockery v2.23.1. DO NOT EDIT.

package job

import mock "github.com/stretchr/testify/mock"

// MockJob is an autogenerated mock type for the Job type
type MockJob struct {
	mock.Mock
}

// Run provides a mock function with given fields:
func (_m *MockJob) Run() {
	_m.Called()
}

type mockConstructorTestingTNewMockJob interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockJob creates a new instance of MockJob. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockJob(t mockConstructorTestingTNewMockJob) *MockJob {
	mock := &MockJob{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}