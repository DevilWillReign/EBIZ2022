package dtos

import "gorm.io/gorm"

type QuantifiedProductDTO struct {
	gorm.Model
	ProductDTOID     uint `gorm:"not null"`
	Quantity         uint `gorm:"not null"`
	TransactionDTOID uint `gorm:"not null"`
}
