package repositories

import (
	"encoding/json"
	"errors"
	"go-service-shortener-url/internal/app/models"
	"go-service-shortener-url/internal/config"
	"go-service-shortener-url/internal/domain"
	"os"
	"strings"
)

// URLRepo стуктура
type URLRepo struct{}

// NewURLRepo возвращает указатель на структуру URLRepo
// со всеми ее методами
func NewURLRepo() *URLRepo {
	return &URLRepo{}
}

// Save сохраняет в базе данных сущность domain.URL
func (ur *URLRepo) Save(url domain.URL) (domain.URL, error) {
	var existingURL domain.URL

	if err := config.DB.
		Table("urls as u").
		Select("u.*").
		Where("u.full_url = ?", url.FullURL).
		Scan(&existingURL).
		Error; err != nil {
		return domain.URL{}, err
	}

	if existingURL.FullURL != "" {
		return existingURL, errors.New("урл уже сохранен")
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

// DeleteURL удаляет сущность domain.URL
func (ur *URLRepo) DeleteURL(id string) error {
	var tmp domain.URL
	if err := config.DB.
		Where("string_short_id = ?", id).
		Delete(&domain.URL{}).Scan(&tmp).
		Error; err != nil {
		return err
	}
	return nil
}

// SaveMany сохраняет сразу несколько сущностей domain.URL
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

// Get возвращает сущность domain.URL по предоставленному сокращенному урлу
func (ur *URLRepo) Get(shortURL string) (domain.URL, error) {
	var url domain.URL
	if err := config.DB.
		Table("urls as u").
		Select("u.*").
		Where("u.short_url = ?", shortURL).
		Scan(&url).
		Error; err != nil {
		return domain.URL{}, err
	}
	return url, nil
}

// GetByFullURL возвращает сущность domain.URL по предоставленному полному урлу
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

// GetByUserID возвращает массив сущностей models.FullURL по пользователю
func (ur *URLRepo) GetByUserID(id string) ([]models.FullURL, error) {
	var url []models.FullURL
	if err := config.DB.Model(&domain.URL{}).Where("user_id = ?", id).Pluck("full_url, short_url", &url).Error; err != nil {
		return []models.FullURL{}, err
	}
	return url, nil
}
