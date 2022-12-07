package entity

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Fruit struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	Owner     string    `json:"owner"`
	Status    string    `json:"status"`
}

func NewFruit(name string, owner string, quantity int, price float64) (*Fruit, error) {
	fruit := &Fruit{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Owner:     owner,
		Quantity:  quantity,
		Price:     price,
		Status:    "comestible",
	}

	err := fruit.Validate()

	if err != nil {
		return nil, err
	}

	return fruit, nil
}

func (f *Fruit) Validate() error {
	if f.Name == "" {
		return errors.New("name is required")
	}

	valid, err := regexp.MatchString("^[a-zA-Z]+$", f.Name)
	if err != nil || !valid {
		return errors.New("name cannot contain numbers or special characters")
	}

	if f.Owner == "" {
		return errors.New("owner is required")
	}

	if f.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	if f.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	return nil
}
