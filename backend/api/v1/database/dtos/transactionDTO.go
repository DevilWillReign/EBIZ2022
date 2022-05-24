package dtos

import "gorm.io/gorm"

type TransactionDTO struct {
	gorm.Model
	UserDTOID         uint `gorm:"not null"`
	PaymentDTO        PaymentDTO
	QuantifiedProduct []QuantifiedProductDTO
}
