package domain

type Base struct {
	ID int64 `gorm:"type:varchar;" json:"id"`
}
