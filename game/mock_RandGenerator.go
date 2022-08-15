// Code generated by mockery v2.14.0. DO NOT EDIT.

package game

import mock "github.com/stretchr/testify/mock"

// MockRandGenerator is an autogenerated mock type for the RandGenerator type
type MockRandGenerator struct {
	mock.Mock
}

// Rand provides a mock function with given fields: end
func (_m *MockRandGenerator) Rand(end int) int {
	ret := _m.Called(end)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(end)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewMockRandGenerator interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRandGenerator creates a new instance of MockRandGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRandGenerator(t mockConstructorTestingTNewMockRandGenerator) *MockRandGenerator {
	mock := &MockRandGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}