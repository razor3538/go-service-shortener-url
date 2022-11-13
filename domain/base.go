package domain

type Base struct {
	ID string `gorm:"type:varchar" json:"correlation_id"`
}
