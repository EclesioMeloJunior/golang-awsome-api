// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package repository

import (
	models "go-challenge/internals/models"

	mock "github.com/stretchr/testify/mock"
)

// Import is an autogenerated mock type for the Import type
type Import struct {
	mock.Mock
}

// ExecuteImport provides a mock function with given fields: _a0, _a1
func (_m *Import) ExecuteImport(_a0 *models.Import, _a1 []interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Import, []interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllImports provides a mock function with given fields:
func (_m *Import) GetAllImports() ([]models.Import, error) {
	ret := _m.Called()

	var r0 []models.Import
	if rf, ok := ret.Get(0).(func() []models.Import); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Import)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastImport provides a mock function with given fields:
func (_m *Import) GetLastImport() (*models.Import, error) {
	ret := _m.Called()

	var r0 *models.Import
	if rf, ok := ret.Get(0).(func() *models.Import); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Import)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}