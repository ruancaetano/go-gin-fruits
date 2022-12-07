package handler_test

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
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

type SearchFruitUseCaseMock struct {
	mock.Mock
}

func (c *SearchFruitUseCaseMock) Execute(ctx context.Context, i *usecase.SearchFruitUseCaseInputDTO) (*usecase.SearchFruitUseCaseOutputDTO, error) {
	args := c.Called(ctx, i)
	return args.Get(0).(*usecase.SearchFruitUseCaseOutputDTO), args.Error(1)
}

func TestSearchFruitHandler(t *testing.T) {
	t.Run("With usecase fail", func(t *testing.T) {
		u := &SearchFruitUseCaseMock{}
		u.On("Execute", mock.Anything, mock.Anything).Return(&usecase.SearchFruitUseCaseOutputDTO{}, errors.New("name is required"))

		h := handler.MakeSearchFruitHandler(u)

		r := httptest.NewRequest("GET", "/fruits/search", nil)
		q := r.URL.Query()
		q.Add("name", "")
		q.Add("status", "status")
		q.Add("offset", "1")
		q.Add("limit", "100")
		r.URL.RawQuery = q.Encode()

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

	t.Run("With usecase success", func(t *testing.T) {
		var fruitResults []*usecase.SearchFruitUseCaseOutputResult

		for i := 0; i < 2; i++ {
			letters := "abcdefghij"
			fruitResults = append(fruitResults, &usecase.SearchFruitUseCaseOutputResult{
				ID:        uuid.NewString(),
				Name:      "fruit" + string(letters[i]),
				Owner:     "owner",
				Price:     10.0,
				Quantity:  1,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Status:    "status",
			})
		}

		useCaseOutputMock := &usecase.SearchFruitUseCaseOutputDTO{
			Paging: &usecase.SearchFruitUseCaseOutputPaging{
				Total:  2,
				Limit:  100,
				Offset: 1,
			},
			Results: fruitResults,
		}

		u := &SearchFruitUseCaseMock{}
		u.On("Execute", mock.Anything, mock.Anything).Return(useCaseOutputMock, nil)

		h := handler.MakeSearchFruitHandler(u)

		r := httptest.NewRequest("GET", "/fruits/search", nil)
		q := r.URL.Query()
		q.Add("name", "fruit")
		q.Add("status", "status")
		q.Add("offset", "1")
		q.Add("limit", "100")
		r.URL.RawQuery = q.Encode()

		rr := httptest.NewRecorder()

		var response handler.SearchFruitResponseDTO
		h(rr, r)
		err := json.Unmarshal([]byte(rr.Body.String()), &response)
		if err != nil {
			t.Error("Parse JSON Data Error")
		}

		assert.Nil(t, err)
		assert.Equal(t, rr.Code, http.StatusOK)
		assert.Equal(t, response.Paging.Total, 2)
		assert.Equal(t, response.Paging.Limit, 100)
		assert.Equal(t, response.Paging.Offset, 1)
		assert.Len(t, response.Results, 2)
		assert.Equal(t, response.Results[0].ID, fruitResults[0].ID)
		assert.Equal(t, response.Results[1].ID, fruitResults[1].ID)
	})
}
