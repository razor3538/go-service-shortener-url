package services

import (
	"errors"
	"example.com/m/v2/config"
	"example.com/m/v2/domain"
	"example.com/m/v2/internal/app/models"
	"example.com/m/v2/repositories"
	"github.com/speps/go-hashids"
	"net/url"
)

type URLService struct{}

func NewURLService() *URLService {
	return &URLService{}
}

var urlRepo = repositories.NewURLRepo()

func (us *URLService) Save(urlModel string, userId string) (domain.URL, error) {
	var address = config.Env.Address

	var urlEntity domain.URL

	_, err := url.ParseRequestURI(urlModel)

	if err != nil {
		return domain.URL{}, errors.New(urlModel)
	}

	hd := hashids.NewData()
	hd.Salt = urlModel

	h, err := hashids.NewWithData(hd)

	if err != nil {
		return domain.URL{}, err
	}

	id, _ := h.Encode([]int{1, 2, 3})

	urlEntity.ShortURL = "http://" + address + "/" + id
	urlEntity.FullURL = urlModel
	urlEntity.UserID = userId

	result, err := urlRepo.Save(urlEntity)
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
	}
	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}

func (us *URLService) GetByUserID(userId string) ([]models.FullURL, error) {
	result, err := urlRepo.GetByUserID(userId)

	if err != nil {
		return []models.FullURL{}, err
	}

	return result, nil
}
