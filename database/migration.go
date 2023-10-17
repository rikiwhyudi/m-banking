package database

import (
	"fmt"
	"m-banking/domain/models"
	"m-banking/pkg/postgresql"
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
