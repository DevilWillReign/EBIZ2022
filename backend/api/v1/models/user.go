package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
