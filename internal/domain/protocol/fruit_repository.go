package protocol

import (
	"context"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/entity"
)

type FruitSearchFilter struct {
	Name   string
	Status string
}

type FruitSearchResultPaging struct {
	Total  int
	Limit  int
	Offset int
}

type FruitSearchResult struct {
	Paging  *FruitSearchResultPaging
	Results []*entity.Fruit
}

type FruitRepository interface {
	Save(context context.Context, fruit *entity.Fruit) error
	Get(context context.Context, id string) (*entity.Fruit, error)
	Search(context context.Context, filter *FruitSearchFilter, offset int, limit int) (*FruitSearchResult, error)
}
