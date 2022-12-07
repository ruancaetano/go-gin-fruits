package handler

import (
	"encoding/json"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/usecase"
	error2 "github.com/ruancaetano/go-gin-fruits/internal/presentation/error"
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

func MakeUpdateFruitHandler(u protocol.UseCase[*usecase.UpdateFruitUseCaseInputDTO, *usecase.UpdateFruitUseCaseOutputDTO]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := r.Context().Value("id").(string)
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			resp, _ := json.Marshal(error2.HttpError{
				Message: "invalid request param",
				Status:  http.StatusBadRequest,
			})
			w.Write(resp)
			return

		}

		body := &UpdateFruitRequestDTO{}
		err := json.NewDecoder(r.Body).Decode(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			resp, _ := json.Marshal(error2.HttpError{
				Message: "invalid request body",
				Status:  http.StatusBadRequest,
			})
			w.Write(resp)
			return
		}

		input := &usecase.UpdateFruitUseCaseInputDTO{
			ID:       id,
			Price:    body.Price,
			Quantity: body.Quantity,
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
		w.WriteHeader(http.StatusOK)
		responseJson, _ := json.Marshal(response)
		w.Write(responseJson)
	}
}
