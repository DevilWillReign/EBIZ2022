package daos

import (
	"apprit/store/api/v1/database/dtos"

	"gorm.io/gorm"
)

func GetCategories(db *gorm.DB) ([]dtos.CategoryDTO, error) {
	var categories []dtos.CategoryDTO
	return GetEntities(db, &categories)
}

func GetCategoryById(db *gorm.DB, id uint64) (dtos.CategoryDTO, error) {
	var categoryDTO dtos.CategoryDTO
	return GetEntityById(db, id, &categoryDTO)
}

func DeleteCategoryById(db *gorm.DB, id uint64) error {
	var categoryDTO dtos.CategoryDTO
	return DeleteEntityById(db, id, &categoryDTO)
}

func AddCategory(db *gorm.DB, categoryDTO dtos.CategoryDTO) error {
	return AddEntity(db, &categoryDTO)
}

func ReplaceCategory(db *gorm.DB, id uint64, categoryDTO dtos.CategoryDTO) error {
	if _, err := GetCategoryById(db, id); err != nil {
		return err
	}
	newValues := map[string]interface{}{"name": categoryDTO.Name}
	return ReplaceEntity(db, id, newValues, &dtos.CategoryDTO{})
}
