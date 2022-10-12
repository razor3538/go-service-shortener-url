package repositories

import (
	"example.com/m/v2/config"
	"example.com/m/v2/domain"
)

type UrlRepo struct{}

func NewUrlRepo() *UrlRepo {
	return &UrlRepo{}
}

func (ur *UrlRepo) Save(url domain.Url) (domain.Url, error) {
	if err := config.DB.
		Create(&url).
		Error; err != nil {
		return domain.Url{}, err
	}
	return url, nil
}

func (ur *UrlRepo) Get(id string) (domain.Url, error) {
	var url domain.Url
	if err := config.DB.
		Table("urls as u").
		Select("u.*").
		Where("u.id = ?", id).
		Scan(&url).
		Error; err != nil {
		return domain.Url{}, err
	}
	return url, nil
}
