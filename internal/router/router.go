package router

import (
	"github.com/abhinav812/cloudy-bookstore/internal/app"
	"github.com/abhinav812/cloudy-bookstore/internal/handler"
	"github.com/abhinav812/cloudy-bookstore/internal/router/middleware"
	"github.com/go-chi/chi"
)

// New - New create a new chi router, with index uri handler.
func New(server *app.Server) *chi.Mux {
	l := server.Logger()

	r := chi.NewRouter()

	r.Method("GET", "/", handler.NewHandler(server.HandleIndex, l))

	// route for healthz
	r.Get("/healthz/liveness", app.HandleLive)
	r.Method("GET", "/healthz/readiness", handler.NewHandler(server.HandleReady, l))

	// Route for api's
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(middleware.ContextTypeJSON)

		// Routes for books
		r.Method("GET", "/books", handler.NewHandler(server.HandleListBooks, l))
		r.Method("POST", "/books", handler.NewHandler(server.HandleCreateBook, l))
		r.Method("GET", "/books/{id}", handler.NewHandler(server.HandleReadBook, l))
		r.Method("PUT", "/books/{id}", handler.NewHandler(server.HandleUpdateBook, l))
		r.Method("DELETE", "/books/{id}", handler.NewHandler(server.HandleDeleteBook, l))
	})

	return r
}
