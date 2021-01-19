// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package services

import (
	models "go-challenge/internals/models"

	mock "github.com/stretchr/testify/mock"
)

// Product is an autogenerated mock type for the Product type
type Product struct {
	mock.Mock
}

// DeleteProductByID provides a mock function with given fields: _a0
func (_m *Product) DeleteProductByID(_a0 string) (*models.Product, error) {
	ret := _m.Called(_a0)

	var r0 *models.Product
	if rf, ok := ret.Get(0).(func(string) *models.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductByID provides a mock function with given fields: _a0
func (_m *Product) GetProductByID(_a0 string) (*models.Product, error) {
	ret := _m.Called(_a0)

	var r0 *models.Product
	if rf, ok := ret.Get(0).(func(string) *models.Product); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProducts provides a mock function with given fields: filter, page, size
func (_m *Product) GetProducts(filter interface{}, page int, size int) ([]models.Product, error) {
	ret := _m.Called(filter, page, size)

	var r0 []models.Product
	if rf, ok := ret.Get(0).(func(interface{}, int, int) []models.Product); ok {
		r0 = rf(filter, page, size)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, int, int) error); ok {
		r1 = rf(filter, page, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProductByID provides a mock function with given fields: _a0, _a1
func (_m *Product) UpdateProductByID(_a0 string, _a1 *models.Product) (*models.Product, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *models.Product
	if rf, ok := ret.Get(0).(func(string, *models.Product) *models.Product); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *models.Product) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}