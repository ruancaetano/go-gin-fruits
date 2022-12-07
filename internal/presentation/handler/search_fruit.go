package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/usecase"
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

// MakeSearchFruitHandler generate handler function to http search fruit request
// @Summary      Search fruits
// @Description  Search fruits by name and status
// @Tags         fruits
// @Accept       json
// @Produce      json
// @Param		 name query string true "Fruit name"
// @Param		 status query string true "Fruit status"
// @Param		 offset query int false "Pagination offset" 1
// @Param		 limit query int false "Pagination limit" 100
// @Success		 200 {object} SearchFruitResponseResult
// @Failure		 400 {object} error.HttpError
// @Failure		 500 {object} error.HttpError
// @Router       /fruits/search [get]
func MakeSearchFruitHandler(u protocol.UseCase[*usecase.SearchFruitUseCaseInputDTO, *usecase.SearchFruitUseCaseOutputDTO]) gin.HandlerFunc {
	return func(c *gin.Context) {

		name := c.Query("name")
		status := c.Query("status")

		var offset int64
		var limit int64
		var err error

		if offset, err = strconv.ParseInt(c.Query("offset"), 10, 64); err != nil {
			offset = 0
		}

		if limit, err = strconv.ParseInt(c.Query("limit"), 10, 64); err != nil {
			limit = 0
		}

		input := &usecase.SearchFruitUseCaseInputDTO{
			Name:   name,
			Status: status,
			Offset: int(offset),
			Limit:  int(limit),
		}

		output, err := u.Execute(c.Request.Context(), input)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"status":  http.StatusBadRequest,
			})
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

		c.JSON(http.StatusOK, response)
	}
}
