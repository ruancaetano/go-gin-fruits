package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/usecase"
	"net/http"
	"time"
)

type UpdateFruitRequestDTO struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type UpdateFruitResponseDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"date_created"`
	UpdatedAt time.Time `json:"date_last_updated"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Owner     string    `json:"owner"`
	Status    string    `json:"status"`
}

func MakeUpdateFruitHandler(u protocol.UseCase[*usecase.UpdateFruitUseCaseInputDTO, *usecase.UpdateFruitUseCaseOutputDTO]) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid request param",
				"status":  http.StatusBadRequest,
			})
			return

		}

		body := &UpdateFruitRequestDTO{}
		err := c.BindJSON(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid request body",
				"status":  http.StatusBadRequest,
			})
			return
		}

		input := &usecase.UpdateFruitUseCaseInputDTO{
			ID:       id,
			Price:    body.Price,
			Quantity: body.Quantity,
		}

		output, err := u.Execute(c.Request.Context(), input)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"status":  http.StatusBadRequest,
			})
			return
		}

		response := &UpdateFruitResponseDTO{
			ID:        output.ID,
			CreatedAt: output.CreatedAt,
			UpdatedAt: output.UpdatedAt,
			Name:      output.Name,
			Status:    output.Status,
			Owner:     output.Owner,
			Price:     output.Price,
			Quantity:  output.Quantity,
		}
		c.JSON(http.StatusOK, response)
	}
}
