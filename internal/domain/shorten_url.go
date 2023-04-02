package domain

// URL Структура хранящихся в базе данных урлов
type URL struct {
	Base
	FullURL       string `gorm:"type:varchar; unique"`
	ShortURL      string `gorm:"type:varchar"`
	UserID        string `gorm:"type:varchar"`
	StringShortID string `gorm:"type:varchar"`
}
