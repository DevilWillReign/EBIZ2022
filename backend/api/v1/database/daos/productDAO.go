package daos

import (
	"apprit/store/api/v1/database/dtos"

	"gorm.io/gorm"
)

func GetProducts(db *gorm.DB) ([]dtos.ProductDTO, error) {
	products := []dtos.ProductDTO{}
	return GetEntities(db, &products)
}

func GetProductById(db *gorm.DB, id uint64) (dtos.ProductDTO, error) {
	productDTO := dtos.ProductDTO{}
	return GetEntityById(db, id, &productDTO)
}

func GetProductByCode(db *gorm.DB, code string) (dtos.ProductDTO, error) {
	product := dtos.ProductDTO{}
	if err := db.Where("code = ?", code).First(&product).Error; err != nil {
		return dtos.ProductDTO{}, err
	}
	return product, nil
}

func DeleteProductById(db *gorm.DB, id uint64) error {
	productDTO := dtos.ProductDTO{}
	return DeleteEntityById(db, id, &productDTO)
}

func AddProduct(db *gorm.DB, productDTO dtos.ProductDTO) (dtos.ProductDTO, error) {
	return AddEntity(db, &productDTO)
}

func ReplaceProduct(db *gorm.DB, id uint64, productDTO dtos.ProductDTO) error {
	if _, err := GetProductById(db, id); err != nil {
		return err
	}
	newValues := map[string]interface{}{"code": productDTO.Code, "category_dto_id": productDTO.CategoryDTOID,
		"name": productDTO.Name, "price": productDTO.Price}
	return ReplaceEntity(db, id, newValues, &dtos.ProductDTO{})
}

func GetProductsByCategoryId(db *gorm.DB, categorydtoid uint64) ([]dtos.ProductDTO, error) {
	products := []dtos.ProductDTO{}
	if err := db.Where("category_dto_id = ?", categorydtoid).Find(&products).Error; err != nil {
		return []dtos.ProductDTO{}, err
	}
	return products, nil
}
