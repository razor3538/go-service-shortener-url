package domain

import (
	"gorm.io/gorm"
)

type Base struct {
	ID      string `gorm:"type:varchar" json:"correlation_id"`
	Deleted gorm.DeletedAt
}
