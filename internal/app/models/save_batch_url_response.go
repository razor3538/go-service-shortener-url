package models

type SaveBatchURLResponse struct {
	ID       string `json:"correlation_id"`
	ShortURL string `json:"short_url"`
}
