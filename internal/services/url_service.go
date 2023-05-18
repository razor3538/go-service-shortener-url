package services

import (
	"errors"
	"net/url"

	"example.com/m/v2/internal/app/models"
	"example.com/m/v2/internal/config"
	"example.com/m/v2/internal/domain"
	"example.com/m/v2/internal/repositories"
	"example.com/m/v2/internal/tools"
	"github.com/google/uuid"
)

// URLService структура
type URLService struct{}

// NewURLService возвращает указатель на структуру URLService
// со всеми ее методами
func NewURLService() *URLService {
	return &URLService{}
}

var urlRepo = repositories.NewURLRepo()

// Delete сервис для удаления урлов
func (us *URLService) Delete(ids []string) {
	for _, id := range ids {
		err := urlRepo.DeleteURL(id)
		if err != nil {
			tools.ErrorLog.Println(err.Error())
		}
	}
}

// Save сервис для сохранения урлов
func (us *URLService) Save(urlModel string, userID string) (domain.URL, error) {
	var address = config.Env.Address

	var urlEntity domain.URL

	_, err := url.ParseRequestURI(urlModel)

	if err != nil {
		return domain.URL{}, errors.New(urlModel)
	}

	id, err := tools.ShortenURL(urlModel)

	if err != nil {
		return domain.URL{}, errors.New(urlModel)
	}

	urlEntity.ShortURL = "http://" + address + "/" + id
	urlEntity.FullURL = urlModel
	urlEntity.UserID = userID
	urlEntity.ID = uuid.New().String()
	urlEntity.StringShortID = id

	result, err := urlRepo.Save(urlEntity)
	if result.FullURL != "" && err != nil {
		return result, err
	}

	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}

// Get сервис для получения урлов
func (us *URLService) Get(id string) (domain.URL, error) {
	var address = config.Env.Address

	result, err := urlRepo.Get("http://" + address + "/" + id)
	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}

// GetByFullURL сервис для получения полной модели урлов
func (us *URLService) GetByFullURL(url string) (domain.URL, error) {
	result, errRepo := urlRepo.GetByFullURL(url)

	if errRepo != nil {
		return domain.URL{}, errRepo
	}

	if result.FullURL == "" {
		urlModel, err := us.Save(url, "")
		if err != nil {
			return domain.URL{}, err
		}
		result, err = urlRepo.GetByFullURL(urlModel.FullURL)
		if err != nil {
			return domain.URL{}, err
		}
	} else {
		return result, errors.New("урл уже сохранен")
	}

	return result, nil
}

// SaveMany сервис для сохранения нескольких урлов
func (us *URLService) SaveMany(urls []models.SaveBatchURLRequest) ([]models.SaveBatchURLResponse, error) {
	var domainUrls []domain.URL
	var response []models.SaveBatchURLResponse

	for i := range urls {
		tmp, err := tools.ShortenURL(urls[i].FullURL)
		if err != nil {
			return []models.SaveBatchURLResponse{}, err
		}
		domainUrls = append(domainUrls, domain.URL{
			Base:     domain.Base{ID: urls[i].ID},
			FullURL:  urls[i].FullURL,
			ShortURL: "http://" + config.Env.Address + "/" + tmp,
			UserID:   "",
		})
	}

	repositoriesResponse, err := urlRepo.SaveMany(domainUrls)
	if err != nil {
		return []models.SaveBatchURLResponse{}, err
	}

	for i := range repositoriesResponse {
		response = append(response, models.SaveBatchURLResponse{
			ID:       repositoriesResponse[i].ID,
			ShortURL: repositoriesResponse[i].ShortURL,
		})
	}

	return response, nil
}

// GetByUserID сервис для получения всех урлов по пользователю
func (us *URLService) GetByUserID(userID string) ([]models.FullURL, error) {
	result, err := urlRepo.GetByUserID(userID)

	if err != nil {
		return []models.FullURL{}, err
	}

	if len(result) == 0 {
		return []models.FullURL{}, errors.New("no content")
	}

	return result, nil
}

// GetAllUsersAndUrls возвращает количество уникальных пользователей и общее количество сокращенных урлов
func (us *URLService) GetAllUsersAndUrls() (models.UserAndUrlsResponse, error) {
	countUser, countUrls, err := urlRepo.GetAllUsersAndUrls()

	if err != nil {
		return models.UserAndUrlsResponse{}, err
	}

	return models.UserAndUrlsResponse{UrlsCount: countUrls, UsersCount: countUser}, nil
}
