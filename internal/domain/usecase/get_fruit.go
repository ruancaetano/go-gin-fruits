package usecase

import (
	"context"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"time"
)

type GetFruitUseCase struct {
	repository protocol.FruitRepository
}

type GetFruitUseCaseInputDTO struct {
	ID string
}

type GetFruitUseCaseOutputDTO struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Owner     string
	Quantity  int
	Price     float64
	Status    string
}

func NewGetFruitUseCase(r protocol.FruitRepository) protocol.UseCase[*GetFruitUseCaseInputDTO, *GetFruitUseCaseOutputDTO] {
	return &GetFruitUseCase{
		r,
	}
}

func (g GetFruitUseCase) Execute(ctx context.Context, input *GetFruitUseCaseInputDTO) (*GetFruitUseCaseOutputDTO, error) {

	fruit, err := g.repository.Get(ctx, input.ID)

	if err != nil {
		return nil, err
	}

	return &GetFruitUseCaseOutputDTO{
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
