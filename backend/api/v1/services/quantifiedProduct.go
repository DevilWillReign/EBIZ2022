package services

import (
	"apprit/store/api/v1/database/daos"
	"apprit/store/api/v1/database/dtos"
	"apprit/store/api/v1/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetQuantifiedProducts(db *gorm.DB) ([]models.QuantifiedProduct, error) {
	quantifiedProductDTOs, err := daos.GetQuantifiedProducts(db)
	if err != nil {
		return []models.QuantifiedProduct{}, nil
	}
	var quantifiedProducts []models.QuantifiedProduct
	for _, quantifiedProductDTO := range quantifiedProductDTOs {
		productDTO, _ := daos.GetProductById(db, uint64(quantifiedProductDTO.ProductDTOID))
		quantifiedProducts = append(quantifiedProducts, copyQuantifiedProductProperties(quantifiedProductDTO, productDTO))
	}
	return quantifiedProducts, nil
}

func GetQuantifiedProductById(db *gorm.DB, id uint64) (models.QuantifiedProduct, error) {
	quantifiedProductDTO, err := daos.GetQuantifiedProductById(db, id)
	productDTO, errP := daos.GetProductById(db, uint64(quantifiedProductDTO.ProductDTOID))
	if err != nil || errP != nil {
		return models.QuantifiedProduct{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return copyQuantifiedProductProperties(quantifiedProductDTO, productDTO), nil
}

func DeleteQuantifiedProductById(db *gorm.DB, id uint64) error {
	err := daos.DeleteQuantifiedProductById(db, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return nil
}

func AddQuantifiedProduct(db *gorm.DB, quantifiedProduct models.QuantifiedProduct) (models.QuantifiedProduct, error) {
	quantifiedProductDTO, err := daos.AddQuantifiedProduct(db, copyQuantifiedProductDTOProperties(quantifiedProduct))
	if err != nil {
		return models.QuantifiedProduct{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	productDTO, _ := daos.GetProductById(db, uint64(quantifiedProductDTO.ProductDTOID))
	return copyQuantifiedProductProperties(quantifiedProductDTO, productDTO), nil
}

func ReplaceQuantifiedProduct(db *gorm.DB, id uint64, quantifiedProduct models.QuantifiedProduct) error {
	err := daos.ReplaceQuantifiedProduct(db, id, copyQuantifiedProductDTOProperties(quantifiedProduct))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func GetQuantifiedProductByTransactionId(db *gorm.DB, transactionId uint64) ([]models.QuantifiedProduct, error) {
	quantifiedProductDTOs, err := daos.GetQuantifiedProductsByTransactionId(db, transactionId)
	if err != nil {
		return []models.QuantifiedProduct{}, nil
	}
	var quantifiedProducts []models.QuantifiedProduct
	for _, quantifiedProductDTO := range quantifiedProductDTOs {
		productDTO, _ := daos.GetProductById(db, uint64(quantifiedProductDTO.ProductDTOID))
		quantifiedProducts = append(quantifiedProducts, copyQuantifiedProductProperties(quantifiedProductDTO, productDTO))
	}
	return quantifiedProducts, nil
}

func copyQuantifiedProductProperties(quantifiedProductDTO dtos.QuantifiedProductDTO, productDTO dtos.ProductDTO) models.QuantifiedProduct {
	return models.QuantifiedProduct{ID: productDTO.ID, Name: productDTO.Name, Code: productDTO.Code, Price: productDTO.Price, CategoryID: productDTO.CategoryDTOID, Quantity: quantifiedProductDTO.Quantity, TransactionID: quantifiedProductDTO.TransactionDTOID}
}

func copyQuantifiedProductDTOProperties(quantifiedProduct models.QuantifiedProduct) dtos.QuantifiedProductDTO {
	return dtos.QuantifiedProductDTO{ProductDTOID: quantifiedProduct.ID, Quantity: quantifiedProduct.Quantity, TransactionDTOID: quantifiedProduct.TransactionID}
}
