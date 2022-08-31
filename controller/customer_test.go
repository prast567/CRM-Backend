package controller_test

import (
	"crm/controller"
	"crm/store"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//  TestGetCustomersHandler tests happy path of submitting a well-formed GET /customers request
func TestGetCustomersHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/customers", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	controller := controller.NewCustomer(store.NewCustomer())
	handler := http.HandlerFunc(controller.Get)
	handler.ServeHTTP(rr, req)

	// Checks for 200 status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Customer.Get returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Checks for JSON response
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("Content-Type does not match: got %v want %v",
			ctype, "application/json")
	}
}

//  TestAddCustomerHandler tests happy path of submitting a well-formed POST /customers request
func TestAddCustomerHandler(t *testing.T) {
	requestBody := strings.NewReader(`
		{
			"name": "Example Name",
			"role": "Example Role",
			"email": "mymail@yahoo.com",
			"phone": "+8801521438989",
			"contacted": true
		}
	`)

	req, err := http.NewRequest("POST", "/customers", requestBody)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	controller := controller.NewCustomer(store.NewCustomer())
	handler := http.HandlerFunc(controller.Add)
	handler.ServeHTTP(rr, req)

	// Checks for 201 status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Customer.Add returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Checks for JSON response
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("Content-Type does not match: got %v want %v",
			ctype, "application/json")
	}
}

// Tests unhappy path of deleting a user that doesn't exist
func TestDeleteCustomerHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/customers/e7847fee-3a0e-455e-b151-519bdb9851c7", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	controller := controller.NewCustomer(store.NewCustomer())
	handler := http.HandlerFunc(controller.Delete)
	handler.ServeHTTP(rr, req)

	// Checks for 404 status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Customer.Delete returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

//  TestGetCustomerHandler Tests unhappy path of getting a user that doesn't exist
func TestGetCustomerHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/customers/e7847fee-3a0e-455e-b151-519bdb9851c7", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	controller := controller.NewCustomer(store.NewCustomer())
	handler := http.HandlerFunc(controller.GetSingle)
	handler.ServeHTTP(rr, req)

	// Checks for 404 status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Customer.GetSingle returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
