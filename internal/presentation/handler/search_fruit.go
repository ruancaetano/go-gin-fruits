package handler

import (
	"encoding/json"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/usecase"
	error2 "github.com/ruancaetano/go-gin-fruits/internal/presentation/error"
	"net/http"
	"strconv"
	"time"
)

type SearchFruitResponsePaging struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type SearchFruitResponseResult struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"date_created"`
	UpdatedAt time.Time `json:"date_last_updated"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Owner     string    `json:"owner"`
	Status    string    `json:"status"`
}

type SearchFruitResponseDTO struct {
	Paging  *SearchFruitResponsePaging
	Results []*SearchFruitResponseResult
}

func MakeSearchFruitHandler(u protocol.UseCase[*usecase.SearchFruitUseCaseInputDTO, *usecase.SearchFruitUseCaseOutputDTO]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		name := query.Get("name")
		status := query.Get("status")

		var offset int64
		var limit int64
		var err error

		if offset, err = strconv.ParseInt(query.Get("offset"), 10, 64); err != nil {
			offset = 0
		}

		if limit, err = strconv.ParseInt(query.Get("limit"), 10, 64); err != nil {
			limit = 0
		}

		input := &usecase.SearchFruitUseCaseInputDTO{
			Name:   name,
			Status: status,
			Offset: int(offset),
			Limit:  int(limit),
		}

		output, err := u.Execute(r.Context(), input)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp, _ := json.Marshal(error2.HttpError{
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			})
			w.Write(resp)
			return
		}

		var mappedResults []*SearchFruitResponseResult
		for _, r := range output.Results {
			mappedResults = append(mappedResults, &SearchFruitResponseResult{
				ID:        r.ID,
				CreatedAt: r.CreatedAt,
				UpdatedAt: r.UpdatedAt,
				Name:      r.Name,
				Owner:     r.Owner,
				Quantity:  r.Quantity,
				Price:     r.Price,
				Status:    r.Status,
			})
		}

		response := &SearchFruitResponseDTO{
			Paging: &SearchFruitResponsePaging{
				Total:  output.Paging.Total,
				Offset: output.Paging.Offset,
				Limit:  output.Paging.Limit,
			},
			Results: mappedResults,
		}
		w.WriteHeader(http.StatusOK)
		responseJson, _ := json.Marshal(response)
		w.Write(responseJson)
	}
}
