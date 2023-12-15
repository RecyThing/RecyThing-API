// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	entity "recything/features/user/entity"

	mock "github.com/stretchr/testify/mock"
)

// UsersRepositoryInterface is an autogenerated mock type for the UsersRepositoryInterface type
type UsersRepositoryInterface struct {
	mock.Mock
}

// FindByEmail provides a mock function with given fields: email
func (_m *UsersRepositoryInterface) FindByEmail(email string) (entity.UsersCore, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for FindByEmail")
	}

	var r0 entity.UsersCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.UsersCore, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) entity.UsersCore); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(entity.UsersCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: userID
func (_m *UsersRepositoryInterface) FindById(userID string) (entity.UsersCore, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 entity.UsersCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.UsersCore, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(string) entity.UsersCore); ok {
		r0 = rf(userID)
	} else {
		r0 = ret.Get(0).(entity.UsersCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *UsersRepositoryInterface) GetById(id string) (entity.UsersCore, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 entity.UsersCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.UsersCore, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) entity.UsersCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(entity.UsersCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByVerificationToken provides a mock function with given fields: token
func (_m *UsersRepositoryInterface) GetByVerificationToken(token string) (entity.UsersCore, error) {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for GetByVerificationToken")
	}

	var r0 entity.UsersCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.UsersCore, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) entity.UsersCore); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(entity.UsersCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// JoinCommunity provides a mock function with given fields: communityId, userId
func (_m *UsersRepositoryInterface) JoinCommunity(communityId string, userId string) error {
	ret := _m.Called(communityId, userId)

	if len(ret) == 0 {
		panic("no return value specified for JoinCommunity")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(communityId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPassword provides a mock function with given fields: email, data
func (_m *UsersRepositoryInterface) NewPassword(email string, data entity.UsersCore) (entity.UsersCore, error) {
	ret := _m.Called(email, data)

	if len(ret) == 0 {
		panic("no return value specified for NewPassword")
	}

	var r0 entity.UsersCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, entity.UsersCore) (entity.UsersCore, error)); ok {
		return rf(email, data)
	}
	if rf, ok := ret.Get(0).(func(string, entity.UsersCore) entity.UsersCore); ok {
		r0 = rf(email, data)
	} else {
		r0 = ret.Get(0).(entity.UsersCore)
	}

	if rf, ok := ret.Get(1).(func(string, entity.UsersCore) error); ok {
		r1 = rf(email, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: data
func (_m *UsersRepositoryInterface) Register(data entity.UsersCore) (entity.UsersCore, error) {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 entity.UsersCore
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.UsersCore) (entity.UsersCore, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(entity.UsersCore) entity.UsersCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(entity.UsersCore)
	}

	if rf, ok := ret.Get(1).(func(entity.UsersCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetOTP provides a mock function with given fields: otp
func (_m *UsersRepositoryInterface) ResetOTP(otp string) (entity.UsersCore, error) {
	ret := _m.Called(otp)

	if len(ret) == 0 {
		panic("no return value specified for ResetOTP")
	}

	var r0 entity.UsersCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.UsersCore, error)); ok {
		return rf(otp)
	}
	if rf, ok := ret.Get(0).(func(string) entity.UsersCore); ok {
		r0 = rf(otp)
	} else {
		r0 = ret.Get(0).(entity.UsersCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(otp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendOTP provides a mock function with given fields: emailUser, otp, expiry
func (_m *UsersRepositoryInterface) SendOTP(emailUser string, otp string, expiry int64) (entity.UsersCore, error) {
	ret := _m.Called(emailUser, otp, expiry)

	if len(ret) == 0 {
		panic("no return value specified for SendOTP")
	}

	var r0 entity.UsersCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, int64) (entity.UsersCore, error)); ok {
		return rf(emailUser, otp, expiry)
	}
	if rf, ok := ret.Get(0).(func(string, string, int64) entity.UsersCore); ok {
		r0 = rf(emailUser, otp, expiry)
	} else {
		r0 = ret.Get(0).(entity.UsersCore)
	}

	if rf, ok := ret.Get(1).(func(string, string, int64) error); ok {
		r1 = rf(emailUser, otp, expiry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBadge provides a mock function with given fields: id
func (_m *UsersRepositoryInterface) UpdateBadge(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBadge")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateById provides a mock function with given fields: id, data
func (_m *UsersRepositoryInterface) UpdateById(id string, data entity.UsersCore) error {
	ret := _m.Called(id, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, entity.UsersCore) error); ok {
		r0 = rf(id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateIsVerified provides a mock function with given fields: id, isVerified
func (_m *UsersRepositoryInterface) UpdateIsVerified(id string, isVerified bool) error {
	ret := _m.Called(id, isVerified)

	if len(ret) == 0 {
		panic("no return value specified for UpdateIsVerified")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool) error); ok {
		r0 = rf(id, isVerified)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdatePassword provides a mock function with given fields: id, data
func (_m *UsersRepositoryInterface) UpdatePassword(id string, data entity.UsersCore) error {
	ret := _m.Called(id, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, entity.UsersCore) error); ok {
		r0 = rf(id, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUserPoint provides a mock function with given fields: id, point
func (_m *UsersRepositoryInterface) UpdateUserPoint(id string, point int) error {
	ret := _m.Called(id, point)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserPoint")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int) error); ok {
		r0 = rf(id, point)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyOTP provides a mock function with given fields: email, otp
func (_m *UsersRepositoryInterface) VerifyOTP(email string, otp string) (entity.UsersCore, error) {
	ret := _m.Called(email, otp)

	if len(ret) == 0 {
		panic("no return value specified for VerifyOTP")
	}

	var r0 entity.UsersCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (entity.UsersCore, error)); ok {
		return rf(email, otp)
	}
	if rf, ok := ret.Get(0).(func(string, string) entity.UsersCore); ok {
		r0 = rf(email, otp)
	} else {
		r0 = ret.Get(0).(entity.UsersCore)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, otp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUsersRepositoryInterface creates a new instance of UsersRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsersRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UsersRepositoryInterface {
	mock := &UsersRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
