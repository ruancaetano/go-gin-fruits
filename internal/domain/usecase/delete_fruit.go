package usecase

import (
	"context"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"time"
)

type DeleteFruitUseCase struct {
	repository protocol.FruitRepository
}

type DeleteFruitUseCaseInputDTO struct {
	ID string
}

type DeleteFruitUseCaseOutputDTO struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Owner     string
	Quantity  int
	Price     float64
	Status    string
}

func NewDeleteFruitUseCase(r protocol.FruitRepository) protocol.UseCase[*DeleteFruitUseCaseInputDTO, *DeleteFruitUseCaseOutputDTO] {
	return &DeleteFruitUseCase{
		r,
	}
}

func (dfu DeleteFruitUseCase) Execute(ctx context.Context, input *DeleteFruitUseCaseInputDTO) (*DeleteFruitUseCaseOutputDTO, error) {

	fruit, err := dfu.repository.Get(ctx, input.ID)

	if err != nil {
		return nil, err
	}

	fruit.Status = "podrido"
	fruit.UpdatedAt = time.Now()

	err = dfu.repository.Save(ctx, fruit)
	if err != nil {
		return nil, err
	}

	return &DeleteFruitUseCaseOutputDTO{
		ID:        fruit.ID,
		CreatedAt: fruit.CreatedAt,
		UpdatedAt: fruit.UpdatedAt,
		Name:      fruit.Name,
		Owner:     fruit.Owner,
		Quantity:  fruit.Quantity,
		Price:     fruit.Price,
		Status:    fruit.Status,
	}, nil
}
