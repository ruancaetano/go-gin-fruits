package usecase_test

import (
	"context"
	"errors"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/entity"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/usecase"
	"github.com/ruancaetano/go-gin-fruits/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestNewGetFruitUseCase(t *testing.T) {
	r := &mocks.FruitRepositoryMock{}
	u := usecase.NewGetFruitUseCase(r)
	assert.NotNil(t, u)
}

func TestGetFruitUseCase_Execute(t *testing.T) {
	t.Run("With not found fruit", func(t *testing.T) {
		r := &mocks.FruitRepositoryMock{}
		r.On("Get", mock.Anything, mock.Anything).Return(&entity.Fruit{}, errors.New("fruit not found"))
		u := usecase.NewGetFruitUseCase(r)

		output, err := u.Execute(context.Background(), &usecase.GetFruitUseCaseInputDTO{
			ID: "invalid-id",
		})

		assert.Nil(t, output)
		assert.EqualError(t, err, "fruit not found")
	})

	t.Run("With found fruit", func(t *testing.T) {
		fruitMock := &entity.Fruit{
			ID:        "valid-id",
			Name:      "name",
			Owner:     "owner",
			Price:     10.0,
			Quantity:  10,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Status:    "comestible",
		}

		r := &mocks.FruitRepositoryMock{}
		r.On("Get", mock.Anything, mock.Anything).Return(fruitMock, nil)
		u := usecase.NewGetFruitUseCase(r)

		output, err := u.Execute(context.Background(), &usecase.GetFruitUseCaseInputDTO{
			ID: fruitMock.ID,
		})

		assert.Nil(t, err)
		assert.Equal(t, output.ID, fruitMock.ID)
		assert.True(t, output.CreatedAt.Equal(fruitMock.CreatedAt))
		assert.True(t, output.UpdatedAt.Equal(fruitMock.UpdatedAt))
		assert.Equal(t, output.Name, fruitMock.Name)
		assert.Equal(t, output.Owner, fruitMock.Owner)
		assert.Equal(t, output.Quantity, fruitMock.Quantity)
		assert.Equal(t, output.Price, fruitMock.Price)
		assert.Equal(t, output.Status, fruitMock.Status)
	})

}
