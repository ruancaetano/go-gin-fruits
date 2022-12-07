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
	"strings"
	"testing"
)

type CreateFruitUseCaseMock struct {
	mock.Mock
}

func (c *CreateFruitUseCaseMock) Execute(ctx context.Context, i *usecase.CreateFruitUseCaseInputDTO) (*usecase.CreateFruitUseCaseOutputDTO, error) {
	args := c.Called(ctx, i)

	return args.Get(0).(*usecase.CreateFruitUseCaseOutputDTO), args.Error(1)
}

func TestCraeteFruitHandler(t *testing.T) {
	t.Run("With empty body", func(t *testing.T) {
		u := &CreateFruitUseCaseMock{}
		h := handler.MakeCreateFruitHandler(u)

		r := httptest.NewRequest("POST", "/fruits", nil)
		rr := httptest.NewRecorder()

		var response error2.HttpError
		h(rr, r)
		err := json.Unmarshal([]byte(rr.Body.String()), &response)
		if err != nil {
			t.Error("Parse JSON Data Error")
		}

		assert.Nil(t, err)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
		assert.Equal(t, response.Message, "invalid request body")
	})

	t.Run("When usecase fails", func(t *testing.T) {
		u := &CreateFruitUseCaseMock{}
		u.On("Execute", mock.Anything, mock.Anything).Return(&usecase.CreateFruitUseCaseOutputDTO{}, errors.New("name is required"))
		h := handler.MakeCreateFruitHandler(u)
		body := `{"name": "", "quantity": 1, "price": 10.10}`
		r := httptest.NewRequest("POST", "/fruits", strings.NewReader(body))
		rr := httptest.NewRecorder()

		var response error2.HttpError
		h(rr, r)
		err := json.Unmarshal([]byte(rr.Body.String()), &response)
		if err != nil {
			t.Error("Parse JSON Data Error")
		}

		assert.Nil(t, err)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
		assert.Equal(t, response.Message, "name is required")
	})

	t.Run("Success", func(t *testing.T) {
		fruitMock, _ := entity.NewFruit("uva", "owner", 1, 10.10)

		u := &CreateFruitUseCaseMock{}
		u.On("Execute", mock.Anything, mock.Anything).Return(&usecase.CreateFruitUseCaseOutputDTO{
			ID:        fruitMock.ID,
			CreatedAt: fruitMock.CreatedAt,
			UpdatedAt: fruitMock.UpdatedAt,
			Name:      fruitMock.Name,
			Quantity:  fruitMock.Quantity,
			Price:     fruitMock.Price,
			Status:    fruitMock.Status,
			Owner:     fruitMock.Owner,
		}, nil)
		h := handler.MakeCreateFruitHandler(u)
		body := `{"name": "uva", "quantity": 1, "price": 10.10}`
		r := httptest.NewRequest("POST", "/fruits", strings.NewReader(body))
		r.Header.Set("x-owner", "owner")
		rr := httptest.NewRecorder()

		var response handler.CreateFruitResponseDTO
		h(rr, r)
		err := json.Unmarshal([]byte(rr.Body.String()), &response)
		if err != nil {
			t.Error("Parse JSON Data Error")
		}

		assert.Nil(t, err)
		assert.Equal(t, rr.Code, http.StatusCreated)
		assert.NotNil(t, response.ID)
		assert.Equal(t, response.Name, fruitMock.Name)
		assert.Equal(t, response.Owner, fruitMock.Owner)
		assert.Equal(t, response.Price, fruitMock.Price)
		assert.Equal(t, response.Quantity, fruitMock.Quantity)
		assert.Equal(t, response.Status, fruitMock.Status)
		assert.True(t, response.CreatedAt.Equal(fruitMock.CreatedAt))
		assert.True(t, response.UpdatedAt.Equal(fruitMock.UpdatedAt))
	})
}
