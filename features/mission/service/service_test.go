package service

import (
	"errors"
	admin "recything/features/admin/entity"
	"recything/features/mission/entity"
	user "recything/features/user/entity"

	"recything/mocks"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteMission(t *testing.T) {
	t.Run("not founc", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)
		missionID := "1"
		expectedError := errors.New("mission not found") // Define the expected error

		missionRepo.On("DeleteMission", missionID).Return(expectedError).Once()

		err := missionService.DeleteMission(missionID)

		assert.EqualError(t, err, expectedError.Error())

		missionRepo.AssertExpectations(t)
	})
	t.Run("success", func(t *testing.T) {
		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		missionID := "1"

		// Set the expectations for a successful deletion by ID
		missionRepo.On("DeleteMission", missionID).Return(nil).Once()

		err := missionService.DeleteMission(missionID)

		assert.NoError(t, err)
		missionRepo.AssertExpectations(t)

	})
}

func TestGetMissionByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		data := entity.Mission{
			ID:               "1",
			Title:            "title",
			Creator:          "creator",
			Status:           "status",
			AdminID:          "1",
			MissionImage:     "image",
			Point:            10,
			Description:      "desc",
			StartDate:        "2023-12-12",
			EndDate:          "2023-12-13",
			TitleStage:       "stage title",
			DescriptionStage: "description title",
		}
		missionID := "1"

		// Set expectations for the FindById call in the missionRepo
		missionRepo.On("FindById", missionID).Return(data, nil).Once()

		// Mock the behavior of SelectById in adminRepo
		dataadmin := admin.AdminCore{
			Id:              "1",
			Fullname:        "admin",
			Image:           "image",
			Role:            "admin",
			Email:           "admin@gmail.com",
			Password:        "12345678",
			ConfirmPassword: "12345678",
			Status:          "aktif",
		}
		adminRepo.On("SelectById", data.AdminID).Return(dataadmin, nil).Once()

		result, err := missionService.FindById(missionID)

		assert.NoError(t, err)
		assert.Equal(t, data.ID, result.ID)
		assert.Equal(t, dataadmin.Fullname, result.Creator)
		missionRepo.AssertExpectations(t)
		adminRepo.AssertExpectations(t)
	})

	t.Run("admin not found", func(t *testing.T) {

		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		// Create mission service instance with mock repositories
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		// Define test data
		data := entity.Mission{
			ID:               "1",
			Title:            "title",
			Creator:          "creator",
			Status:           "status",
			AdminID:          "1",
			MissionImage:     "image",
			Point:            10,
			Description:      "desc",
			StartDate:        "2023-12-12",
			EndDate:          "2023-12-13",
			TitleStage:       "stage title",
			DescriptionStage: "description title",
		}
		missionID := "1"

		// dataadmin := admin.AdminCore{
		// 	Id:              "2",
		// 	Fullname:        "admin",
		// 	Image:           "image",
		// 	Role:            "admin",
		// 	Email:           "admin@gmail.com",
		// 	Password:        "12345678",
		// 	ConfirmPassword: "12345678",
		// 	Status:          "aktif",
		// }
		missionRepo.On("FindById", missionID).Return(data, errors.New("data tidak ditemukan")).Once()
		result, err := missionService.FindById(missionID)

		adminRepo.On("SelectById", data.AdminID).Return(admin.AdminCore{}, err).Once()
		adminRepo.SelectById(data.AdminID)
		assert.Error(t, err)
		assert.NotEqual(t, data.ID, result.ID)

		missionRepo.AssertExpectations(t)
		adminRepo.AssertExpectations(t)
	})
	t.Run("mission not found", func(t *testing.T) {

		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		// Create mission service instance with mock repositories
		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

		// Define test data
		data := entity.Mission{
			ID:               "1",
			Title:            "title",
			Creator:          "creator",
			Status:           "status",
			AdminID:          "1",
			MissionImage:     "image",
			Point:            10,
			Description:      "desc",
			StartDate:        "2023-12-12",
			EndDate:          "2023-12-13",
			TitleStage:       "stage title",
			DescriptionStage: "description title",
		}
		missionID := "2"

		dataadmin := admin.AdminCore{
			Id:              "1",
			Fullname:        "admin",
			Image:           "image",
			Role:            "admin",
			Email:           "admin@gmail.com",
			Password:        "12345678",
			ConfirmPassword: "12345678",
			Status:          "aktif",
		}
		missionRepo.On("FindById", missionID).Return(data, errors.New("data tidak ditemukan")).Once()
		result, err := missionService.FindById(missionID)

		adminRepo.On("SelectById", data.AdminID).Return(dataadmin, err).Once()
		adminRepo.SelectById(data.AdminID)
		assert.Error(t, err)
		assert.NotEqual(t, data.ID, result.ID)

		missionRepo.AssertExpectations(t)
		adminRepo.AssertExpectations(t)
	})
}

func TestFindMissionApprovalById(t *testing.T) {
	t.Run("success", func(t *testing.T) {

		missionRepo := new(mocks.MissionRepositoryInterface)
		adminRepo := new(mocks.AdminRepositoryInterface)
		userRepo := new(mocks.UsersRepositoryInterface)

		missionService := NewMissionService(missionRepo, adminRepo, userRepo)
		uploadMissionID := "1"
		missionUpproval := entity.UploadMissionTaskCore{
			ID:          "1",
			UserID:      "1",
			User:        "YAYA",
			MissionID:   "1",
			MissionName: "mission name",
			Description: "description",
			Reason:      "ini reason",
			Status:      "perlu tinjauaan",
		}

		missionRepo.On("FindMissionApprovalById", uploadMissionID).Return(missionUpproval, nil).Once()
		missionRepo.On("FindById", missionUpproval.MissionID).Return(entity.Mission{
			ID:    missionUpproval.MissionID,
			Title: "Mission Title",
		}, nil).Once()

		userRepo.On("GetById", missionUpproval.UserID).Return(user.UsersCore{
			Id:       "1",
			Fullname: "User Fullname",
		}, nil).Once()

		result, err := missionService.FindMissionApprovalById(uploadMissionID)

		assert.NoError(t, err)
		assert.Equal(t, missionUpproval.ID, result.ID)
		assert.Equal(t, "Mission Title", result.MissionName)
		assert.Equal(t, "User Fullname", result.User)
		missionRepo.AssertExpectations(t)
		userRepo.AssertExpectations(t)
	})
}

// func TestUpdateStatusMissionApproval(t *testing.T) {

// 	t.Run("Approval Success with Reason", func(t *testing.T) {
// 		missionRepo := new(mocks.MissionRepositoryInterface)
// 		adminRepo := new(mocks.AdminRepositoryInterface)
// 		userRepo := new(mocks.UsersRepositoryInterface)
// 		missionService := NewMissionService(missionRepo, adminRepo, userRepo)

// 		status := "disetujui"
// 		reason := "approval reason"

// 		task := entity.UploadMissionTaskCore{
// 			ID:        "1",
// 			UserID:    "1",
// 			User:      "user",
// 			MissionID: "1",
// 		}

// 		expectedTask := entity.UploadMissionTaskCore{
// 			ID:        "1",
// 			UserID:    "1",
// 			User:      "user",
// 			MissionID: "1",
// 			Status:    "disetujui",
// 			Reason:    "approval reason",
// 		}

// 		// mission := entity.Mission{
// 		// 	ID:               "1",
// 		// 	Title:            "title",
// 		// 	Creator:          "creator",
// 		// 	Status:           "active",
// 		// 	AdminID:          "1",
// 		// 	MissionImage:     "image",
// 		// 	Point:            10,
// 		// 	Description:      "desc",
// 		// 	StartDate:        "2023-12-12",
// 		// 	EndDate:          "2023-12-13",
// 		// 	TitleStage:       "stage title",
// 		// 	DescriptionStage: "description title",
// 		// }

		
		
// 		missionRepo.On("UpdateStatusMissionApproval", task.ID, status, reason).Return(nil)
// 		err := missionService.UpdateStatusMissionApproval(task.ID, status, reason)

// 		assert.Error(t, err)
// 		assert.Equal(t, expectedTask.Status, status)
// 		userRepo.AssertExpectations(t )
// 	})

// }
