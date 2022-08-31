package controller

import (
	"crm/service"
	"crm/store"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Customer is a controller for customer
type Customer struct {
	store *store.Customer
}

// NewCustomer creates a new customer controller
func NewCustomer(store *store.Customer) *Customer {
	return &Customer{
		store: store,
	}
}

// Get responses with all customer from store
func (c *Customer) Get(w http.ResponseWriter, r *http.Request) {
	Response{
		HTTPSTatus: 200,
		Status:     StatusSuccess,
		Data:       c.store.List(),
	}.WriteJson(w)
}

// GetSingle responses with a single customer for a specific id from store
func (c *Customer) GetSingle(w http.ResponseWriter, r *http.Request) {

	// parsing id from url parameter
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		Response{
			HTTPSTatus: http.StatusNotFound,
			Status:     StatusFailed,
			Message:    "customer not found",
		}.WriteJson(w)
		return
	}

	customer, err := c.store.FindById(id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			Response{
				HTTPSTatus: http.StatusNotFound,
				Status:     StatusFailed,
				Message:    "customer not found",
			}.WriteJson(w)
			return
		}
		Response{
			HTTPSTatus: http.StatusInternalServerError,
			Status:     StatusFailed,
			Message:    err.Error(),
		}.WriteJson(w)
		return
	}

	Response{
		HTTPSTatus: http.StatusOK,
		Status:     "success",
		Data:       customer,
	}.WriteJson(w)
}

// Add creates a new customer in store
func (c *Customer) Add(w http.ResponseWriter, r *http.Request) {

	// Getting input from request body
	var customer service.CustomerInput
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		Response{
			HTTPSTatus: http.StatusBadRequest,
			Status:     StatusFailed,
			Message:    "failed to get input body. error:" + err.Error(),
		}.WriteJson(w)
		return
	}

	// input field validation
	if err := customer.Validate(); err != nil {
		Response{
			HTTPSTatus: http.StatusUnprocessableEntity,
			Status:     StatusFailed,
			Message:    "failed to validate customer input.",
			Data:       err,
		}.WriteJson(w)
		return
	}

	c.store.Create(store.CustomerInput{
		Name:        customer.Name,
		Role:        customer.Role,
		Email:       customer.Email,
		Phone:       customer.Phone,
		IsContacted: customer.IsContacted,
	})

	Response{
		HTTPSTatus: 201,
		Status:     "success",
	}.WriteJson(w)
}

// Update updates an exisiting customer in store
func (c *Customer) Update(w http.ResponseWriter, r *http.Request) {
	// parsing id from url parameter
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		Response{
			HTTPSTatus: http.StatusNotFound,
			Status:     StatusFailed,
			Message:    "customer not found",
		}.WriteJson(w)
		return
	}

	// Getting input from request body
	var customer service.CustomerInput
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		Response{
			HTTPSTatus: http.StatusBadRequest,
			Status:     StatusFailed,
			Message:    "failed to get input body. error:" + err.Error(),
		}.WriteJson(w)
		return
	}

	// validating input
	if err := customer.Validate(); err != nil {
		Response{
			HTTPSTatus: http.StatusUnprocessableEntity,
			Status:     StatusFailed,
			Message:    "failed to validate customer input.",
			Data:       err,
		}.WriteJson(w)
		return
	}

	err = c.store.Update(id, store.CustomerInput{
		Name:        customer.Name,
		Role:        customer.Role,
		Email:       customer.Email,
		Phone:       customer.Phone,
		IsContacted: customer.IsContacted,
	})
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			Response{
				HTTPSTatus: http.StatusNotFound,
				Status:     StatusFailed,
				Message:    "customer not found",
			}.WriteJson(w)
			return
		}
		Response{
			HTTPSTatus: http.StatusInternalServerError,
			Status:     StatusFailed,
			Message:    err.Error(),
		}.WriteJson(w)
		return
	}

	Response{
		HTTPSTatus: 200,
		Status:     "success",
	}.WriteJson(w)

}

// Delete deletes an existing customer from store
func (c *Customer) Delete(w http.ResponseWriter, r *http.Request) {

	// parsing id from url parameter
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		Response{
			HTTPSTatus: http.StatusNotFound,
			Status:     StatusFailed,
			Message:    "customer not found",
		}.WriteJson(w)
		return
	}

	err = c.store.Delete(id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			Response{
				HTTPSTatus: http.StatusNotFound,
				Status:     StatusFailed,
				Message:    "customer not found",
			}.WriteJson(w)
			return
		}
		Response{
			HTTPSTatus: http.StatusInternalServerError,
			Status:     StatusFailed,
			Message:    err.Error(),
		}.WriteJson(w)
		return
	}

	Response{
		HTTPSTatus: 200,
		Status:     "success",
	}.WriteJson(w)
}
