package service

import (
	"errors"
	"mime/multipart"
	"testing"
	"recything/features/admin/entity"
	"recything/mocks"
	"recything/utils/pagination"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAdmin(t *testing.T) {
	mockRepo := new(mocks.AdminRepositoryInterface)
	adminService := NewAdminService(mockRepo)

	// Mock data
	mockData := entity.AdminCore{
		Fullname:        "John Doe",
		Email:           "john@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
		Status:          "aktif",
	}

	mockRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(errors.New("not found"))
	mockRepo.On("Create", mock.AnythingOfType("*multipart.FileHeader"), mockData).Return(mockData, nil)

	createdAdmin, err := adminService.Create(nil, mockData)

	assert.NoError(t, err)
	assert.Equal(t, mockData.Fullname, createdAdmin.Fullname)
	assert.Equal(t, mockData.Email, createdAdmin.Email)

	mockRepo.AssertExpectations(t)
}

func TestGetAllAdmins(t *testing.T) {
	mockRepo := new(mocks.AdminRepositoryInterface)
	adminService := NewAdminService(mockRepo)

	// Mock data
	mockData := []entity.AdminCore{
		{Id: "1", Fullname: "John Doe", Email: "john@example.com", Status: "aktif"},
		{Id: "2", Fullname: "Jane Doe", Email: "jane@example.com", Status: "aktif"},
	}

	// Mock repository response
	mockRepo.On("SelectAll", mock.AnythingOfType("int"), mock.AnythingOfType("int"), mock.AnythingOfType("string")).
		Return(mockData, pagination.PageInfo{}, len(mockData), nil)

	// Test case
	admins, _, _, err := adminService.GetAll("1", "10", "")

	assert.NoError(t, err)
	assert.Len(t, admins, len(mockData))
	mockRepo.AssertExpectations(t)
}

func TestGetAdminById(t *testing.T) {
	mockRepo := new(mocks.AdminRepositoryInterface)
	adminService := NewAdminService(mockRepo)

	// Mock data
	mockData := entity.AdminCore{
		Id:       "1",
		Fullname: "John Doe",
		Email:    "john@example.com",
		Status:   "aktif",
	}

	mockRepo.On("SelectById", mock.AnythingOfType("string")).Return(mockData, nil)

	admin, err := adminService.GetById("1")

	assert.NoError(t, err)
	assert.Equal(t, mockData.Id, admin.Id)
	mockRepo.AssertExpectations(t)
}

func TestUpdateAdminById(t *testing.T) {
	mockRepo := new(mocks.AdminRepositoryInterface)
	adminService := NewAdminService(mockRepo)

	// Mock data
	updateAdmin := entity.AdminCore{
		Fullname:        "Updated Admin",
		Email:           "updatedadmin@example.com",
		Password:        "123456789",
		ConfirmPassword: "123456789",
		Status:          "aktif",
	}

	mockFileHeader := &multipart.FileHeader{
		Filename: "testfile.png",
	}

	mockRepo.On("Update", mock.AnythingOfType("*multipart.FileHeader"), mock.AnythingOfType("string"), mock.AnythingOfType("entity.AdminCore")).
		Return(nil)
	mockRepo.On("SelectById", mock.AnythingOfType("string")).Return(updateAdmin, nil)

	// Test case
	admin, _ := adminService.GetById("1")
	err := adminService.UpdateById(mockFileHeader, admin.Id, updateAdmin)

	// Assertions
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}


func TestDeleteAdmin(t *testing.T) {
	mockRepo := new(mocks.AdminRepositoryInterface)
	adminService := NewAdminService(mockRepo)

	adminID := "1"
	mockRepo.On("Delete", mock.AnythingOfType("string")).
		Return(nil)

	// Test case
	err := adminService.DeleteById(adminID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
