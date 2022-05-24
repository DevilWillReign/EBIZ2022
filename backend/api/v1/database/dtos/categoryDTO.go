package dtos

import (
	"gorm.io/gorm"
)

type CategoryDTO struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Description string
	Product     []ProductDTO
}
