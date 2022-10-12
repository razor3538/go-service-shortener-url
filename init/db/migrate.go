package main

import (
	"example.com/m/v2/config"
	"example.com/m/v2/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gopkg.in/gormigrate.v1"
	"log"
)

func main() {
	db := config.DB

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202210111116",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&domain.URL{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("url").Error
			},
		},
	})

	err := m.Migrate()

	log.Println("=== ADD MIGRATIONS ===")

	if err == nil {
		println("Migration did run successfully")
	} else {
		println("Could not migrate: ", err.Error())
	}
}
