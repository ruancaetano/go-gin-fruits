package usecase

import (
	"context"
	"errors"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"time"
)

type UpdateFruitUseCase struct {
	repository protocol.FruitRepository
}

type UpdateFruitUseCaseInputDTO struct {
	ID       string
	Quantity int
	Price    float64
}

type UpdateFruitUseCaseOutputDTO struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Owner     string
	Quantity  int
	Price     float64
	Status    string
}

func NewUpdateFruitUseCase(r protocol.FruitRepository) protocol.UseCase[*UpdateFruitUseCaseInputDTO, *UpdateFruitUseCaseOutputDTO] {
	return &UpdateFruitUseCase{
		repository: r,
	}
}

func (cf *UpdateFruitUseCase) Execute(ctx context.Context, i *UpdateFruitUseCaseInputDTO) (*UpdateFruitUseCaseOutputDTO, error) {
	err := cf.validateInput(i)

	if err != nil {
		return nil, err
	}

	fruit, err := cf.repository.Get(ctx, i.ID)

	if err != nil {
		return nil, err
	}

	fruit.Quantity = i.Quantity
	fruit.Price = i.Price
	fruit.UpdatedAt = time.Now()

	err = cf.repository.Save(ctx, fruit)

	if err != nil {
		return nil, err
	}

	return &UpdateFruitUseCaseOutputDTO{
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

func (*UpdateFruitUseCase) validateInput(i *UpdateFruitUseCaseInputDTO) error {
	if i.ID == "" {
		return errors.New("id is required")
	}

	if i.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	if i.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	return nil
}
