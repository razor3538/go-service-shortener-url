package models

// PathID модель ID для удаления урлов
type PathID struct {
	ID string `uri:"id" json:"id" binding:"required"`
}
