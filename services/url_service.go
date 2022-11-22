package services

import (
	"errors"
	"example.com/m/v2/config"
	"example.com/m/v2/domain"
	"example.com/m/v2/internal/app/models"
	"example.com/m/v2/repositories"
	"example.com/m/v2/tools"
	"github.com/google/uuid"
	"net/url"
)

type URLService struct{}

func NewURLService() *URLService {
	return &URLService{}
}

var urlRepo = repositories.NewURLRepo()

func (us *URLService) Delete(ids []string, token string) {
	for _, id := range ids {
		err := urlRepo.DeleteURL(id, token)
		if err != nil {
			println(err.Error())
		}
	}
}

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

	result, err := urlRepo.Save(urlEntity)
	if result.FullURL != "" && err != nil {
		return result, err
	}

	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}

func (us *URLService) Get(id string) (domain.URL, error) {
	var address = config.Env.Address

	result, err := urlRepo.Get("http://" + address + "/" + id)
	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}

func (us *URLService) GetByFullURL(url string) (domain.URL, error) {
	result, err := urlRepo.GetByFullURL(url)

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

	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}

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
