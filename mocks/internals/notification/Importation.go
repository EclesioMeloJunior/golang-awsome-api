// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package notification

import mock "github.com/stretchr/testify/mock"

// Importation is an autogenerated mock type for the Importation type
type Importation struct {
	mock.Mock
}

// Notify provides a mock function with given fields:
func (_m *Importation) Notify() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NotifyFail provides a mock function with given fields: _a0
func (_m *Importation) NotifyFail(_a0 error) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(error) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NotifySuccess provides a mock function with given fields: _a0
func (_m *Importation) NotifySuccess(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
