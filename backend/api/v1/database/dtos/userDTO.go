package dtos

import "gorm.io/gorm"

type UserDTO struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Admin    bool   `gorm:"not null"`
	Password string
	Salt     []byte
}
