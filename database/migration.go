package database

import (
	"e-wallet/models"
	"e-wallet/pkg/postgresql"
	"fmt"
)

func RunMigration() {
	err := postgresql.DB.AutoMigrate(
		&models.Customer{},
		&models.AccountNumber{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("migration failed...")
	}

	fmt.Println("migration successfully...")
}
