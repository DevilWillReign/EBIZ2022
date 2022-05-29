package services

import (
	"apprit/store/api/v1/database/daos"
	"apprit/store/api/v1/database/dtos"
	"apprit/store/api/v1/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetPayments(db *gorm.DB) ([]models.Payment, error) {
	paymentDTOs, err := daos.GetPayments(db)
	if err != nil {
		return []models.Payment{}, nil
	}
	var payments []models.Payment
	for _, paymentDTO := range paymentDTOs {
		payments = append(payments, copyPaymentProperties(paymentDTO))
	}
	return payments, nil
}

func GetPaymentById(db *gorm.DB, id uint64) (models.Payment, error) {
	paymentDTO, err := daos.GetPaymentById(db, id)
	if err != nil {
		return models.Payment{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return copyPaymentProperties(paymentDTO), nil
}

func DeletePaymentById(db *gorm.DB, id uint64) error {
	err := daos.DeletePaymentById(db, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return nil
}

func AddPayment(db *gorm.DB, payment models.Payment) (models.Payment, error) {
	paymentDTO, err := daos.AddPayment(db, copyPaymentDTOProperties(payment))
	if err != nil {
		return models.Payment{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return copyPaymentProperties(paymentDTO), nil
}

func ReplacePayment(db *gorm.DB, id uint64, payment models.Payment) error {
	err := daos.ReplacePayment(db, id, copyPaymentDTOProperties(payment))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func copyPaymentProperties(paymentDTO dtos.PaymentDTO) models.Payment {
	return models.Payment{ID: paymentDTO.ID, Total: paymentDTO.Total, TransactionID: paymentDTO.TransactionDTOID}
}

func copyPaymentDTOProperties(payment models.Payment) dtos.PaymentDTO {
	return dtos.PaymentDTO{Total: payment.Total, TransactionDTOID: payment.TransactionID}
}
