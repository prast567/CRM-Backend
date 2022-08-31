package store

import (
	"crm/model"
	"errors"
	"sync"
)

// ErrNotFound is returned when a customer is not found
var ErrNotFound = errors.New("resource not found")

// Customer defines a customer store
type Customer struct {
	mu             sync.Mutex // required to avoid overlapping operations
	customers      map[int]model.Customer
	count          int
	lastInsertedId int
}

// NewCustomer returns a new CustomerStore
func NewCustomer() *Customer {
	return &Customer{
		customers:      map[int]model.Customer{},
		count:          0,
		lastInsertedId: 0,
	}
}

// Len gets the number of record in customer store
func (c *Customer) Len() int {
	return c.count
}

// List retuns all customer in a array
func (c *Customer) List() []model.Customer {
	customers := []model.Customer{}
	for _, customer := range c.customers {
		customers = append(customers, customer)
	}
	return customers
}

// FindById returns a customer in store with the given ID.
// If the customer doesn't exists then it will return ErrNotFound
func (c *Customer) FindById(id int) (*model.Customer, error) {
	if customer, ok := c.customers[id]; ok {
		return &customer, nil
	}

	return nil, ErrNotFound
}

// CustomerInput is used to create and update a customer
type CustomerInput struct {
	Name        string
	Role        string
	Email       string
	Phone       string
	IsContacted bool
}

// Creates add a new customer in the store
func (c *Customer) Create(input CustomerInput) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
	c.lastInsertedId++

	c.customers[c.lastInsertedId] = model.Customer{
		Id:          c.lastInsertedId,
		Name:        input.Name,
		Role:        input.Role,
		Email:       input.Email,
		Phone:       input.Phone,
		IsContacted: input.IsContacted,
	}

}

// Update updates a customer in store.
// If the customer doesn't exists then it will return ErrNotFound
func (c *Customer) Update(id int, input CustomerInput) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.customers[id]; ok {
		c.customers[id] = model.Customer{
			Id:          id,
			Name:        input.Name,
			Role:        input.Role,
			Email:       input.Email,
			Phone:       input.Phone,
			IsContacted: input.IsContacted,
		}
		return nil
	}

	return ErrNotFound
}

// Delete deletes a customer from store.
// If the customer doesn't exists then it will return ErrNotFound
func (c *Customer) Delete(id int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.customers[id]; ok {
		delete(c.customers, id)
		c.count--
		return nil
	}

	return ErrNotFound
}
