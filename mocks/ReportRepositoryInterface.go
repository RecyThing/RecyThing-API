// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	entity "recything/features/report/entity"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"
)

// ReportRepositoryInterface is an autogenerated mock type for the ReportRepositoryInterface type
type ReportRepositoryInterface struct {
	mock.Mock
}

// Insert provides a mock function with given fields: reportInput, images
func (_m *ReportRepositoryInterface) Insert(reportInput entity.ReportCore, images []*multipart.FileHeader) (entity.ReportCore, error) {
	ret := _m.Called(reportInput, images)

	if len(ret) == 0 {
		panic("no return value specified for Insert")
	}

	var r0 entity.ReportCore
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.ReportCore, []*multipart.FileHeader) (entity.ReportCore, error)); ok {
		return rf(reportInput, images)
	}
	if rf, ok := ret.Get(0).(func(entity.ReportCore, []*multipart.FileHeader) entity.ReportCore); ok {
		r0 = rf(reportInput, images)
	} else {
		r0 = ret.Get(0).(entity.ReportCore)
	}

	if rf, ok := ret.Get(1).(func(entity.ReportCore, []*multipart.FileHeader) error); ok {
		r1 = rf(reportInput, images)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAllReport provides a mock function with given fields: idUser
func (_m *ReportRepositoryInterface) ReadAllReport(idUser string) ([]entity.ReportCore, error) {
	ret := _m.Called(idUser)

	if len(ret) == 0 {
		panic("no return value specified for ReadAllReport")
	}

	var r0 []entity.ReportCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]entity.ReportCore, error)); ok {
		return rf(idUser)
	}
	if rf, ok := ret.Get(0).(func(string) []entity.ReportCore); ok {
		r0 = rf(idUser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.ReportCore)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(idUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectById provides a mock function with given fields: idReport
func (_m *ReportRepositoryInterface) SelectById(idReport string) (entity.ReportCore, error) {
	ret := _m.Called(idReport)

	if len(ret) == 0 {
		panic("no return value specified for SelectById")
	}

	var r0 entity.ReportCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.ReportCore, error)); ok {
		return rf(idReport)
	}
	if rf, ok := ret.Get(0).(func(string) entity.ReportCore); ok {
		r0 = rf(idReport)
	} else {
		r0 = ret.Get(0).(entity.ReportCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(idReport)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewReportRepositoryInterface creates a new instance of ReportRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewReportRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ReportRepositoryInterface {
	mock := &ReportRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
