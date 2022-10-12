package services

import (
	"example.com/m/v2/config"
	"example.com/m/v2/domain"
	"example.com/m/v2/internal/app/models"
	"example.com/m/v2/repositories"
	"math/rand"
	"net/url"
)

type UrlService struct{}

func NewUrlService() *UrlService {
	return &UrlService{}
}

var geolocationRepo = repositories.NewUrlRepo()

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (us *UrlService) Save(urlModel models.ShortenURL) (domain.URL, error) {
	var urlEntity domain.URL

	_, err := url.ParseRequestURI(urlModel.URL)
	if err != nil {
		return domain.URL{}, err
	}

	u, err := url.Parse(urlModel.URL)
	if err != nil || u.Host == "" {
		return domain.URL{}, err
	}

	shortUrl := make([]byte, 5)

	for i := range shortUrl {
		shortUrl[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	urlEntity.ShortURL = "http://" + config.Env.Host + "/" + string(shortUrl)
	urlEntity.FullURL = urlModel.URL

	result, err := geolocationRepo.Save(urlEntity)
	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}

func (us *UrlService) Get(id string) (domain.URL, error) {
	result, err := geolocationRepo.Get(id)
	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}
