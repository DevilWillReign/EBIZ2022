package services

import (
	"apprit/store/api/v1/database/daos"
	"apprit/store/api/v1/database/dtos"
	"apprit/store/api/v1/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) ([]models.User, error) {
	userDTOs, err := daos.GetUsers(db)
	if err != nil {
		return []models.User{}, nil
	}
	var users []models.User
	for _, userDTO := range userDTOs {
		users = append(users, copyUserProperties(userDTO))
	}
	return users, nil
}

func GetUserById(db *gorm.DB, id uint64) (models.User, error) {
	userDTO, err := daos.GetUserById(db, id)
	if err != nil {
		return models.User{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return copyUserProperties(userDTO), nil
}

func DeleteUserById(db *gorm.DB, id uint64) error {
	err := daos.DeleteUserById(db, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return nil
}

func AddUser(db *gorm.DB, user models.User) error {
	err := daos.AddUser(db, copyUserDTOProperties(user))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func ReplaceUser(db *gorm.DB, user models.User) error {
	err := daos.ReplaceUser(db, uint64(user.ID), copyUserDTOProperties(user))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func copyUserProperties(userDTO dtos.UserDTO) models.User {
	return models.User{ID: userDTO.ID, Username: userDTO.Username, Email: userDTO.Email, Password: userDTO.Password}
}

func copyUserDTOProperties(user models.User) dtos.UserDTO {
	return dtos.UserDTO{Username: user.Username, Email: user.Email, Password: user.Password}
}
