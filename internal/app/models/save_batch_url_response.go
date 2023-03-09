package models

// SaveBatchURLResponse модель респонса после сохранения урлов
type SaveBatchURLResponse struct {
	ID       string `json:"correlation_id"`
	ShortURL string `json:"short_url"`
}
