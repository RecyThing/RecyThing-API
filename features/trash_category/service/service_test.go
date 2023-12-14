package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"recything/features/trash_category/entity"
	"recything/features/trash_category/entity/mocks"
)

// MockRepository adalah mock dari TrashCategoryRepositoryInterface
type MockRepository struct {
	mock.Mock
}

// Implementasi method Create dari TrashCategoryRepositoryInterface untuk mock
func (m *MockRepository) Create(data entity.TrashCategoryCore) error {
	args := m.Called(data)
	return args.Error(0)
}

func TestTrashCategoryService_CreateCategory(t *testing.T) {
	// Membuat mock repository
	mockRepo := new(mocks.TrashCategoryRepositoryInterface)

	// Membuat instance baru dari service menggunakan mock repository
	trashService := NewTrashCategoryService(mockRepo)

	// Mock data untuk pengujian
	testData := entity.TrashCategoryCore{
		Unit:      "barang",
		TrashType: "kaca",
		Point:     10,
	}

	// Mengatur perilaku yang diharapkan dari mock repository untuk method Create
	mockRepo.On("Create", testData).Return(nil)

	// Pengujian kasus 1: Pembuatan berhasil
	err := trashService.CreateCategory(testData)
	assert.NoError(t, err, "Expected no error")


	
	// Mengatur perilaku dari mock repository untuk mengembalikan error
	mockRepo.On("Create", testData).Return(errors.New("failed to create"))

	// Pengujian kasus 4: Error dari repository saat pembuatan
	err = trashService.CreateCategory(testData)
	assert.NoError(t, err, "Expected no error")
}
