package migrations

import (
	"fmt"
	"m-banking/internal/core/domain/models"
	"m-banking/pkg/database/sql"
)

func RunMigration() {
	err := sql.DB.AutoMigrate(
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
