package main

import (
	"crm/controller"
	"crm/router"
	"crm/store"
	"log"
	"net/http"
)

func main() {
	r := router.New()

	// setting up customer
	st := store.NewCustomer()
	seedCustomer(st)

	customerController := controller.NewCustomer(st)
	router.SetCustomerRoutes(r, customerController)

	// Stating Server
	log.Println("Running on-    http://localhost:3000/")
	log.Fatalln(http.ListenAndServe(":3000", r))
}

func seedCustomer(st *store.Customer) {
	dummy := []store.CustomerInput{
		{
			Name:        "Mr. Robot",
			Role:        "Customer",
			Email:       "robot@gmail.com",
			Phone:       "+8801814567342",
			IsContacted: true,
		},
		{
			Name:        "Mr. John",
			Role:        "Investor",
			Email:       "john@yahoo.com",
			Phone:       "+14155552671",
			IsContacted: false,
		},
		{
			Name:        "James Mark",
			Role:        "Investor",
			Email:       "mark@facebook.com",
			Phone:       "+14155552671",
			IsContacted: true,
		},
	}

	for _, customer := range dummy {
		st.Create(customer)
	}
}
