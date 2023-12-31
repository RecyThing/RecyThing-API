// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	entity "recything/features/admin/entity"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"

	pagination "recything/utils/pagination"

	reportentity "recything/features/report/entity"

	userentity "recything/features/user/entity"
)

// AdminRepositoryInterface is an autogenerated mock type for the AdminRepositoryInterface type
type AdminRepositoryInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: image, data
func (_m *AdminRepositoryInterface) Create(image *multipart.FileHeader, data entity.AdminCore) (entity.AdminCore, error) {
	ret := _m.Called(image, data)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 entity.AdminCore
	var r1 error
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader, entity.AdminCore) (entity.AdminCore, error)); ok {
		return rf(image, data)
	}
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader, entity.AdminCore) entity.AdminCore); ok {
		r0 = rf(image, data)
	} else {
		r0 = ret.Get(0).(entity.AdminCore)
	}

	if rf, ok := ret.Get(1).(func(*multipart.FileHeader, entity.AdminCore) error); ok {
		r1 = rf(image, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: adminId
func (_m *AdminRepositoryInterface) Delete(adminId string) error {
	ret := _m.Called(adminId)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(adminId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUsers provides a mock function with given fields: adminId
func (_m *AdminRepositoryInterface) DeleteUsers(adminId string) error {
	ret := _m.Called(adminId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUsers")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(adminId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByEmail provides a mock function with given fields: email
func (_m *AdminRepositoryInterface) FindByEmail(email string) error {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for FindByEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByEmailANDPassword provides a mock function with given fields: data
func (_m *AdminRepositoryInterface) FindByEmailANDPassword(data entity.AdminCore) (entity.AdminCore, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for FindByEmailANDPassword")
	}

	var r0 entity.AdminCore
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.AdminCore) (entity.AdminCore, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(entity.AdminCore) entity.AdminCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(entity.AdminCore)
	}

	if rf, ok := ret.Get(1).(func(entity.AdminCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllReport provides a mock function with given fields: status, search, page, limit
func (_m *AdminRepositoryInterface) GetAllReport(status string, search string, page int, limit int) ([]reportentity.ReportCore, pagination.PageInfo, pagination.CountDataInfo, error) {
	ret := _m.Called(status, search, page, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetAllReport")
	}

	var r0 []reportentity.ReportCore
	var r1 pagination.PageInfo
	var r2 pagination.CountDataInfo
	var r3 error
	if rf, ok := ret.Get(0).(func(string, string, int, int) ([]reportentity.ReportCore, pagination.PageInfo, pagination.CountDataInfo, error)); ok {
		return rf(status, search, page, limit)
	}
	if rf, ok := ret.Get(0).(func(string, string, int, int) []reportentity.ReportCore); ok {
		r0 = rf(status, search, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]reportentity.ReportCore)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, int, int) pagination.PageInfo); ok {
		r1 = rf(status, search, page, limit)
	} else {
		r1 = ret.Get(1).(pagination.PageInfo)
	}

	if rf, ok := ret.Get(2).(func(string, string, int, int) pagination.CountDataInfo); ok {
		r2 = rf(status, search, page, limit)
	} else {
		r2 = ret.Get(2).(pagination.CountDataInfo)
	}

	if rf, ok := ret.Get(3).(func(string, string, int, int) error); ok {
		r3 = rf(status, search, page, limit)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetAllUsers provides a mock function with given fields: search, page, limit
func (_m *AdminRepositoryInterface) GetAllUsers(search string, page int, limit int) ([]userentity.UsersCore, pagination.PageInfo, int, error) {
	ret := _m.Called(search, page, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsers")
	}

	var r0 []userentity.UsersCore
	var r1 pagination.PageInfo
	var r2 int
	var r3 error
	if rf, ok := ret.Get(0).(func(string, int, int) ([]userentity.UsersCore, pagination.PageInfo, int, error)); ok {
		return rf(search, page, limit)
	}
	if rf, ok := ret.Get(0).(func(string, int, int) []userentity.UsersCore); ok {
		r0 = rf(search, page, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]userentity.UsersCore)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int, int) pagination.PageInfo); ok {
		r1 = rf(search, page, limit)
	} else {
		r1 = ret.Get(1).(pagination.PageInfo)
	}

	if rf, ok := ret.Get(2).(func(string, int, int) int); ok {
		r2 = rf(search, page, limit)
	} else {
		r2 = ret.Get(2).(int)
	}

	if rf, ok := ret.Get(3).(func(string, int, int) error); ok {
		r3 = rf(search, page, limit)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// GetByIdUser provides a mock function with given fields: userId
func (_m *AdminRepositoryInterface) GetByIdUser(userId string) (userentity.UsersCore, error) {
	ret := _m.Called(userId)

	if len(ret) == 0 {
		panic("no return value specified for GetByIdUser")
	}

	var r0 userentity.UsersCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (userentity.UsersCore, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) userentity.UsersCore); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(userentity.UsersCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCount provides a mock function with given fields: fullName, role
func (_m *AdminRepositoryInterface) GetCount(fullName string, role string) (int, error) {
	ret := _m.Called(fullName, role)

	if len(ret) == 0 {
		panic("no return value specified for GetCount")
	}

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (int, error)); ok {
		return rf(fullName, role)
	}
	if rf, ok := ret.Get(0).(func(string, string) int); ok {
		r0 = rf(fullName, role)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(fullName, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCountByStatus provides a mock function with given fields: status, search
func (_m *AdminRepositoryInterface) GetCountByStatus(status string, search string) (int64, error) {
	ret := _m.Called(status, search)

	if len(ret) == 0 {
		panic("no return value specified for GetCountByStatus")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (int64, error)); ok {
		return rf(status, search)
	}
	if rf, ok := ret.Get(0).(func(string, string) int64); ok {
		r0 = rf(status, search)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(status, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReportById provides a mock function with given fields: id
func (_m *AdminRepositoryInterface) GetReportById(id string) (reportentity.ReportCore, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetReportById")
	}

	var r0 reportentity.ReportCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (reportentity.ReportCore, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) reportentity.ReportCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(reportentity.ReportCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAll provides a mock function with given fields: page, limit, search
func (_m *AdminRepositoryInterface) SelectAll(page int, limit int, search string) ([]entity.AdminCore, pagination.PageInfo, int, error) {
	ret := _m.Called(page, limit, search)

	if len(ret) == 0 {
		panic("no return value specified for SelectAll")
	}

	var r0 []entity.AdminCore
	var r1 pagination.PageInfo
	var r2 int
	var r3 error
	if rf, ok := ret.Get(0).(func(int, int, string) ([]entity.AdminCore, pagination.PageInfo, int, error)); ok {
		return rf(page, limit, search)
	}
	if rf, ok := ret.Get(0).(func(int, int, string) []entity.AdminCore); ok {
		r0 = rf(page, limit, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.AdminCore)
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

// SelectById provides a mock function with given fields: adminId
func (_m *AdminRepositoryInterface) SelectById(adminId string) (entity.AdminCore, error) {
	ret := _m.Called(adminId)

	if len(ret) == 0 {
		panic("no return value specified for SelectById")
	}

	var r0 entity.AdminCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.AdminCore, error)); ok {
		return rf(adminId)
	}
	if rf, ok := ret.Get(0).(func(string) entity.AdminCore); ok {
		r0 = rf(adminId)
	} else {
		r0 = ret.Get(0).(entity.AdminCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(adminId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: image, adminId, data
func (_m *AdminRepositoryInterface) Update(image *multipart.FileHeader, adminId string, data entity.AdminCore) error {
	ret := _m.Called(image, adminId, data)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*multipart.FileHeader, string, entity.AdminCore) error); ok {
		r0 = rf(image, adminId, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatusReport provides a mock function with given fields: id, status, reason
func (_m *AdminRepositoryInterface) UpdateStatusReport(id string, status string, reason string) (reportentity.ReportCore, error) {
	ret := _m.Called(id, status, reason)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatusReport")
	}

	var r0 reportentity.ReportCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (reportentity.ReportCore, error)); ok {
		return rf(id, status, reason)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) reportentity.ReportCore); ok {
		r0 = rf(id, status, reason)
	} else {
		r0 = ret.Get(0).(reportentity.ReportCore)
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(id, status, reason)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAdminRepositoryInterface creates a new instance of AdminRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAdminRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AdminRepositoryInterface {
	mock := &AdminRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
