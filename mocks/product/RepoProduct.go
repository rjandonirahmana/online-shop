// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	product "graphql/product"

	mock "github.com/stretchr/testify/mock"
)

// RepoProduct is an autogenerated mock type for the RepoProduct type
type RepoProduct struct {
	mock.Mock
}

// AddProduct provides a mock function with given fields: _a0
func (_m *RepoProduct) AddProduct(_a0 product.Product) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(product.Product) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateCart provides a mock function with given fields: customerID, id
func (_m *RepoProduct) CreateCart(customerID int, id int) error {
	ret := _m.Called(customerID, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(customerID, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDetailsProductByID provides a mock function with given fields: id
func (_m *RepoProduct) GetDetailsProductByID(id int) []product.Product {
	ret := _m.Called(id)

	var r0 []product.Product
	if rf, ok := ret.Get(0).(func(int) []product.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Product)
		}
	}

	return r0
}

// GetLastID provides a mock function with given fields:
func (_m *RepoProduct) GetLastID() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListCartByID provides a mock function with given fields: cartid, customerid
func (_m *RepoProduct) GetListCartByID(cartid int, customerid int) ([]product.Product, error) {
	ret := _m.Called(cartid, customerid)

	var r0 []product.Product
	if rf, ok := ret.Get(0).(func(int, int) []product.Product); ok {
		r0 = rf(cartid, customerid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(cartid, customerid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductByCategoryName provides a mock function with given fields: name_category
func (_m *RepoProduct) GetProductByCategoryName(name_category string) ([]product.Product, error) {
	ret := _m.Called(name_category)

	var r0 []product.Product
	if rf, ok := ret.Get(0).(func(string) []product.Product); ok {
		r0 = rf(name_category)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name_category)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductByID provides a mock function with given fields: id
func (_m *RepoProduct) GetProductByID(id int) (product.Product, error) {
	ret := _m.Called(id)

	var r0 product.Product
	if rf, ok := ret.Get(0).(func(int) product.Product); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(product.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductsOrderByprice provides a mock function with given fields: name
func (_m *RepoProduct) GetProductsOrderByprice(name string) ([]product.Product, error) {
	ret := _m.Called(name)

	var r0 []product.Product
	if rf, ok := ret.Get(0).(func(string) []product.Product); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductsOrderBypriceDSC provides a mock function with given fields: name
func (_m *RepoProduct) GetProductsOrderBypriceDSC(name string) ([]product.Product, error) {
	ret := _m.Called(name)

	var r0 []product.Product
	if rf, ok := ret.Get(0).(func(string) []product.Product); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertDetailProduct provides a mock function with given fields: _a0
func (_m *RepoProduct) InsertDetailProduct(_a0 product.ProductDesc) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(product.ProductDesc) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertShoppingCart provides a mock function with given fields: cartid, productid, price, name
func (_m *RepoProduct) InsertShoppingCart(cartid int, productid int, price int32, name string) error {
	ret := _m.Called(cartid, productid, price, name)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, int32, string) error); ok {
		r0 = rf(cartid, productid, price, name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePrice provides a mock function with given fields: id, price
func (_m *RepoProduct) UpdatePrice(id int, price int) error {
	ret := _m.Called(id, price)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(id, price)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
