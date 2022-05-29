package daos

import (
	"apprit/store/api/v1/database/dtos"
	"crypto/rand"

	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) ([]dtos.UserDTO, error) {
	var users []dtos.UserDTO
	return GetEntities(db, &users)
}

func GetUserById(db *gorm.DB, id uint64) (dtos.UserDTO, error) {
	var userDTO dtos.UserDTO
	return GetEntityById(db, id, &userDTO)
}

func GetUserByEmail(db *gorm.DB, email string) (dtos.UserDTO, error) {
	var userDTO dtos.UserDTO
	if err := db.Where("email = ?", email).First(&userDTO).Error; err != nil {
		return userDTO, err
	}
	return userDTO, nil
}

func DeleteUserById(db *gorm.DB, id uint64) error {
	var userDTO dtos.UserDTO
	return DeleteEntityById(db, id, &userDTO)
}

func AddUser(db *gorm.DB, userDTO dtos.UserDTO) (dtos.UserDTO, error) {
	salt := make([]byte, 10)
	rand.Read(salt)
	userDTO.Salt = salt
	userDTO.Password = string(argon2.IDKey([]byte(userDTO.Password), salt, 3, 32*1024, 4, 32))
	return AddEntity(db, &userDTO)
}

func ReplaceUser(db *gorm.DB, id uint64, userDTO dtos.UserDTO) error {
	if _, err := GetUserById(db, id); err != nil {
		return err
	}
	salt := make([]byte, 10)
	rand.Read(salt)
	userDTO.Salt = salt
	userDTO.Password = string(argon2.IDKey([]byte(userDTO.Password), salt, 3, 32*1024, 4, 32))
	newValues := map[string]interface{}{"username": userDTO.Username, "password": userDTO.Password, "salt": userDTO.Salt, "email": userDTO.Email}
	return ReplaceEntity(db, id, newValues, &dtos.UserDTO{})
}
