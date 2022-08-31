package router

import (
	"crm/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// New creates a new router with neccessary middleware, static endpoints
func New() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)

	r.NotFound(controller.HandleNotFound)
	r.MethodNotAllowed(controller.HandleMethodNotAllowed)

	// serving static files
	fs := http.FileServer(http.Dir("./static/"))
	r.Handle("/", fs)

	return r
}

// SetCustomerRoutes links customer endpoints with customer controller
func SetCustomerRoutes(r *chi.Mux, controller *controller.Customer) {
	r.Route("/customers", func(r chi.Router) {
		r.Get("/", controller.Get)
		r.Get("/{id:[1-9][0-9]*}", controller.GetSingle)

		r.Post("/", controller.Add)

		r.Put("/{id:[1-9][0-9]*}", controller.Update)
		r.Delete("/{id:[1-9][0-9]*}", controller.Delete)
	})
}
