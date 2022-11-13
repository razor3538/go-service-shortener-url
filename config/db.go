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

func init() {
	err := initPgSQL()
	println(err.Error())
	println(err.Error())
	println(err.Error())
	println(err.Error())
	println(err.Error())
	if err != nil {
		println("подключаемся в склайт")
		initMySQL()
	}

	fmt.Println("You connected to your database.")
}

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

	println("постря настроенна")
	return nil
}
