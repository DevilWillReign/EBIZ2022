package dtos

import (
	"gorm.io/gorm"
)

type PaymentType int

const (
	Card PaymentType = iota
	Transfer
	PayPal
)

type PaymentDTO struct {
	gorm.Model
	PaymentType      PaymentType `gorm:"not null"`
	TransactionDTOID uint        `gorm:"unique;not null"`
}
