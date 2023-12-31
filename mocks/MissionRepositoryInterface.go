// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	entity "recything/features/mission/entity"
	helper "recything/utils/helper"

	mock "github.com/stretchr/testify/mock"

	model "recything/features/mission/model"

	multipart "mime/multipart"

	pagination "recything/utils/pagination"
)

// MissionRepositoryInterface is an autogenerated mock type for the MissionRepositoryInterface type
type MissionRepositoryInterface struct {
	mock.Mock
}

// ClaimMission provides a mock function with given fields: userID, data
func (_m *MissionRepositoryInterface) ClaimMission(userID string, data entity.ClaimedMission) error {
	ret := _m.Called(userID, data)

	if len(ret) == 0 {
		panic("no return value specified for ClaimMission")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, entity.ClaimedMission) error); ok {
		r0 = rf(userID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMission provides a mock function with given fields: input
func (_m *MissionRepositoryInterface) CreateMission(input entity.Mission) error {
	ret := _m.Called(input)

	if len(ret) == 0 {
		panic("no return value specified for CreateMission")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Mission) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUploadMissionTask provides a mock function with given fields: userID, data, images
func (_m *MissionRepositoryInterface) CreateUploadMissionTask(userID string, data entity.UploadMissionTaskCore, images []*multipart.FileHeader) (entity.UploadMissionTaskCore, error) {
	ret := _m.Called(userID, data, images)

	if len(ret) == 0 {
		panic("no return value specified for CreateUploadMissionTask")
	}

	var r0 entity.UploadMissionTaskCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, entity.UploadMissionTaskCore, []*multipart.FileHeader) (entity.UploadMissionTaskCore, error)); ok {
		return rf(userID, data, images)
	}
	if rf, ok := ret.Get(0).(func(string, entity.UploadMissionTaskCore, []*multipart.FileHeader) entity.UploadMissionTaskCore); ok {
		r0 = rf(userID, data, images)
	} else {
		r0 = ret.Get(0).(entity.UploadMissionTaskCore)
	}

	if rf, ok := ret.Get(1).(func(string, entity.UploadMissionTaskCore, []*multipart.FileHeader) error); ok {
		r1 = rf(userID, data, images)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteMission provides a mock function with given fields: missionID
func (_m *MissionRepositoryInterface) DeleteMission(missionID string) error {
	ret := _m.Called(missionID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMission")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(missionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAllMission provides a mock function with given fields: page, limit, search, filter
func (_m *MissionRepositoryInterface) FindAllMission(page int, limit int, search string, filter string) ([]entity.Mission, pagination.PageInfo, helper.CountMission, error) {
	ret := _m.Called(page, limit, search, filter)

	if len(ret) == 0 {
		panic("no return value specified for FindAllMission")
	}

	var r0 []entity.Mission
	var r1 pagination.PageInfo
	var r2 helper.CountMission
	var r3 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]entity.Mission, pagination.PageInfo, helper.CountMission, error)); ok {
		return rf(page, limit, search, filter)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []entity.Mission); ok {
		r0 = rf(page, limit, search, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Mission)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) pagination.PageInfo); ok {
		r1 = rf(page, limit, search, filter)
	} else {
		r1 = ret.Get(1).(pagination.PageInfo)
	}

	if rf, ok := ret.Get(2).(func(int, int, string, string) helper.CountMission); ok {
		r2 = rf(page, limit, search, filter)
	} else {
		r2 = ret.Get(2).(helper.CountMission)
	}

	if rf, ok := ret.Get(3).(func(int, int, string, string) error); ok {
		r3 = rf(page, limit, search, filter)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// FindAllMissionApproval provides a mock function with given fields: page, limit, search, filter
func (_m *MissionRepositoryInterface) FindAllMissionApproval(page int, limit int, search string, filter string) ([]entity.UploadMissionTaskCore, pagination.PageInfo, helper.CountMissionApproval, error) {
	ret := _m.Called(page, limit, search, filter)

	if len(ret) == 0 {
		panic("no return value specified for FindAllMissionApproval")
	}

	var r0 []entity.UploadMissionTaskCore
	var r1 pagination.PageInfo
	var r2 helper.CountMissionApproval
	var r3 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]entity.UploadMissionTaskCore, pagination.PageInfo, helper.CountMissionApproval, error)); ok {
		return rf(page, limit, search, filter)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []entity.UploadMissionTaskCore); ok {
		r0 = rf(page, limit, search, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.UploadMissionTaskCore)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) pagination.PageInfo); ok {
		r1 = rf(page, limit, search, filter)
	} else {
		r1 = ret.Get(1).(pagination.PageInfo)
	}

	if rf, ok := ret.Get(2).(func(int, int, string, string) helper.CountMissionApproval); ok {
		r2 = rf(page, limit, search, filter)
	} else {
		r2 = ret.Get(2).(helper.CountMissionApproval)
	}

	if rf, ok := ret.Get(3).(func(int, int, string, string) error); ok {
		r3 = rf(page, limit, search, filter)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// FindAllMissionUser provides a mock function with given fields: userID, filter
func (_m *MissionRepositoryInterface) FindAllMissionUser(userID string, filter string) ([]entity.MissionHistories, error) {
	ret := _m.Called(userID, filter)

	if len(ret) == 0 {
		panic("no return value specified for FindAllMissionUser")
	}

	var r0 []entity.MissionHistories
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]entity.MissionHistories, error)); ok {
		return rf(userID, filter)
	}
	if rf, ok := ret.Get(0).(func(string, string) []entity.MissionHistories); ok {
		r0 = rf(userID, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.MissionHistories)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: missionID
func (_m *MissionRepositoryInterface) FindById(missionID string) (entity.Mission, error) {
	ret := _m.Called(missionID)

	if len(ret) == 0 {
		panic("no return value specified for FindById")
	}

	var r0 entity.Mission
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.Mission, error)); ok {
		return rf(missionID)
	}
	if rf, ok := ret.Get(0).(func(string) entity.Mission); ok {
		r0 = rf(missionID)
	} else {
		r0 = ret.Get(0).(entity.Mission)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(missionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindClaimed provides a mock function with given fields: userID, missionID
func (_m *MissionRepositoryInterface) FindClaimed(userID string, missionID string) error {
	ret := _m.Called(userID, missionID)

	if len(ret) == 0 {
		panic("no return value specified for FindClaimed")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, missionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindHistoryById provides a mock function with given fields: userID, transactionID
func (_m *MissionRepositoryInterface) FindHistoryById(userID string, transactionID string) (entity.UploadMissionTaskCore, error) {
	ret := _m.Called(userID, transactionID)

	if len(ret) == 0 {
		panic("no return value specified for FindHistoryById")
	}

	var r0 entity.UploadMissionTaskCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (entity.UploadMissionTaskCore, error)); ok {
		return rf(userID, transactionID)
	}
	if rf, ok := ret.Get(0).(func(string, string) entity.UploadMissionTaskCore); ok {
		r0 = rf(userID, transactionID)
	} else {
		r0 = ret.Get(0).(entity.UploadMissionTaskCore)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, transactionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindHistoryByIdTransaction provides a mock function with given fields: userID, transactionID
func (_m *MissionRepositoryInterface) FindHistoryByIdTransaction(userID string, transactionID string) (map[string]interface{}, error) {
	ret := _m.Called(userID, transactionID)

	if len(ret) == 0 {
		panic("no return value specified for FindHistoryByIdTransaction")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (map[string]interface{}, error)); ok {
		return rf(userID, transactionID)
	}
	if rf, ok := ret.Get(0).(func(string, string) map[string]interface{}); ok {
		r0 = rf(userID, transactionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, transactionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMissionApprovalById provides a mock function with given fields: UploadMissionTaskID
func (_m *MissionRepositoryInterface) FindMissionApprovalById(UploadMissionTaskID string) (entity.UploadMissionTaskCore, error) {
	ret := _m.Called(UploadMissionTaskID)

	if len(ret) == 0 {
		panic("no return value specified for FindMissionApprovalById")
	}

	var r0 entity.UploadMissionTaskCore
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (entity.UploadMissionTaskCore, error)); ok {
		return rf(UploadMissionTaskID)
	}
	if rf, ok := ret.Get(0).(func(string) entity.UploadMissionTaskCore); ok {
		r0 = rf(UploadMissionTaskID)
	} else {
		r0 = ret.Get(0).(entity.UploadMissionTaskCore)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(UploadMissionTaskID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUploadById provides a mock function with given fields: id
func (_m *MissionRepositoryInterface) FindUploadById(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindUploadById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUploadMissionStatus provides a mock function with given fields: id, missionID, userID, status
func (_m *MissionRepositoryInterface) FindUploadMissionStatus(id string, missionID string, userID string, status string) error {
	ret := _m.Called(id, missionID, userID, status)

	if len(ret) == 0 {
		panic("no return value specified for FindUploadMissionStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(id, missionID, userID, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCountDataMission provides a mock function with given fields: filter, search
func (_m *MissionRepositoryInterface) GetCountDataMission(filter string, search string) (helper.CountMission, error) {
	ret := _m.Called(filter, search)

	if len(ret) == 0 {
		panic("no return value specified for GetCountDataMission")
	}

	var r0 helper.CountMission
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (helper.CountMission, error)); ok {
		return rf(filter, search)
	}
	if rf, ok := ret.Get(0).(func(string, string) helper.CountMission); ok {
		r0 = rf(filter, search)
	} else {
		r0 = ret.Get(0).(helper.CountMission)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(filter, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCountDataMissionApproval provides a mock function with given fields: search
func (_m *MissionRepositoryInterface) GetCountDataMissionApproval(search string) (helper.CountMissionApproval, error) {
	ret := _m.Called(search)

	if len(ret) == 0 {
		panic("no return value specified for GetCountDataMissionApproval")
	}

	var r0 helper.CountMissionApproval
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (helper.CountMissionApproval, error)); ok {
		return rf(search)
	}
	if rf, ok := ret.Get(0).(func(string) helper.CountMissionApproval); ok {
		r0 = rf(search)
	} else {
		r0 = ret.Get(0).(helper.CountMissionApproval)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetImageURL provides a mock function with given fields: missionID
func (_m *MissionRepositoryInterface) GetImageURL(missionID string) (string, error) {
	ret := _m.Called(missionID)

	if len(ret) == 0 {
		panic("no return value specified for GetImageURL")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(missionID)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(missionID)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(missionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMission provides a mock function with given fields: missionID, data
func (_m *MissionRepositoryInterface) UpdateMission(missionID string, data entity.Mission) error {
	ret := _m.Called(missionID, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMission")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, entity.Mission) error); ok {
		r0 = rf(missionID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateMissionStatus provides a mock function with given fields: data
func (_m *MissionRepositoryInterface) UpdateMissionStatus(data model.Mission) error {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMissionStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Mission) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStatusMissionApproval provides a mock function with given fields: uploadMissionTaskID, status, reason
func (_m *MissionRepositoryInterface) UpdateStatusMissionApproval(uploadMissionTaskID string, status string, reason string) error {
	ret := _m.Called(uploadMissionTaskID, status, reason)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatusMissionApproval")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(uploadMissionTaskID, status, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUploadMissionTask provides a mock function with given fields: id, images, data
func (_m *MissionRepositoryInterface) UpdateUploadMissionTask(id string, images []*multipart.FileHeader, data entity.UploadMissionTaskCore) error {
	ret := _m.Called(id, images, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUploadMissionTask")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []*multipart.FileHeader, entity.UploadMissionTaskCore) error); ok {
		r0 = rf(id, images, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMissionRepositoryInterface creates a new instance of MissionRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMissionRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MissionRepositoryInterface {
	mock := &MissionRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
