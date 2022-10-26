package models

type PathID struct {
	ID string `uri:"id" json:"id" binding:"required"`
}
