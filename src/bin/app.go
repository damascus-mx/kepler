package app

import (
	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"github.com/damascus-mx/kepler/src/core"
	"github.com/damascus-mx/kepler/src/routes"
)

// InitApplication Returns a new chi Router pointer
func InitApplication() *chi.Mux {
	// Create a new Chi router
	fmt.Print("Running new REST Microservice")
	app := chi.NewRouter()

	// Set CORS policies
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// Use CORS
	app.Use(cors.Handler)

	// Set default middlewares
	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)

	// Set routes
	var user core.IRoute = routes.UserRoutes{}
	user.SetRoutes(app)

	return app
}
