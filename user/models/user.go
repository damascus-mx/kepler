package models

// User The User domain model
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
