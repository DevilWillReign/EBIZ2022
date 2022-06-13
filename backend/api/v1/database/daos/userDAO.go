package daos

import (
	"apprit/store/api/v1/database/dtos"
	"apprit/store/api/v1/utils"
	"crypto/rand"

	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) ([]dtos.UserDTO, error) {
	users := []dtos.UserDTO{}
	return GetEntities(db, &users)
}

func GetUserById(db *gorm.DB, id uint64) (dtos.UserDTO, error) {
	userDTO := dtos.UserDTO{}
	return GetEntityById(db, id, &userDTO)
}

func GetUserByEmail(db *gorm.DB, email string) (dtos.UserDTO, error) {
	userDTO := dtos.UserDTO{}
	if err := db.Where("email = ?", email).First(&userDTO).Error; err != nil {
		return userDTO, err
	}
	return userDTO, nil
}

func DeleteUserById(db *gorm.DB, id uint64) error {
	userDTO := dtos.UserDTO{}
	return DeleteEntityById(db, id, &userDTO)
}

func AddUser(db *gorm.DB, userDTO dtos.UserDTO) (dtos.UserDTO, error) {
	salt := make([]byte, 10)
	rand.Read(salt)
	userDTO.Salt = salt
	userDTO.Password = utils.HashPassword([]byte(userDTO.Password), salt)
	return AddEntity(db, &userDTO)
}

func ReplaceUser(db *gorm.DB, id uint64, userDTO dtos.UserDTO) error {
	if _, err := GetUserById(db, id); err != nil {
		return err
	}
	salt := make([]byte, 10)
	rand.Read(salt)
	userDTO.Salt = salt
	userDTO.Password = utils.HashPassword([]byte(userDTO.Password), salt)
	newValues := map[string]interface{}{"username": userDTO.Username, "password": userDTO.Password, "salt": userDTO.Salt, "email": userDTO.Email}
	return ReplaceEntity(db, id, newValues, &dtos.UserDTO{})
}
