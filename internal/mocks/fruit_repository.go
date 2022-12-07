package mocks

import (
	"context"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/entity"
	"github.com/ruancaetano/go-gin-fruits/internal/domain/protocol"
	"github.com/stretchr/testify/mock"
)

type FruitRepositoryMock struct {
	mock.Mock
}

func (fr *FruitRepositoryMock) Save(c context.Context, f *entity.Fruit) error {
	args := fr.Called(c, f)
	return args.Error(0)
}

func (fr *FruitRepositoryMock) Get(c context.Context, id string) (*entity.Fruit, error) {
	args := fr.Called(c, id)
	return args.Get(0).(*entity.Fruit), args.Error(1)
}

func (fr *FruitRepositoryMock) Search(ctx context.Context, filter *protocol.FruitSearchFilter, offset int, limit int) (*protocol.FruitSearchResult, error) {
	args := fr.Called(ctx, filter, offset, limit)
	return args.Get(0).(*protocol.FruitSearchResult), args.Error(1)
}
