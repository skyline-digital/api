package models

type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"username"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}
