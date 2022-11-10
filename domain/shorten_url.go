package domain

type URL struct {
	Base
	FullURL  string `gorm:"type:varchar"`
	ShortURL string `gorm:"type:varchar"`
	UserID   string `gorm:"type:varchar"`
}
