package router

import (
	"github.com/abhinav812/cloudy-bookstore/internal/app"
	"github.com/abhinav812/cloudy-bookstore/internal/handler"
	"github.com/go-chi/chi"
)

// New - New create a new chi router, with index uri handler.
func New(server *app.Server) *chi.Mux {
	l := server.Logger()

	r := chi.NewRouter()

	r.Method("GET", "/", handler.NewHandler(server.HandleIndex, l))

	return r
}
