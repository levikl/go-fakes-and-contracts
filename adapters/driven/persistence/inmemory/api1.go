package inmemory

import (
	"context"
	"errors"
	"strconv"

	"github.com/levikl/go-fakes-and-contracts/domain/planner"
)

func NewAPI1() *API1 {
	return &API1{customers: make(map[string]planner.API1Customer)}
}

type API1 struct {
	i         int
	customers map[string]planner.API1Customer
}

var ErrDaveIsForbidden = errors.New("u r banned Dave")

func (a *API1) CreateCustomer(_ context.Context, name string) (planner.API1Customer, error) {
	if name == "Dave" {
		return planner.API1Customer{}, ErrDaveIsForbidden
	}

	newCustomer := planner.API1Customer{
		Name: name,
		ID:   strconv.Itoa(a.i),
	}
	a.customers[newCustomer.ID] = newCustomer
	a.i++
	return newCustomer, nil
}

func (a *API1) GetCustomer(_ context.Context, id string) (planner.API1Customer, error) {
	return a.customers[id], nil
}

func (a *API1) UpdateCustomer(_ context.Context, id, name string) error {
	customer, ok := a.customers[id]
	if !ok {
		return errors.New("oh no")
	}

	customer.Name = name
	a.customers[id] = customer
	return nil
}
