package usecase_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/entity"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/usecase"
	"github.com/ruancaetano/go-gin-fruits/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewUpdateFruitUseCase(t *testing.T) {
	repository := &mocks.FruitRepositoryMock{}
	u := usecase.NewUpdateFruitUseCase(repository)
	assert.NotNil(t, u)
}

func TestUpdateFruitUseCase_Execute(t *testing.T) {
	t.Run("With Invalid Params", func(t *testing.T) {
		repository := &mocks.FruitRepositoryMock{}
		repository.On("Save", mock.Anything).Return(nil)
		u := usecase.NewUpdateFruitUseCase(repository)

		input := &usecase.UpdateFruitUseCaseInputDTO{
			ID:       "",
			Quantity: 0,
			Price:    0,
		}

		output, err := u.Execute(context.Background(), input)
		assert.Nil(t, output)
		assert.Error(t, err, "id is required")

		input.ID = "id"
		output, err = u.Execute(context.Background(), input)
		assert.Nil(t, output)
		assert.Error(t, err, "quantity must be greater than zero")

		input.Quantity = 1
		output, err = u.Execute(context.Background(), input)
		assert.Nil(t, output)
		assert.Error(t, err, "price must be greater than zero")
	})

	t.Run("Fail if fruit not found", func(t *testing.T) {
		repository := &mocks.FruitRepositoryMock{}
		repository.On("Get", mock.Anything, mock.Anything).Return(&entity.Fruit{}, errors.New("fruit not found"))
		u := usecase.NewUpdateFruitUseCase(repository)

		output, err := u.Execute(context.Background(), &usecase.UpdateFruitUseCaseInputDTO{
			ID:       "not-found-id",
			Price:    10.0,
			Quantity: 1,
		})

		fmt.Println(err)
		assert.Nil(t, output)
		assert.EqualError(t, err, "fruit not found")
	})

	t.Run("Fail if repository save fail", func(t *testing.T) {
		fruitMock, err := entity.NewFruit("fruit", "owner", 1, 10.0)
		assert.Nil(t, err)

		repository := &mocks.FruitRepositoryMock{}
		repository.On("Get", mock.Anything, mock.Anything).Return(fruitMock, nil)
		repository.On("Save", mock.Anything, mock.Anything).Return(errors.New("repository save fail"))
		u := usecase.NewUpdateFruitUseCase(repository)

		output, err := u.Execute(context.Background(), &usecase.UpdateFruitUseCaseInputDTO{
			ID:       fruitMock.ID,
			Price:    10.0,
			Quantity: 1,
		})

		fmt.Println(err)
		assert.Nil(t, output)
		assert.EqualError(t, err, "repository save fail")
	})

	t.Run("With Valid Params", func(t *testing.T) {
		fruitMock, err := entity.NewFruit("fruit", "owner", 1, 10.0)
		// just to skip the pass by reference
		fruitMockCopy := *fruitMock
		assert.Nil(t, err)

		repository := &mocks.FruitRepositoryMock{}
		repository.On("Get", mock.Anything, mock.Anything).Return(&fruitMockCopy, nil)
		repository.On("Save", mock.Anything, mock.Anything).Return(nil)

		u := usecase.NewUpdateFruitUseCase(repository)

		input := &usecase.UpdateFruitUseCaseInputDTO{
			ID:       fruitMock.ID,
			Quantity: 100,
			Price:    100.0,
		}

		output, err := u.Execute(context.Background(), input)

		repository.AssertNumberOfCalls(t, "Save", 1)

		assert.Nil(t, err)
		assert.Equal(t, output.ID, fruitMock.ID)
		assert.True(t, output.CreatedAt.Equal(fruitMock.CreatedAt))
		assert.False(t, output.UpdatedAt.Equal(fruitMock.UpdatedAt))
		assert.Equal(t, output.Name, fruitMock.Name)
		assert.Equal(t, output.Owner, fruitMock.Owner)
		assert.Equal(t, output.Quantity, 100)
		assert.Equal(t, output.Price, 100.0)
		assert.Equal(t, output.Status, fruitMock.Status)
	})
}
