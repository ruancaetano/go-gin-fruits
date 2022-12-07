package handler

import (
	"encoding/json"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/usecase"
	error2 "github.com/ruancaetano/go-gin-fruits/internal/presentation/error"
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

func MakeCreateFruitHandler(u protocol.UseCase[*usecase.CreateFruitUseCaseInputDTO, *usecase.CreateFruitUseCaseOutputDTO]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body := &CreateFruitRequestDTO{}
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

		owner := r.Header.Get("x-owner")

		input := &usecase.CreateFruitUseCaseInputDTO{
			Name:     body.Name,
			Price:    body.Price,
			Quantity: body.Quantity,
			Owner:    owner,
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

		w.WriteHeader(http.StatusCreated)
		responseJson, _ := json.Marshal(response)
		w.Write(responseJson)
	}
}
