package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/damascus-mx/kepler/src/models"
	"github.com/go-chi/chi"
)

// UserActions Exports all user actions
type UserActions struct {
}

// Create creates a user
func (a UserActions) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello from user!")
}

// Update Updates a user
func (a UserActions) Update(w http.ResponseWriter, r *http.Request) {
}

// Delete Deletes a user
func (a UserActions) Delete(w http.ResponseWriter, r *http.Request) {
}

// GetAll Gets a user
func (a UserActions) GetAll(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	user.Username = "elRuelasLGBT"

	userJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

// GetByID Gets a user by ID
func (a UserActions) GetByID(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "userID")
	if user == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}

	w.Write([]byte(fmt.Sprintf("user:%s", user)))
}

// Context User HTTP context
func (a UserActions) Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		ctx := context.WithValue(r.Context(), "user", userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
