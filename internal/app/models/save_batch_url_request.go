package models

type SaveBatchURLRequest struct {
	ID      string `json:"correlation_id"`
	FullURL string `json:"original_url"`
}
