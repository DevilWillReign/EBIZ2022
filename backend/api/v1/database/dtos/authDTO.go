package dtos

import (
	"gorm.io/gorm"
)

type AuthType int

const (
	_ AuthType = iota
	Github
	Google
	Slack
)

type AuthDTO struct {
	gorm.Model
	Authtype  AuthType `gorm:"not null"`
	UserDTOID uint     `gorm:"unique;not null"`
}
