package models

// SaveBatchURLRequest модель для сохранения множества урлов 1 запросом
type SaveBatchURLRequest struct {
	ID      string `json:"correlation_id"`
	FullURL string `json:"original_url"`
}
