package usecase_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/usecase"
	"github.com/ruancaetano/go-gin-fruits/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewCreateFruitUseCase(t *testing.T) {
	repository := &mocks.FruitRepositoryMock{}
	u := usecase.NewCreateFruitUseCase(repository)
	assert.NotNil(t, u)
}

func TestCreateFruitUseCase_Execute(t *testing.T) {
	t.Run("With Invalid Params", func(t *testing.T) {
		repository := &mocks.FruitRepositoryMock{}
		repository.On("Save", mock.Anything).Return(nil)
		u := usecase.NewCreateFruitUseCase(repository)

		input := &usecase.CreateFruitUseCaseInputDTO{
			Name:     "",
			Owner:    "",
			Quantity: 0,
			Price:    0,
		}

		output, err := u.Execute(context.Background(), input)
		assert.Nil(t, output)
		assert.Error(t, err, "name is required")

		input.Name = "12312312@3123asdas"
		output, err = u.Execute(context.Background(), input)
		assert.Nil(t, output)
		assert.Error(t, err, "name cannot contain numbers or special characters")

		input.Name = "Name"
		output, err = u.Execute(context.Background(), input)
		assert.Nil(t, output)
		assert.Error(t, err, "owner is required")

		input.Owner = "Owner"
		output, err = u.Execute(context.Background(), input)
		assert.Nil(t, output)
		assert.Error(t, err, "quantity must be greater than zero")

		input.Quantity = 1
		output, err = u.Execute(context.Background(), input)
		assert.Nil(t, output)
		assert.Error(t, err, "price must be greater than zero")
	})

	t.Run("Fail if repository fail", func(t *testing.T) {
		repository := &mocks.FruitRepositoryMock{}
		repository.On("Save", mock.Anything, mock.Anything).Return(errors.New("repository save fail"))
		u := usecase.NewCreateFruitUseCase(repository)

		output, err := u.Execute(context.Background(), &usecase.CreateFruitUseCaseInputDTO{
			Name:     "name",
			Owner:    "owner",
			Price:    10.0,
			Quantity: 1,
		})

		fmt.Println(err)
		assert.Nil(t, output)
		assert.EqualError(t, err, "repository save fail")
	})

	t.Run("With Valid Params", func(t *testing.T) {
		repository := &mocks.FruitRepositoryMock{}
		repository.On("Save", mock.Anything, mock.Anything).Return(nil)

		u := usecase.NewCreateFruitUseCase(repository)

		input := &usecase.CreateFruitUseCaseInputDTO{
			Name:     "Name",
			Owner:    "Owner",
			Quantity: 1,
			Price:    100.0,
		}

		output, err := u.Execute(context.Background(), input)

		repository.AssertNumberOfCalls(t, "Save", 1)

		assert.Nil(t, err)

		assert.Nil(t, err)
		assert.NotNil(t, output.ID)
		assert.NotNil(t, output.CreatedAt)
		assert.NotNil(t, output.UpdatedAt)
		assert.Equal(t, output.Name, "Name")
		assert.Equal(t, output.Owner, "Owner")
		assert.Equal(t, output.Quantity, 1)
		assert.Equal(t, output.Price, 100.0)
		assert.Equal(t, output.Status, "comestible")
	})
}
