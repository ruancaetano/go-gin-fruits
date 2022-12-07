package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
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
	"time"
)

type UpdateFruitUseCaseMock struct {
	mock.Mock
}

func (c *UpdateFruitUseCaseMock) Execute(ctx context.Context, i *usecase.UpdateFruitUseCaseInputDTO) (*usecase.UpdateFruitUseCaseOutputDTO, error) {
	args := c.Called(ctx, i)

	return args.Get(0).(*usecase.UpdateFruitUseCaseOutputDTO), args.Error(1)
}

func TestUpdateFruitHandler(t *testing.T) {
	t.Run("With invalid param", func(t *testing.T) {
		u := &UpdateFruitUseCaseMock{}
		h := handler.MakeUpdateFruitHandler(u)

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		ctx.Params = []gin.Param{
			{Key: "id", Value: ""},
		}

		r := httptest.NewRequest("PUT", "/fruits/{id}", nil)
		ctx.Request = r

		var response error2.HttpError
		h(ctx)
		err := json.Unmarshal([]byte(rr.Body.String()), &response)
		if err != nil {
			t.Error("Parse JSON Data Error")
		}

		assert.Nil(t, err)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
		assert.Equal(t, response.Message, "invalid request param")
	})

	t.Run("With empty body", func(t *testing.T) {
		u := &UpdateFruitUseCaseMock{}
		h := handler.MakeUpdateFruitHandler(u)

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		ctx.Params = []gin.Param{
			{Key: "id", Value: "some-uuid"},
		}

		r := httptest.NewRequest("PUT", "/fruits/{id}", nil)
		ctx.Request = r

		var response error2.HttpError
		h(ctx)
		err := json.Unmarshal([]byte(rr.Body.String()), &response)
		if err != nil {
			t.Error("Parse JSON Data Error")
		}

		assert.Nil(t, err)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
		assert.Equal(t, response.Message, "invalid request body")
	})

	t.Run("When usecase fails", func(t *testing.T) {
		u := &UpdateFruitUseCaseMock{}
		u.On("Execute", mock.Anything, mock.Anything).Return(&usecase.UpdateFruitUseCaseOutputDTO{}, errors.New("quantity must be greater than zero"))
		h := handler.MakeUpdateFruitHandler(u)

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		ctx.Params = []gin.Param{
			{Key: "id", Value: "some-uuid"},
		}

		body := `{"quantity": 0, "price": 10.10}`
		r := httptest.NewRequest("PUT", "/fruits/{id}", strings.NewReader(body))
		ctx.Request = r

		var response error2.HttpError
		h(ctx)
		err := json.Unmarshal([]byte(rr.Body.String()), &response)
		if err != nil {
			t.Error("Parse JSON Data Error")
		}

		assert.Nil(t, err)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
		assert.Equal(t, response.Message, "quantity must be greater than zero")
	})

	t.Run("When usecase success", func(t *testing.T) {
		fruitMock, _ := entity.NewFruit("uva", "owner", 1, 1.0)

		u := &UpdateFruitUseCaseMock{}
		u.On("Execute", mock.Anything, mock.Anything).Return(&usecase.UpdateFruitUseCaseOutputDTO{
			ID:        fruitMock.ID,
			CreatedAt: fruitMock.CreatedAt,
			UpdatedAt: fruitMock.CreatedAt.Add(time.Second),
			Name:      fruitMock.Name,
			Quantity:  100,
			Price:     100.0,
			Status:    fruitMock.Status,
			Owner:     fruitMock.Owner,
		}, nil)
		h := handler.MakeUpdateFruitHandler(u)

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		ctx.Params = []gin.Param{
			{Key: "id", Value: "some-uuid"},
		}

		body := `{"quantity": 100, "price": 100.0}`
		r := httptest.NewRequest("PUT", "/fruits/{id}", strings.NewReader(body))
		ctx.Request = r

		var response handler.UpdateFruitResponseDTO
		h(ctx)
		err := json.Unmarshal([]byte(rr.Body.String()), &response)
		if err != nil {
			t.Error("Parse JSON Data Error")
		}

		assert.Nil(t, err)
		assert.Equal(t, rr.Code, http.StatusOK)
		assert.Equal(t, response.ID, fruitMock.ID)
		assert.Equal(t, response.Name, fruitMock.Name)
		assert.Equal(t, response.Owner, fruitMock.Owner)
		assert.Equal(t, response.Price, 100.0)
		assert.Equal(t, response.Quantity, 100)
		assert.Equal(t, response.Status, fruitMock.Status)
		assert.True(t, response.CreatedAt.Equal(fruitMock.CreatedAt))
		assert.False(t, response.UpdatedAt.Equal(fruitMock.UpdatedAt))
	})
}
