package services

import (
	"apprit/store/api/v1/database/daos"
	"apprit/store/api/v1/database/dtos"
	"apprit/store/api/v1/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetCategories(db *gorm.DB) ([]models.Category, error) {
	categoryDTOs, err := daos.GetCategories(db)
	if err != nil {
		return []models.Category{}, nil
	}
	categories := []models.Category{}
	for _, category := range categoryDTOs {
		categories = append(categories, copyCategoryProperties(category))
	}
	return categories, nil
}

func GetCategoryById(db *gorm.DB, id uint64) (models.Category, error) {
	categoryDTO, err := daos.GetCategoryById(db, id)
	if err != nil {
		return models.Category{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return copyCategoryProperties(categoryDTO), nil
}

func DeleteCategoryById(db *gorm.DB, id uint64) error {
	err := daos.DeleteCategoryById(db, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return nil
}

func AddCategory(db *gorm.DB, category models.PostCategory) (models.Category, error) {
	categoryDTO, err := daos.AddCategory(db, copyCategoryDTOPropertiesFromPost(category))
	if err != nil {
		return models.Category{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return copyCategoryProperties(categoryDTO), nil
}

func ReplaceCategory(db *gorm.DB, id uint64, category models.Category) error {
	err := daos.ReplaceCategory(db, id, copyCategoryDTOProperties(category))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func copyCategoryProperties(categoryDTO dtos.CategoryDTO) models.Category {
	return models.Category{
		ID:          categoryDTO.ID,
		Name:        categoryDTO.Name,
		Description: categoryDTO.Description,
	}
}

func copyCategoryDTOProperties(category models.Category) dtos.CategoryDTO {
	return dtos.CategoryDTO{
		Name:        category.Name,
		Description: category.Description,
	}
}

func copyCategoryDTOPropertiesFromPost(category models.PostCategory) dtos.CategoryDTO {
	return dtos.CategoryDTO{
		Name:        category.Name,
		Description: category.Description,
	}
}
