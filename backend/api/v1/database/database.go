package database

import (
	"crypto/rand"
	"fmt"
	mathrand "math/rand"

	"github.com/shopspring/decimal"
	"golang.org/x/crypto/argon2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"apprit/store/api/v1/database/dtos"
	"apprit/store/api/v1/utils"
)

func CreateDatabase() *gorm.DB {
	databaseUser := utils.GetEnv("DATABASE_USER", "")
	databasePassword := utils.GetEnv("DATABASE_PASS", "")
	databaseAddress := utils.GetEnv("DATABASE_ADDRESS", "")
	databasePort := utils.GetEnv("DATABASE_PORT", "")
	databaseName := utils.GetEnv("DATABASE_NAME", "test")
	dsn := ""
	if databasePassword == "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseAddress, databasePort, databaseName)
	} else {
		dsn = fmt.Sprintf("%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databaseAddress, databasePort, databaseName)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&dtos.CategoryDTO{})
	db.AutoMigrate(&dtos.ProductDTO{})
	db.AutoMigrate(&dtos.UserDTO{})
	db.AutoMigrate(&dtos.AuthDTO{})
	db.AutoMigrate(&dtos.PaymentDTO{})
	db.AutoMigrate(&dtos.TransactionDTO{})
	db.AutoMigrate(&dtos.QuantifiedProductDTO{})
	return db
}

func InitializeDatabaseData(db *gorm.DB) {
	productPrice, _ := decimal.NewFromString("1000000.05")
	salt := make([]byte, 10)
	rand.Read(salt)
	key := argon2.IDKey([]byte("admin@example.com1234"), salt, 3, 32*1024, 4, 32)
	db.Create(&dtos.UserDTO{Username: "adminin", Email: "admin@example.com", Password: string(key), Salt: salt, Admin: true})
	for i := 1; i <= 5; i++ {
		num := fmt.Sprint(i)
		addCategoryIfNotExists(db, "Category"+num)
		addProductIfNotExists(db, "P"+num, "Product"+num, productPrice)
		addUserIfNotExists(db, "user"+num, "user"+num+"@example.com")
		addAuthIfNotExists(db, dtos.AuthType(i), uint(i))
		addTransactionIfNotExists(db, uint(i), uint(i))
		total := decimal.NewFromInt(0)
		for j := 1; j <= 5; j++ {
			quantity := mathrand.Intn(32)
			quantityD := decimal.NewFromInt(int64(quantity))
			addQuantifiedProductIfNotExists(db, uint(j), uint(quantity), uint(i))
			total = total.Add(productPrice.Mul(quantityD))
		}
		addPaymentIfNotExists(db, uint(i), total)
	}
}

func addCategoryIfNotExists(db *gorm.DB, name string) {
	c := new(dtos.CategoryDTO)
	err := db.Where("name = ?", name).Take(&c)
	if err.RowsAffected == 0 {
		db.Create(&dtos.CategoryDTO{
			Name:        name,
			Description: utils.RandStringBytes(20),
		})
	}
}

func addProductIfNotExists(db *gorm.DB, code string, name string, price decimal.Decimal) {
	c := new(dtos.ProductDTO)
	err := db.Where("code = ?", code).Take(&c)
	if err.RowsAffected == 0 {
		db.Create(&dtos.ProductDTO{
			Name:          name,
			Code:          code,
			Price:         price,
			Availability:  uint(mathrand.Int63n(64)),
			Description:   utils.RandStringBytes(20),
			CategoryDTOID: 1,
		})
	}
}

func addUserIfNotExists(db *gorm.DB, username string, email string) {
	c := new(dtos.UserDTO)
	err := db.Where("email = ?", email).Take(&c)
	if err.RowsAffected == 0 {
		salt := make([]byte, 10)
		rand.Read(salt)
		key := argon2.IDKey([]byte(username+email), salt, 3, 32*1024, 4, 32)
		db.Create(&dtos.UserDTO{Username: username, Email: email, Password: string(key), Salt: salt})
	}
}

func addAuthIfNotExists(db *gorm.DB, auth dtos.AuthType, userDTOID uint) {
	c := new(dtos.AuthDTO)
	err := db.Where("userDTOID = ?", userDTOID).Take(&c)
	if err.RowsAffected == 0 {
		db.Create(&dtos.AuthDTO{UserDTOID: userDTOID, Authtype: auth})
	}
}

func addTransactionIfNotExists(db *gorm.DB, id uint, userDTOID uint) {
	c := new(dtos.TransactionDTO)
	err := db.Where("id = ?", id).Take(&c)
	if err.RowsAffected == 0 {
		db.Create(&dtos.TransactionDTO{UserDTOID: userDTOID})
	}
}

func addQuantifiedProductIfNotExists(db *gorm.DB, productDTOID uint, quantity uint, transactionDTOID uint) {
	c := new(dtos.QuantifiedProductDTO)
	err := db.Where("transactiondtoid = ? and productdtoid = ?", transactionDTOID, productDTOID).Take(&c)
	if err.RowsAffected == 0 {
		db.Create(&dtos.QuantifiedProductDTO{ProductDTOID: productDTOID, TransactionDTOID: transactionDTOID, Quantity: quantity})
	}
}

func addPaymentIfNotExists(db *gorm.DB, transactiondtoid uint, total decimal.Decimal) {
	c := new(dtos.PaymentDTO)
	err := db.Where("transactiondtoid = ?", transactiondtoid).Take(&c)
	if err.RowsAffected == 0 {
		db.Create(&dtos.PaymentDTO{TransactionDTOID: transactiondtoid, Total: total})
	}
}
