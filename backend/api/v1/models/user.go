package models

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Admin    bool   `json:"admin" validate:"required"`
}

type UserData struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Admin bool   `json:"admin"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func ConverUserToUserData(user User) UserData {
	return UserData{ID: user.ID, Name: user.Username, Email: user.Email, Admin: user.Admin}
}
