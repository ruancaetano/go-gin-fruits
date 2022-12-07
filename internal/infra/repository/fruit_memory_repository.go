package repository

import (
	"context"
	"errors"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/entity"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"strings"
)

type FruitMemoryRepository struct {
	fruits []*entity.Fruit
}

func NewFruitMemoryRepository() *FruitMemoryRepository {
	return &FruitMemoryRepository{
		fruits: []*entity.Fruit{},
	}
}

func (fmr *FruitMemoryRepository) Save(_ context.Context, fruit *entity.Fruit) error {
	for index, f := range fmr.fruits {
		if f.ID == fruit.ID {
			fmr.fruits[index] = fruit
			return nil
		}
	}
	fmr.fruits = append(fmr.fruits, fruit)

	return nil
}

func (fmr *FruitMemoryRepository) Get(_ context.Context, id string) (*entity.Fruit, error) {
	for _, f := range fmr.fruits {
		if f.ID == id {
			return f, nil
		}
	}

	return nil, errors.New("fruit not found")
}

func (fmr *FruitMemoryRepository) Search(_ context.Context, filter *protocol.FruitSearchFilter, offset int, limit int) (*protocol.FruitSearchResult, error) {
	var founds []*entity.Fruit
	var results []*entity.Fruit

	for _, f := range fmr.fruits {
		if strings.Contains(strings.ToLower(f.Name), strings.ToLower(filter.Name)) && f.Status == filter.Status {
			founds = append(founds, f)
		}
	}

	if len(founds) > 0 {
		start := (offset - 1) * limit
		end := offset * limit

		if start > len(founds) {
			start = len(founds)
		}

		if end > len(founds) {
			end = len(founds)
		}

		results = founds[start:end]
	}

	return &protocol.FruitSearchResult{
		Paging: &protocol.FruitSearchResultPaging{
			Total:  len(founds),
			Offset: offset,
			Limit:  limit,
		},
		Results: results,
	}, nil
}
