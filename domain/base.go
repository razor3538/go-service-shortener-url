package domain

type Base struct {
	ID int64 `gorm:"type:primaryKey" json:"id"`
}
