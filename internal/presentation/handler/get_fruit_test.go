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
	"testing"
	"time"
)

type GetFruitUseCaseMock struct {
	mock.Mock
}

func (c *GetFruitUseCaseMock) Execute(ctx context.Context, i *usecase.GetFruitUseCaseInputDTO) (*usecase.GetFruitUseCaseOutputDTO, error) {
	args := c.Called(ctx, i)
	return args.Get(0).(*usecase.GetFruitUseCaseOutputDTO), args.Error(1)
}

func TestGetFruitHandler(t *testing.T) {
	t.Run("With invalid param", func(t *testing.T) {
		u := &GetFruitUseCaseMock{}
		h := handler.MakeGetFruitHandler(u)

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		ctx.Params = []gin.Param{
			{Key: "id", Value: ""},
		}

		r := httptest.NewRequest("GET", "/fruits/{id}", nil)
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

	t.Run("With usecase fail", func(t *testing.T) {
		u := &GetFruitUseCaseMock{}
		u.On("Execute", mock.Anything, mock.Anything).Return(&usecase.GetFruitUseCaseOutputDTO{}, errors.New("item not found"))

		h := handler.MakeGetFruitHandler(u)

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "some-uuid",
			},
		}

		r := httptest.NewRequest("GET", "/fruits/{id}", nil)
		ctx.Request = r

		var response error2.HttpError
		h(ctx)
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

		u := &GetFruitUseCaseMock{}
		u.On("Execute", mock.Anything, mock.Anything).Return(&usecase.GetFruitUseCaseOutputDTO{
			ID:        fruitMock.ID,
			Name:      fruitMock.Name,
			Owner:     fruitMock.Owner,
			Price:     fruitMock.Price,
			Quantity:  fruitMock.Quantity,
			CreatedAt: fruitMock.CreatedAt,
			UpdatedAt: fruitMock.UpdatedAt,
			Status:    fruitMock.Status,
		}, nil)

		h := handler.MakeGetFruitHandler(u)

		rr := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(rr)
		ctx.Params = []gin.Param{
			{
				Key:   "id",
				Value: "some-uuid",
			},
		}

		r := httptest.NewRequest("GET", "/fruits/{id}", nil)
		ctx.Request = r

		var response handler.GetFruitResponseDTO
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
		assert.True(t, response.CreatedAt.Equal(fruitMock.CreatedAt))
		assert.True(t, response.UpdatedAt.Equal(fruitMock.UpdatedAt))
		assert.Equal(t, response.Price, fruitMock.Price)
		assert.Equal(t, response.Quantity, fruitMock.Quantity)
	})
}
