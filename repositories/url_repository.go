package repositories

import (
	"example.com/m/v2/config"
	"example.com/m/v2/domain"
)

type URLRepo struct{}

func NewURLRepo() *URLRepo {
	return &URLRepo{}
}

func (ur *URLRepo) Save(url domain.URL) (domain.URL, error) {
	if err := config.DB.
		Create(&url).
		Error; err != nil {
		return domain.URL{}, err
	}
	return url, nil
}

func (ur *URLRepo) Get(id string) (domain.URL, error) {
	println(id)
	var url domain.URL
	if err := config.DB.
		Table("urls as u").
		Select("u.*").
		Where("u.short_url = ?", id).
		Scan(&url).
		Error; err != nil {
		return domain.URL{}, err
	}
	return url, nil
}
