// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	entity "recything/features/trash_exchange/entity"

	mock "github.com/stretchr/testify/mock"

	pagination "recything/utils/pagination"
)

// TrashExchangeRepositoryInterface is an autogenerated mock type for the TrashExchangeRepositoryInterface type
type TrashExchangeRepositoryInterface struct {
	mock.Mock
}

// CreateTrashExchange provides a mock function with given fields: data
func (_m *TrashExchangeRepositoryInterface) CreateTrashExchange(data entity.TrashExchangeCore) (entity.TrashExchangeCore, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for CreateTrashExchange")
	}

	var r0 entity.TrashExchangeCore
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.TrashExchangeCore) (entity.TrashExchangeCore, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(entity.TrashExchangeCore) entity.TrashExchangeCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(entity.TrashExchangeCore)
	}

	if rf, ok := ret.Get(1).(func(entity.TrashExchangeCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTrashExchangeDetails provides a mock function with given fields: data
func (_m *TrashExchangeRepositoryInterface) CreateTrashExchangeDetails(data entity.TrashExchangeDetailCore) (entity.TrashExchangeDetailCore, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for CreateTrashExchangeDetails")
	}

	var r0 entity.TrashExchangeDetailCore
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.TrashExchangeDetailCore) (entity.TrashExchangeDetailCore, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(entity.TrashExchangeDetailCore) entity.TrashExchangeDetailCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(entity.TrashExchangeDetailCore)
	}

	if rf, ok := ret.Get(1).(func(entity.TrashExchangeDetailCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTrashExchangeById provides a mock function with given fields: id
func (_m *TrashExchangeRepositoryInterface) DeleteTrashExchangeById(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTrashExchangeById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTrashExchange provides a mock function with given fields: page, limit, search
func (_m *TrashExchangeRepositoryInterface) GetAllTrashExchange(page int, limit int, search string) ([]entity.TrashExchangeCore, pagination.PageInfo, int, error) {
	ret := _m.Called(page, limit, search)

	if len(ret) == 0 {
		panic("no return value specified for GetAllTrashExchange")
	}

	var r0 []entity.TrashExchangeCore
	var r1 pagination.PageInfo
	var r2 int
	var r3 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]entity.TrashExchangeCore, pagination.PageInfo, int, error)); ok {
		return rf(page, limit, search)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []entity.TrashExchangeCore); ok {
		r0 = rf(page, limit, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.TrashExchangeCore)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string) pagination.PageInfo); ok {
		r1 = rf(page, limit, search)
	} else {
		r1 = ret.Get(1).(pagination.PageInfo)
	}

	if rf, ok := ret.Get(2).(func(int, int, string) int); ok {
		r2 = rf(page, limit, search)
	} else {
		r2 = ret.Get(2).(int)
	}

	if rf, ok := ret.Get(3).(func(int, int, string) error); ok {
		r3 = rf(page, limit, search)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetByEmail provides a mock function with given fields: email
func (_m *TrashExchangeRepositoryInterface) GetByEmail(email string) ([]map[string]interface{}, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 []map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]map[string]interface{}, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) []map[string]interface{}); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTrashExchangeById provides a mock function with given fields: id
func (_m *TrashExchangeRepositoryInterface) GetTrashExchangeById(id string) (entity.TrashExchangeCore, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetTrashExchangeById")
	}

	var r0 entity.TrashExchangeCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.TrashExchangeCore, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) entity.TrashExchangeCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.TrashExchangeCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTrashExchangeByIdTransaction provides a mock function with given fields: email, idTransaction
func (_m *TrashExchangeRepositoryInterface) GetTrashExchangeByIdTransaction(email string, idTransaction string) (map[string]interface{}, error) {
	ret := _m.Called(email, idTransaction)

	if len(ret) == 0 {
		panic("no return value specified for GetTrashExchangeByIdTransaction")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (map[string]interface{}, error)); ok {
		return rf(email, idTransaction)
	}
	if rf, ok := ret.Get(0).(func(string, string) map[string]interface{}); ok {
		r0 = rf(email, idTransaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, idTransaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTrashExchangeRepositoryInterface creates a new instance of TrashExchangeRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTrashExchangeRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *TrashExchangeRepositoryInterface {
	mock := &TrashExchangeRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
