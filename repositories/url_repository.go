package repositories

import (
	"example.com/m/v2/config"
	"example.com/m/v2/domain"
)

type UrlRepo struct{}

func NewUrlRepo() *UrlRepo {
	return &UrlRepo{}
}

func (ur *UrlRepo) Save(url domain.URL) (domain.URL, error) {
	if err := config.DB.
		Create(&url).
		Error; err != nil {
		return domain.URL{}, err
	}
	return url, nil
}

func (ur *UrlRepo) Get(id string) (domain.URL, error) {
	var url domain.URL
	if err := config.DB.
		Table("urls as u").
		Select("u.*").
		Where("u.id = ?", id).
		Scan(&url).
		Error; err != nil {
		return domain.URL{}, err
	}
	return url, nil
}
