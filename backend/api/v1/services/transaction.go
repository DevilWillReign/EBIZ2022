package services

import (
	"apprit/store/api/v1/database/daos"
	"apprit/store/api/v1/database/dtos"
	"apprit/store/api/v1/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetTransactions(db *gorm.DB) ([]models.Transaction, error) {
	transactionDTOs, err := daos.GetTransactions(db)
	if err != nil {
		return []models.Transaction{}, nil
	}
	var transactions []models.Transaction
	for _, transactionDTO := range transactionDTOs {
		transactions = append(transactions, copyTransactionProperties(transactionDTO))
	}
	return transactions, nil
}

func GetTransactionById(db *gorm.DB, id uint64) (models.Transaction, error) {
	transactionDTO, err := daos.GetTransactionById(db, id)
	if err != nil {
		return models.Transaction{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return copyTransactionProperties(transactionDTO), nil
}

func DeleteTransactionById(db *gorm.DB, id uint64) error {
	err := daos.DeleteTransactionById(db, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return nil
}

func AddTransaction(db *gorm.DB, transaction models.Transaction) (models.Transaction, error) {
	transactionDTO, err := daos.AddTransaction(db, copyTransactionDTOProperties(transaction))
	if err != nil {
		return models.Transaction{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return copyTransactionProperties(transactionDTO), nil
}

func ReplaceTransaction(db *gorm.DB, id uint64, transaction models.Transaction) error {
	err := daos.ReplaceTransaction(db, id, copyTransactionDTOProperties(transaction))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func copyTransactionProperties(transactionDTO dtos.TransactionDTO) models.Transaction {
	return models.Transaction{ID: transactionDTO.ID, UserID: transactionDTO.UserDTOID, CreatedAt: transactionDTO.CreatedAt}
}

func copyTransactionDTOProperties(transaction models.Transaction) dtos.TransactionDTO {
	return dtos.TransactionDTO{UserDTOID: transaction.UserID}
}
