package usecase

import (
	"context"
	"errors"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"time"
)

type SearchFruitUseCase struct {
	repository protocol.FruitRepository
}

type SearchFruitUseCaseInputDTO struct {
	Name   string
	Status string
	Offset int
	Limit  int
}

type SearchFruitUseCaseOutputPaging struct {
	Total  int
	Limit  int
	Offset int
}

type SearchFruitUseCaseOutputResult struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Owner     string
	Quantity  int
	Price     float64
	Status    string
}

type SearchFruitUseCaseOutputDTO struct {
	Paging  *SearchFruitUseCaseOutputPaging
	Results []*SearchFruitUseCaseOutputResult
}

func NewSearchFruitUseCase(r protocol.FruitRepository) protocol.UseCase[*SearchFruitUseCaseInputDTO, *SearchFruitUseCaseOutputDTO] {
	return &SearchFruitUseCase{
		r,
	}
}

func (sfu *SearchFruitUseCase) Execute(ctx context.Context, input *SearchFruitUseCaseInputDTO) (*SearchFruitUseCaseOutputDTO, error) {
	err := sfu.validateInput(input)

	if err != nil {
		return nil, err
	}

	filter := &protocol.FruitSearchFilter{
		Name:   input.Name,
		Status: input.Status,
	}

	result, err := sfu.repository.Search(ctx, filter, input.Offset, input.Limit)

	if err != nil {
		return nil, err
	}

	var mappedResult []*SearchFruitUseCaseOutputResult

	for _, r := range result.Results {
		mappedResult = append(mappedResult, &SearchFruitUseCaseOutputResult{
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

	return &SearchFruitUseCaseOutputDTO{
		Paging: &SearchFruitUseCaseOutputPaging{
			Total:  result.Paging.Total,
			Offset: result.Paging.Offset,
			Limit:  result.Paging.Limit,
		},
		Results: mappedResult,
	}, nil
}

func (sfu *SearchFruitUseCase) validateInput(i *SearchFruitUseCaseInputDTO) error {
	if i.Name == "" {
		return errors.New("name is required")
	}

	if i.Status == "" {
		return errors.New("status is required")
	}

	if i.Offset <= 0 {
		return errors.New("offset must be greater than 0")
	}

	if i.Limit < 1 || i.Limit > 100 {
		return errors.New("limit must be a number between 1 and 100")
	}

	return nil
}
