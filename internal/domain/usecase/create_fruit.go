package usecase

import (
	"context"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/entity"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"time"
)

type CreateFruitUseCase struct {
	repository protocol.FruitRepository
}

type CreateFruitUseCaseInputDTO struct {
	Name     string
	Owner    string
	Quantity int
	Price    float64
}

type CreateFruitUseCaseOutputDTO struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Owner     string
	Quantity  int
	Price     float64
	Status    string
}

func NewCreateFruitUseCase(r protocol.FruitRepository) protocol.UseCase[*CreateFruitUseCaseInputDTO, *CreateFruitUseCaseOutputDTO] {
	return &CreateFruitUseCase{
		repository: r,
	}
}

func (cf *CreateFruitUseCase) Execute(ctx context.Context, i *CreateFruitUseCaseInputDTO) (*CreateFruitUseCaseOutputDTO, error) {
	fruit, err := entity.NewFruit(i.Name, i.Owner, i.Quantity, i.Price)

	if err != nil {
		return nil, err
	}

	err = cf.repository.Save(ctx, fruit)

	if err != nil {
		return nil, err
	}

	return &CreateFruitUseCaseOutputDTO{
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
