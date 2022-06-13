package dtos

import "gorm.io/gorm"

type UserDTO struct {
	gorm.Model
	Username    string `gorm:"unique;not null"`
	Email       string `gorm:"unique;not null"`
	Admin       bool   `gorm:"not null"`
	Password    []byte
	Salt        []byte
	Auth        AuthDTO        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Transaction TransactionDTO `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
