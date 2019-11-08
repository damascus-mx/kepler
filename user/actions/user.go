package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/damascus-mx/kepler/user/models"
)

// Actions Exports all user actions
type Actions struct {
}

// CreateUser creates a user
func (a Actions) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello from user!")
}

// GetUser Gets a user
func (a Actions) GetUser(w http.ResponseWriter, r *http.Request) {
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
