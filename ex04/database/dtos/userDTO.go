package dtos

import "gorm.io/gorm"

type UserDTO struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Salt     []byte `gorm:"not null"`
}
