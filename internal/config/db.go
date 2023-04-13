package config

import (
	"database/sql"
	"go-service-shortener-url/internal/domain"
	"go-service-shortener-url/internal/tools"
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
		tools.InfoLog.Println("You connected to your database.")
	}
}

// Инициализация базы данных MySQL
func initMySQL() {
	dirname, err := os.Getwd()
	if err != nil {
		tools.ErrorLog.Fatal(err)
	}

	var db *gorm.DB

	sqlDB, err := sql.Open("sqlite3", "gorm.db")

	if err != nil {
		tools.ErrorLog.Fatal(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		tools.ErrorLog.Fatal(err)
	}

	err = sqlDB.Close()
	if err != nil {
		tools.ErrorLog.Fatal(err)
	}

	if filepath.Base(dirname) == "go-service-shortener-url" {
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
		if err != nil {
			tools.ErrorLog.Fatal(err)
		}
	} else {
		db, err = gorm.Open(sqlite.Open("../gorm.db"), &gorm.Config{})
		if err != nil {
			tools.ErrorLog.Fatal(err)
		}
	}

	DB = db

	err = db.Table("urls").AutoMigrate(&domain.URL{})
	if err != nil {
		tools.ErrorLog.Fatal(err)
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