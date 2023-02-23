package models

// FullURL модель полной сущности URL
type FullURL struct {
	FullURL  string `json:"original_url"`
	ShortURL string `json:"short_url"`
}
