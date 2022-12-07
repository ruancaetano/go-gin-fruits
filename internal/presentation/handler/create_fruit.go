package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/usecase"
	"net/http"
	"time"
)

type CreateFruitRequestDTO struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type CreateFruitResponseDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"date_created"`
	UpdatedAt time.Time `json:"date_last_updated"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Owner     string    `json:"owner"`
	Status    string    `json:"status"`
}

func MakeCreateFruitHandler(u protocol.UseCase[*usecase.CreateFruitUseCaseInputDTO, *usecase.CreateFruitUseCaseOutputDTO]) gin.HandlerFunc {
	return func(c *gin.Context) {

		body := &CreateFruitRequestDTO{}
		err := c.BindJSON(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid request body",
				"status":  http.StatusBadRequest,
			})
			return
		}

		owner := c.GetHeader("x-owner")

		input := &usecase.CreateFruitUseCaseInputDTO{
			Name:     body.Name,
			Price:    body.Price,
			Quantity: body.Quantity,
			Owner:    owner,
		}

		output, err := u.Execute(c.Request.Context(), input)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"status":  http.StatusBadRequest,
			})
			return
		}

		response := &CreateFruitResponseDTO{
			ID:        output.ID,
			CreatedAt: output.CreatedAt,
			UpdatedAt: output.UpdatedAt,
			Name:      output.Name,
			Status:    output.Status,
			Owner:     output.Owner,
			Price:     output.Price,
			Quantity:  output.Quantity,
		}
		c.JSON(http.StatusCreated, response)
	}
}
