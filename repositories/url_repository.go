package repositories

import (
	"encoding/json"
	"example.com/m/v2/config"
	"example.com/m/v2/domain"
	"os"
	"strings"
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
	filePath := config.Env.FilePath

	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
		if err != nil {
			err = os.Mkdir(strings.Split(filePath, "/")[0], 0777)
			if err != nil {
				println(err.Error())
				return domain.URL{}, err
			}
			file, err = os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
			if err != nil {
				return domain.URL{}, err
			}
		}

		data, err := json.Marshal(url)
		if err != nil {
			return domain.URL{}, err
		}

		data = append(data, '\n')

		_, err = file.Write(data)
		if err != nil {
			return domain.URL{}, err
		}
	}

	return url, nil
}

func (ur *URLRepo) Get(id string) (domain.URL, error) {
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

func (ur *URLRepo) GetByFullURL(id string) (domain.URL, error) {
	var url domain.URL
	if err := config.DB.
		Table("urls as u").
		Select("u.*").
		Where("u.full_url = ?", id).
		Scan(&url).
		Error; err != nil {
		return domain.URL{}, err
	}
	return url, nil
}
