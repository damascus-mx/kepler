package models

// User User model
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
