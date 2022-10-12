package domain

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

// Base содержит базовые поля для всех таблиц БД
type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt string    `gorm:"type:varchar(10);" json:"-"`
}

// BeforeCreate создает уникальный UUID и назначает дату удаления элемента
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	if base.ID.IsNil() {
		deleteDay := time.Now().AddDate(0, 2, 0)
		uuidv4, _ := uuid.NewV4()

		err := scope.SetColumn("deleted_at", deleteDay.Format("01-02-2006"))
		if err != nil {
			return err
		}
		return scope.SetColumn("ID", uuidv4)
	}
	return nil
}
