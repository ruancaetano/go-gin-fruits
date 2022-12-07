package usecase_test

import (
	"context"
	"errors"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/entity"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/usecase"
	"github.com/ruancaetano/go-gin-fruits/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewSearchFruitUseCase(t *testing.T) {
	r := &mocks.FruitRepositoryMock{}
	u := usecase.NewSearchFruitUseCase(r)
	assert.NotNil(t, u)
}

func TestSearchFruitUseCase_Execute(t *testing.T) {
	t.Run("With invalid input", func(t *testing.T) {
		r := &mocks.FruitRepositoryMock{}
		u := usecase.NewSearchFruitUseCase(r)

		input := &usecase.SearchFruitUseCaseInputDTO{
			Name:   "",
			Status: "",
			Offset: 0,
			Limit:  0,
		}

		output, err := u.Execute(context.Background(), input)
		assert.Nil(t, output)
		assert.EqualError(t, err, "name is required")

		input.Name = "name"
		output, err = u.Execute(context.Background(), input)
		assert.Nil(t, output)
		assert.EqualError(t, err, "status is required")

		input.Status = "status"
		output, err = u.Execute(context.Background(), input)
		assert.Nil(t, output)
		assert.EqualError(t, err, "offset must be greater than 0")

		input.Offset = 1
		output, err = u.Execute(context.Background(), input)
		assert.Nil(t, output)
		assert.EqualError(t, err, "limit must be a number between 1 and 100")
	})

	t.Run("With search fail", func(t *testing.T) {

		r := &mocks.FruitRepositoryMock{}
		r.On("Search", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&protocol.FruitSearchResult{}, errors.New("search failed"))

		u := usecase.NewSearchFruitUseCase(r)

		input := &usecase.SearchFruitUseCaseInputDTO{
			Name:   "fruit",
			Status: "comestible",
			Offset: 1,
			Limit:  10,
		}

		output, err := u.Execute(context.Background(), input)

		assert.Nil(t, output)
		assert.EqualError(t, err, "search failed")
	})

	t.Run("With valid input", func(t *testing.T) {
		fruit1, err := entity.NewFruit("fruita", "owner", 1, 10.0)
		assert.Nil(t, err)
		fruit2, err := entity.NewFruit("fruitb", "owner", 1, 10.0)
		assert.Nil(t, err)
		fruits := []*entity.Fruit{
			fruit1, fruit2,
		}

		searchResult := &protocol.FruitSearchResult{
			Paging: &protocol.FruitSearchResultPaging{
				Total:  10,
				Limit:  10,
				Offset: 1,
			},
			Results: fruits,
		}

		r := &mocks.FruitRepositoryMock{}
		r.On("Search", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(searchResult, nil)

		u := usecase.NewSearchFruitUseCase(r)

		input := &usecase.SearchFruitUseCaseInputDTO{
			Name:   "fruit",
			Status: "comestible",
			Offset: 1,
			Limit:  10,
		}

		output, err := u.Execute(context.Background(), input)

		assert.Nil(t, err)
		assert.Equal(t, output.Paging.Total, 10)
		assert.Equal(t, output.Paging.Limit, 10)
		assert.Equal(t, output.Paging.Offset, 1)
		assert.Len(t, output.Results, 2)
	})

}
