package router

import (
	"github.com/abhinav812/cloudy-bookstore/internal/app"
	"github.com/go-chi/chi"
)

// New - New create a new chi router, with index uri handler.
func New() *chi.Mux {
	r := chi.NewRouter()

	r.MethodFunc("GET", "/", app.HandleIndex)

	return r
}
