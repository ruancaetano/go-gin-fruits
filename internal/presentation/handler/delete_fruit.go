package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/usecase"
	"net/http"
	"time"
)

type DeleteFruitResponseDTO struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"date_created"`
	UpdatedAt time.Time `json:"date_last_updated"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Owner     string    `json:"owner"`
	Status    string    `json:"status"`
}

// MakeDeleteFruitHandler generate handler function to http delete fruit request
// @Summary      Delete a fruit
// @Description  Update fruit status to podrido
// @Tags         fruits
// @Accept       json
// @Produce      json
// @Param		 id path string true "Fruit id"
// @Success		 200 {object} DeleteFruitResponseDTO
// @Failure		 400 {object} error.HttpError
// @Failure		 500 {object} error.HttpError
// @Router       /fruits/{id} [delete]
func MakeDeleteFruitHandler(u protocol.UseCase[*usecase.DeleteFruitUseCaseInputDTO, *usecase.DeleteFruitUseCaseOutputDTO]) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid request param",
				"status":  http.StatusBadRequest,
			})
			return
		}

		input := &usecase.DeleteFruitUseCaseInputDTO{
			ID: id,
		}

		output, err := u.Execute(c.Request.Context(), input)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
				"status":  http.StatusBadRequest,
			})
			return
		}

		response := &DeleteFruitResponseDTO{
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
