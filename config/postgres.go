package config

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB сущность базы данных
var DB *gorm.DB

func init() {

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = db

	DB.Raw("CREATE TABLE `urls` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `full_url` VARCHAR(164), `short_url` VARCHAR(164));")

	fmt.Println("You connected to your database.")
}
