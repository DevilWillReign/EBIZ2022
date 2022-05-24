package dtos

import (
	"gorm.io/gorm"
)

type AuthType int

const (
	Github AuthType = iota
	Google
	Slack
)

type AuthDTO struct {
	gorm.Model
	Authtype  AuthType `gorm:"not null"`
	UserDTOID uint     `gorm:"unique;not null"`
}
