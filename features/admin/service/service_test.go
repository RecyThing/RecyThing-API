package service

import (
	"errors"
	"mime/multipart"

	"recything/features/admin/entity"
	user "recything/features/user/entity"
	"recything/mocks"

	"recything/utils/constanta"
	"recything/utils/pagination"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAdmin(t *testing.T) {
	mockData := entity.AdminCore{
		Fullname:        "John Doe",
		Email:           "john@example.com",
		Password:        "password123",
		ConfirmPassword: "password123",
		Status:          "aktif",
	}

	t.Run("Succes Create", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "johnny@example.com",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "aktif",
		}
		mockRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(errors.New("not found"))
		mockRepo.On("Create", mock.AnythingOfType("*multipart.FileHeader"), requestBody).Return(requestBody, nil)

		_, err := adminService.Create(nil, requestBody)

		assert.NoError(t, err)
		assert.NotEqual(t, requestBody.Email, mockData.Email)

		mockRepo.AssertExpectations(t)
	})
	t.Run("Data Empty", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "",
			Email:           "",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "aktif",
		}

		_, err := adminService.Create(nil, requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Status Input Invalid", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "john@example.com",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "berjalan",
		}

		_, err := adminService.Create(nil, requestBody)

		assert.Error(t, err)
		assert.NotEqualValues(t, []string{"aktif", "tidak aktif"}, requestBody.Status)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Email Invalid Fomat", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "johnexecom",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "aktif",
		}

		_, err := adminService.Create(nil, requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Password Invalid Length", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "john@example.com",
			Password:        "pass",
			ConfirmPassword: "pass",
			Status:          "aktif",
		}

		_, err := adminService.Create(nil, requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Email Registered", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "john@example.com",
			Password:        "password123456",
			ConfirmPassword: "password123456",
			Status:          "aktif",
		}

		mockRepo.On("FindByEmail", mock.AnythingOfType("string")).Return(errors.New("not found"))
		mockRepo.On("Create", mock.AnythingOfType("*multipart.FileHeader"), requestBody).Return(requestBody, errors.New("failed"))

		_, err := adminService.Create(nil, requestBody)

		assert.Error(t, err)
		assert.Equal(t, requestBody.Email, mockData.Email)

		mockRepo.AssertExpectations(t)
	})

	// INFO THIS FUNCTION
	// - Password Not Match No testing

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

	// Mock data
	mockData := entity.AdminCore{
		Id:       "1",
		Fullname: "John Doe",
		Email:    "john@example.com",
		Status:   "aktif",
	}

	t.Run("Succes GetBYID", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		mockRepo.On("SelectById", mock.AnythingOfType("string")).Return(mockData, nil)

		admin, err := adminService.GetById("1")

		assert.NoError(t, err)
		assert.Equal(t, mockData.Id, admin.Id)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		adminID := "2"
		mockRepo.On("SelectById", mock.AnythingOfType("string")).Return(mockData, errors.New(constanta.ERROR_RECORD_NOT_FOUND))

		admin, err := adminService.GetById(adminID)

		assert.Error(t, err)
		assert.NotEqual(t, adminID, mockData.Id)
		assert.Empty(t, admin)

		mockRepo.AssertExpectations(t)
	})

}

func TestUpdateAdminById(t *testing.T) {

	// Mock data
	updateAdmin := entity.AdminCore{
		Fullname:        "Updated Admin",
		Email:           "updatedadmin@example.com",
		Password:        "123456789",
		ConfirmPassword: "123456789",
		Status:          "aktif",
	}

	t.Run("Succes Update", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)
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
	})
	t.Run("Data Empty", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "",
			Email:           "",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "aktif",
		}

		err := adminService.UpdateById(nil, "1", requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Status Input Invalid", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "john@example.com",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "berjalan",
		}

		err := adminService.UpdateById(nil, "1", requestBody)

		assert.Error(t, err)
		assert.NotEqualValues(t, []string{"aktif", "tidak aktif"}, requestBody.Status)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Email Invalid Fomat", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "johnexecom",
			Password:        "password123",
			ConfirmPassword: "password123",
			Status:          "aktif",
		}

		err := adminService.UpdateById(nil, "1", requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Password Invalid Length", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		requestBody := entity.AdminCore{
			Fullname:        "John Doe",
			Email:           "john@example.com",
			Password:        "pass",
			ConfirmPassword: "pass",
			Status:          "aktif",
		}

		err := adminService.UpdateById(nil, "1", requestBody)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

}

func TestDeleteAdmin(t *testing.T) {

	dataAdmin := entity.AdminCore{
		Id:       "1",
		Email:    "admin@example.com",
		Password: "hashedpassword",
	}

	t.Run("Succes Delete", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)

		adminID := "1"
		mockRepo.On("Delete", mock.AnythingOfType("string")).Return(nil)

		// Test case
		err := adminService.DeleteById(adminID)

		assert.NoError(t, err)
		assert.Equal(t, adminID, dataAdmin.Id)
		mockRepo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		mockRepo := new(mocks.AdminRepositoryInterface)
		adminService := NewAdminService(mockRepo)
		adminID := "2"

		mockRepo.On("Delete", mock.AnythingOfType("string")).Return(errors.New("failed"))

		// Test case
		err := adminService.DeleteById(adminID)

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)

	})

}

func TestAdminService_FindByEmailANDPassword(t *testing.T) {
	dataAdmin := entity.AdminCore{
		Email:    "admin@example.com",
		Password: "hashedpassword",
	}

	t.Run("Succes Login", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)
		mockAdmin := entity.AdminCore{
			Email:    "admin@example.com",
			Password: "hashedpassword",
		}

		mockRepo.On("FindByEmailANDPassword", mockAdmin).Return(mockAdmin, nil)

		// Function Test
		admin, token, err := adminService.FindByEmailANDPassword(mockAdmin)

		assert.NoError(t, err)
		assert.Equal(t, dataAdmin, admin)
		assert.NotEmpty(t, token)
		mockRepo.AssertExpectations(t)

	})
	t.Run("Wrong Email or Password", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)
		mockAdmin := entity.AdminCore{
			Email:    "admin@example.com",
			Password: "hashedpasswordnewwrong",
		}

		mockRepo.On("FindByEmailANDPassword", mockAdmin).Return(mockAdmin, errors.New("failed"))

		// Function Test
		_, token, err := adminService.FindByEmailANDPassword(mockAdmin)

		assert.NotEqual(t, mockAdmin, dataAdmin)
		assert.EqualError(t, err, "error : email atau password salah")
		assert.Empty(t, token)
		mockRepo.AssertExpectations(t)
	})

}

// Manage User
func TestGetAllUsers(t *testing.T) {
	expectedUsers := []user.UsersCore{
		{Id: "1", Fullname: "User1", Email: "user1@example.com", Point: 20000},
		{Id: "2", Fullname: "User2", Email: "user2@example.com", Point: 3000},
	}

	t.Run("Succes GetAll", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)
		mockRepo.On("GetAllUsers", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedUsers, pagination.PageInfo{}, len(expectedUsers), nil)

		// Panggil fungsi GetAllUsers dari AdminService
		users, _, count, err := adminService.GetAllUsers("", "", "")

		assert.NoError(t, err)
		assert.NotNil(t, users)
		assert.NotEmpty(t, count)
		assert.Equal(t, len(expectedUsers), count)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Limit > 10", func(t *testing.T) {
		mockRepo := mocks.NewAdminRepositoryInterface(t)
		adminService := NewAdminService(mockRepo)
		users, pageInfo, count, err := adminService.GetAllUsers("", "", "20")

		assert.Error(t, err)
		assert.Nil(t, users)
		assert.Equal(t, pagination.PageInfo{}, pageInfo)
		assert.Equal(t, 0, count)
	})

}
