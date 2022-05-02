package dtos

import (
	"gorm.io/gorm"
)

type AuthType int

const (
	Facebook AuthType = iota
	Github
	Google
	Twitter
)

type AuthDTO struct {
	gorm.Model
	Authtype  AuthType `gorm:"not null"`
	UserDTOID uint     `gorm:"unique;not null"`
}
