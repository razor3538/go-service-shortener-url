package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB сущность базы данных
var DB *gorm.DB

func init() {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		Env.DBHost, Env.DBPort, Env.DBUser, Env.DBName, Env.DBPassword,
	)

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	DB = db

	fmt.Println("You connected to your database.")
}
