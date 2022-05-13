package daos

import (
	"apprit/store/api/v1/database/dtos"

	"gorm.io/gorm"
)

func GetProducts(db *gorm.DB) ([]dtos.ProductDTO, error) {
	var products []dtos.ProductDTO
	return GetEntities(db, &products)
}

func GetProductById(db *gorm.DB, id uint64) (dtos.ProductDTO, error) {
	var productDTO dtos.ProductDTO
	return GetEntityById(db, id, &productDTO)
}

func DeleteProductById(db *gorm.DB, id uint64) error {
	var productDTO dtos.ProductDTO
	return DeleteEntityById(db, id, &productDTO)
}

func AddProduct(db *gorm.DB, productDTO dtos.ProductDTO) error {
	return AddEntity(db, &productDTO)
}

func ReplaceProduct(db *gorm.DB, id uint64, productDTO dtos.ProductDTO) error {
	newValues := map[string]interface{}{"code": productDTO.Code, "categorydtoid": productDTO.CategoryDTOID,
		"name": productDTO.Name, "price": productDTO.Price}
	return ReplaceEntity(db, id, newValues, &dtos.ProductDTO{})
}
