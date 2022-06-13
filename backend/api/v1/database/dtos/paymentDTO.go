package dtos

import (
	"gorm.io/gorm"
)

type PaymentType int

const (
	_ PaymentType = iota
	Card
	Transfer
	PayPal
)

type PaymentDTO struct {
	gorm.Model
	PaymentType      PaymentType `gorm:"not null"`
	TransactionDTOID uint        `gorm:"unique;not null"`
}
