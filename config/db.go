package config

import (
	"database/sql"
	"example.com/m/v2/domain"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

// DB сущность базы данных
var DB *gorm.DB

// InitBD инициализирует подключение к базе данных либо MySQL либо PgSQL
func InitBD() {
	err := initPgSQL()

	if err != nil {
		initMySQL()
		_, err := fmt.Println("You connected to your database.")

		if err != nil {
			println(err)
		}
	}
}

// Инициализация базы данных MySQL
func initMySQL() {
	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var db *gorm.DB

	sqlDB, err := sql.Open("sqlite3", "gorm.db")

	if err != nil {
		panic(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}

	if filepath.Base(dirname) == "go-service-shortener-url" {
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	} else {
		db, err = gorm.Open(sqlite.Open("../gorm.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	}

	DB = db

	err = db.Table("urls").AutoMigrate(&domain.URL{})
	if err != nil {
		panic(err)
	}
}

// Инициализация базы данных PgSQL
func initPgSQL() error {
	var db *gorm.DB

	db, err := gorm.Open(postgres.Open(Env.BdConnection), &gorm.Config{})

	if err != nil {
		return err
	}

	DB = db

	err = db.Table("urls").AutoMigrate(&domain.URL{})
	if err != nil {
		return err
	}

	return nil
}
