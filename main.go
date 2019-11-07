package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/damascus-mx/kepler/user/models"
)

func main() {
	fmt.Print("Running new HTTP Microservice")
	app := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	app.Use(cors.Handler)

	app.Get("/user", userHandler)

	http.ListenAndServe(": 3000", app)
}

func customMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "123")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{}
	user.Username = "elRuelasLGBT"

	userJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
