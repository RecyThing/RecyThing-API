package service

import (
	// "errors"
	// "mime/multipart"

	user_entity "recything/features/user/entity"
	"recything/mocks"

	// "recything/utils/constanta"
	// "recything/utils/pagination"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockData = []map[string]interface{}{
	{"ID": 1, "Point": 200, "Description": "Poin hari 1"},
	{"ID": 2, "Point": 200, "Description": "Poin hari 2"},
	{"ID": 3, "Point": 200, "Description": "Poin hari 3"},
}

var dataHistory = map[string]interface{}{
	"ID": 1, "Point": 200, "Description": "Poin hari 1", "transactionId": "2",
}

var claimedData = []user_entity.UserDailyPointsCore{
	{UsersID: "1", DailyPointID: 1, Claim: false},
	{UsersID: "2", DailyPointID: 2, Claim: true},
	{UsersID: "3", DailyPointID: 3, Claim: false},
}

func TestGetAllHistoryPoint(t *testing.T) {
	mockRepo := new(mocks.DailyPointRepositoryInterface)
	dailyPointService := NewDailyPointService(mockRepo)

	userId := "1"

	// Mock repository response
	mockRepo.On("GetAllHistoryPoint", userId).
		Return(mockData, nil)

	// Test case
	dailypoint, err := dailyPointService.GetAllHistoryPoint(userId)

	assert.NoError(t, err)
	assert.Len(t, dailypoint, len(mockData))
	mockRepo.AssertExpectations(t)
}

func TestGetByIdHistoryPoint(t *testing.T) {
	mockRepo := new(mocks.DailyPointRepositoryInterface)
	dailyPointService := NewDailyPointService(mockRepo)

	userId := "1"
	transactionId := "2"

	// Mock repository response
	mockRepo.On("GetByIdHistoryPoint", userId, transactionId).
		Return(dataHistory, nil)

	// Test case
	dailypoint, err := dailyPointService.GetByIdHistoryPoint(userId, transactionId)

	assert.NoError(t, err)
	assert.NotNil(t, dailypoint)
	mockRepo.AssertExpectations(t)
}

func TestGetAllClaimedDaily(t *testing.T) {
	mockRepo := new(mocks.DailyPointRepositoryInterface)
	dailyPointService := NewDailyPointService(mockRepo)

	userId := "1"

	// Mock repository response
	mockRepo.On("GetAllClaimedDaily", userId).
		Return(claimedData, nil)

	// Test case
	dailypoint, err := dailyPointService.GetAllClaimedDaily(userId)

	assert.NoError(t, err)
	assert.Len(t, dailypoint, len(claimedData))
	mockRepo.AssertExpectations(t)
}
