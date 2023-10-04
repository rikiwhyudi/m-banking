package mysql

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	MYSQL_USER := os.Getenv("DB_USER")
	MYSQL_PASSWORD := os.Getenv("DB_PASSWORD")
	MYSQL_HOST := os.Getenv("DB_HOST")
	MYSQL_PORT := os.Getenv("DB_PORT")
	MYSQL_DB := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT, MYSQL_DB)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("connected to MySQL Database...")

}
