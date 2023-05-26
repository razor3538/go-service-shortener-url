package models

// UserAndUrlsResponse модель количества пользователей и сокращенных урлов
type UserAndUrlsResponse struct {
	UsersCount int `json:"users"`
	UrlsCount  int `json:"urls"`
}
