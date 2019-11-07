package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	_userModels "github.com/damascus-mx/kepler/user/models"
)

// CreateUser creates a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello from user!")
}

// GetUser Gets a user
func GetUser(w http.ResponseWriter, r *http.Request) {
	user := _userModels.User{}
	user.Username = "elRuelasLGBT"

	userJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}
