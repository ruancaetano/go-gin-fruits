package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/entity"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/usecase"
	error2 "github.com/ruancaetano/go-gin-fruits/internal/presentation/error"
	"github.com/ruancaetano/go-gin-fruits/internal/presentation/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type DeleteFruitUseCaseMock struct {
	mock.Mock
}

func (c *DeleteFruitUseCaseMock) Execute(ctx context.Context, i *usecase.DeleteFruitUseCaseInputDTO) (*usecase.DeleteFruitUseCaseOutputDTO, error) {
	args := c.Called(ctx, i)
	return args.Get(0).(*usecase.DeleteFruitUseCaseOutputDTO), args.Error(1)
}

func TestDeleteFruitHandler(t *testing.T) {
	t.Run("With invalid param", func(t *testing.T) {
		u := &DeleteFruitUseCaseMock{}
		h := handler.MakeDeleteFruitHandler(u)

		r := httptest.NewRequest("DELETE", "/fruits/{id}", nil)
		rr := httptest.NewRecorder()

		r = r.WithContext(context.WithValue(r.Context(), "id", ""))

		var response error2.HttpError
		h(rr, r)
		err := json.Unmarshal([]byte(rr.Body.String()), &response)
		if err != nil {
			t.Error("Parse JSON Data Error")
		}

		assert.Nil(t, err)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
		assert.Equal(t, response.Message, "invalid request param")
	})

	t.Run("With usecase fail", func(t *testing.T) {
		u := &DeleteFruitUseCaseMock{}
		u.On("Execute", mock.Anything, mock.Anything).Return(&usecase.DeleteFruitUseCaseOutputDTO{}, errors.New("item not found"))

		h := handler.MakeDeleteFruitHandler(u)

		r := httptest.NewRequest("DELETE", "/fruits/{id}", nil)
		rr := httptest.NewRecorder()

		r = r.WithContext(context.WithValue(r.Context(), "id", "some-uuid"))

		var response error2.HttpError
		h(rr, r)
		err := json.Unmarshal([]byte(rr.Body.String()), &response)
		if err != nil {
			t.Error("Parse JSON Data Error")
		}

		assert.Nil(t, err)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
		assert.Equal(t, response.Message, "item not found")

	})

	t.Run("With usecase success", func(t *testing.T) {
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

		u := &DeleteFruitUseCaseMock{}
		u.On("Execute", mock.Anything, mock.Anything).Return(&usecase.DeleteFruitUseCaseOutputDTO{
			ID:        fruitMock.ID,
			Name:      fruitMock.Name,
			Owner:     fruitMock.Owner,
			Price:     fruitMock.Price,
			Quantity:  fruitMock.Quantity,
			CreatedAt: fruitMock.CreatedAt,
			UpdatedAt: fruitMock.UpdatedAt.Add(time.Second),
			Status:    "podrido",
		}, nil)

		h := handler.MakeDeleteFruitHandler(u)

		r := httptest.NewRequest("DELETE", "/fruits/{id}", nil)
		rr := httptest.NewRecorder()

		r = r.WithContext(context.WithValue(r.Context(), "id", fruitMock.ID))

		var response handler.DeleteFruitResponseDTO
		h(rr, r)
		err := json.Unmarshal([]byte(rr.Body.String()), &response)
		if err != nil {
			t.Error("Parse JSON Data Error")
		}

		assert.Nil(t, err)
		assert.Equal(t, rr.Code, http.StatusOK)
		assert.Equal(t, response.ID, fruitMock.ID)
		assert.Equal(t, response.Name, fruitMock.Name)
		assert.Equal(t, response.Owner, fruitMock.Owner)
		assert.True(t, response.CreatedAt.Equal(fruitMock.CreatedAt))
		assert.False(t, response.UpdatedAt.Equal(fruitMock.UpdatedAt))
		assert.Equal(t, response.Price, fruitMock.Price)
		assert.Equal(t, response.Quantity, fruitMock.Quantity)
		assert.Equal(t, response.Status, "podrido")
	})
}
