package services

import (
	"apprit/store/api/v1/database/daos"
	"apprit/store/api/v1/database/dtos"
	"apprit/store/api/v1/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAuths(db *gorm.DB) ([]models.Auth, error) {
	authDTOs, err := daos.GetAuths(db)
	if err != nil {
		return []models.Auth{}, nil
	}
	var auths []models.Auth
	for _, authDTO := range authDTOs {
		auths = append(auths, copyAuthProperties(authDTO))
	}
	return auths, nil
}

func GetAuthById(db *gorm.DB, id uint64) (models.Auth, error) {
	authDTO, err := daos.GetAuthById(db, id)
	if err != nil {
		return models.Auth{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return copyAuthProperties(authDTO), nil
}

func DeleteAuthById(db *gorm.DB, id uint64) error {
	err := daos.DeleteAuthById(db, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return nil
}

func AddAuth(db *gorm.DB, auth models.Auth) error {
	err := daos.AddAuth(db, copyAuthDTOProperties(auth))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func ReplaceAuth(db *gorm.DB, id uint64, auth models.Auth) error {
	if err := daos.ReplaceAuth(db, id, copyAuthDTOProperties(auth)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func copyAuthProperties(authDTO dtos.AuthDTO) models.Auth {
	return models.Auth{ID: authDTO.ID, Authtype: authDTO.Authtype, UserID: authDTO.UserDTOID}
}

func copyAuthDTOProperties(auth models.Auth) dtos.AuthDTO {
	return dtos.AuthDTO{Authtype: auth.Authtype, UserDTOID: auth.UserID}
}
