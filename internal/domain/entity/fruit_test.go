package entity_test

import (
	"github.com/ruancaetano/go-gin-fruits/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFruit(t *testing.T) {

	t.Run("With invalid params", func(t *testing.T) {
		fruit, err := entity.NewFruit("", "owner", 1, 10)
		assert.Nil(t, fruit)
		assert.Error(t, err, "name is required")

		fruit, err = entity.NewFruit("123123@23123", "owner", 1, 10)
		assert.Nil(t, fruit)
		assert.Error(t, err, "name cannot contain numbers or special characters")

		fruit, err = entity.NewFruit("name", "", 1, 10)
		assert.Nil(t, fruit)
		assert.Error(t, err, "owner is required")

		fruit, err = entity.NewFruit("name", "owner", 0, 10)
		assert.Nil(t, fruit)
		assert.Error(t, err, "quantity must be greater than zero")

		fruit, err = entity.NewFruit("name", "owner", 1, 0)
		assert.Nil(t, fruit)
		assert.Error(t, err, "price must be greater than zero")
	})

	t.Run("With valid params", func(t *testing.T) {
		fruit, err := entity.NewFruit("Name", "Owner", 1, 10)

		assert.Nil(t, err)
		assert.NotNil(t, fruit.ID)
		assert.NotNil(t, fruit.CreatedAt)
		assert.NotNil(t, fruit.UpdatedAt)
		assert.Equal(t, fruit.Name, "Name")
		assert.Equal(t, fruit.Owner, "Owner")
		assert.Equal(t, fruit.Quantity, 1)
		assert.Equal(t, fruit.Price, 10.0)
		assert.Equal(t, fruit.Status, "comestible")
	})
}
