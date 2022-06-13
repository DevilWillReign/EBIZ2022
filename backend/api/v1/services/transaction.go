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
	transactions := []models.Transaction{}
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

func AddTransaction(db *gorm.DB, transaction models.PostTransaction) (models.Transaction, error) {
	transactionDTO, err := daos.AddTransaction(db, copyTransactionDTOPropertiesFromPost(transaction))
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

func GetUserTransactions(db *gorm.DB, userId uint64) ([]models.Transaction, error) {
	transactionDTOs, err := daos.GetTransactionsByUserId(db, userId)
	if err != nil {
		return []models.Transaction{}, nil
	}
	transactions := []models.Transaction{}
	for _, transactionDTO := range transactionDTOs {
		transactions = append(transactions, copyTransactionProperties(transactionDTO))
	}
	return transactions, nil
}

func AddUserTransaction(db *gorm.DB, transaction models.UserTransaction) (int64, error) {
	addedTransaction, err := AddTransaction(db, convertUserTransactionToTransaction(transaction))
	if err != nil {
		return -1, err
	}
	for i := range transaction.QuantifiedProducts {
		transaction.QuantifiedProducts[i].TransactionID = addedTransaction.ID
	}
	if err := AddAllQuantifiedProduct(db, transaction.QuantifiedProducts); err != nil {
		return -1, err
	}
	return int64(addedTransaction.ID), nil
}

func GetUserTransactionById(db *gorm.DB, userId uint64, id uint64) (models.Transaction, error) {
	transactionDTO, err := daos.GetUserTransactionById(db, userId, id)
	if err != nil {
		return models.Transaction{}, echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return copyTransactionProperties(transactionDTO), nil
}

func copyTransactionProperties(transactionDTO dtos.TransactionDTO) models.Transaction {
	return models.Transaction{ID: transactionDTO.ID, UserID: transactionDTO.UserDTOID, CreatedAt: transactionDTO.CreatedAt, Total: transactionDTO.Total}
}

func convertUserTransactionToTransaction(userTransaction models.UserTransaction) models.PostTransaction {
	return models.PostTransaction{ID: userTransaction.ID, UserID: userTransaction.UserID, CreatedAt: userTransaction.CreatedAt, Total: userTransaction.Total}
}

func copyTransactionDTOProperties(transaction models.Transaction) dtos.TransactionDTO {
	return dtos.TransactionDTO{UserDTOID: transaction.UserID, Total: transaction.Total}
}

func copyTransactionDTOPropertiesFromPost(transaction models.PostTransaction) dtos.TransactionDTO {
	return dtos.TransactionDTO{UserDTOID: transaction.UserID, Total: transaction.Total}
}
