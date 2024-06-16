// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	decimal "github.com/shopspring/decimal"
	mock "github.com/stretchr/testify/mock"
)

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// GetRate provides a mock function with given fields: base, target
func (_m *Provider) GetRate(base string, target string) (decimal.Decimal, error) {
	ret := _m.Called(base, target)

	if len(ret) == 0 {
		panic("no return value specified for GetRate")
	}

	var r0 decimal.Decimal
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (decimal.Decimal, error)); ok {
		return rf(base, target)
	}
	if rf, ok := ret.Get(0).(func(string, string) decimal.Decimal); ok {
		r0 = rf(base, target)
	} else {
		r0 = ret.Get(0).(decimal.Decimal)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(base, target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewProvider creates a new instance of Provider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *Provider {
	mock := &Provider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
