package services

import (
	"apprit/store/api/v1/database/daos"
	"apprit/store/api/v1/database/dtos"
	"apprit/store/api/v1/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetProducts(db *gorm.DB) ([]models.Product, error) {
	productDTOs, err := daos.GetProducts(db)
	if err != nil {
		return []models.Product{}, nil
	}
	var products []models.Product
	for _, productDTO := range productDTOs {
		products = append(products, copyProductProperties(productDTO))
	}
	return products, nil
}

func GetProductById(db *gorm.DB, id uint64) (models.Product, error) {
	productDTO, err := daos.GetProductById(db, id)
	if err != nil {
		return models.Product{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return copyProductProperties(productDTO), nil
}

func DeleteProductById(db *gorm.DB, id uint64) error {
	err := daos.DeleteProductById(db, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return nil
}

func AddProduct(db *gorm.DB, product models.Product) (models.Product, error) {
	productDTO, err := daos.AddProduct(db, copyProductDTOProperties(product))
	if err != nil {
		return models.Product{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return copyProductProperties(productDTO), nil
}

func ReplaceProduct(db *gorm.DB, id uint64, product models.Product) error {
	err := daos.ReplaceProduct(db, id, copyProductDTOProperties(product))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func copyProductProperties(productDTO dtos.ProductDTO) models.Product {
	return models.Product{
		ID:           productDTO.ID,
		Name:         productDTO.Name,
		Code:         productDTO.Code,
		Price:        productDTO.Price,
		Availability: productDTO.Availability,
		Description:  productDTO.Description,
		CategoryID:   productDTO.CategoryDTOID,
	}
}

func copyProductDTOProperties(product models.Product) dtos.ProductDTO {
	return dtos.ProductDTO{
		Name:          product.Name,
		Code:          product.Code,
		Price:         product.Price,
		Availability:  product.Availability,
		Description:   product.Description,
		CategoryDTOID: product.CategoryID,
	}
}
