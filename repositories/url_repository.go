package repositories

import (
	"encoding/json"
	"errors"
	"example.com/m/v2/config"
	"example.com/m/v2/domain"
	"example.com/m/v2/internal/app/models"
	"os"
	"strings"
)

type URLRepo struct{}

func NewURLRepo() *URLRepo {
	return &URLRepo{}
}

func (ur *URLRepo) Save(url domain.URL) (domain.URL, error) {
	var existingUrl domain.URL

	if err := config.DB.
		Table("urls as u").
		Select("u.*").
		Where("u.full_url = ?", url.FullURL).
		Scan(&existingUrl).
		Error; err != nil {
		return domain.URL{}, err
	}

	if existingUrl.FullURL != "" {
		return existingUrl, errors.New("урл уже сохранен")
	}

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

func (ur *URLRepo) SaveMany(urls []domain.URL) ([]domain.URL, error) {
	var urlsResponse []domain.URL
	var urlsID []string

	for _, urlID := range urls {
		urlsID = append(urlsID, urlID.ID)
	}

	if err := config.DB.Create(&urls).Error; err != nil {
		return []domain.URL{}, err
	}

	if err := config.DB.
		Table("urls as u").
		Select("u.*").
		Where("u.id in ?", urlsID).
		Scan(&urlsResponse).Error; err != nil {
		return []domain.URL{}, err
	}

	return urlsResponse, nil
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

func (ur *URLRepo) GetByUserID(id string) ([]models.FullURL, error) {
	var url []models.FullURL
	if err := config.DB.Model(&domain.URL{}).Where("user_id = ?", id).Pluck("full_url, short_url", &url).Error; err != nil {
		return []models.FullURL{}, err
	}
	return url, nil
}
