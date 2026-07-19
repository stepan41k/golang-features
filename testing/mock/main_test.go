package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepo struct {
	mock.Mock
}

func (m *MockProductRepo) GetStockCount(productID int) (int, error) {
	args := m.Called(productID)
	return args.Int(0), args.Error(1)
}

func TestCanShowAd(t *testing.T) {
	t.Run("Should return true when stock is plenty", func(t *testing.T) {
		mockRepo := new(MockProductRepo)

		mockRepo.On("GetStockCount", 123).Return(15, nil)

		service := NewAdService(mockRepo)

		result, err := service.CanShowAd(123)

		assert.NoError(t, err)
		assert.True(t, result)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Should return false when repo fails", func(t *testing.T) {
		mockRepo := new(MockProductRepo)

		mockRepo.On("GetStockCount", 999).Return(0, errors.New("db error"))

		service := NewAdService(mockRepo)

		result, err := service.CanShowAd(999)

		assert.Error(t, err)
		assert.False(t, result)

		mockRepo.AssertExpectations(t)
	})
}