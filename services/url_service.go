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

func (us *UrlService) Save(urlModel models.ShortenUrl) (domain.Url, error) {
	var urlEntity domain.Url

	_, err := url.ParseRequestURI(urlModel.Url)
	if err != nil {
		return domain.Url{}, err
	}

	u, err := url.Parse(urlModel.Url)
	if err != nil || u.Host == "" {
		return domain.Url{}, err
	}

	shortUrl := make([]byte, 5)

	for i := range shortUrl {
		shortUrl[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	urlEntity.ShortUrl = "http://" + config.Env.Host + "/" + string(shortUrl)
	urlEntity.FullUrl = urlModel.Url

	result, err := geolocationRepo.Save(urlEntity)
	if err != nil {
		return domain.Url{}, err
	}

	return result, nil
}

func (us *UrlService) Get(id string) (domain.Url, error) {
	result, err := geolocationRepo.Get(id)
	if err != nil {
		return domain.Url{}, err
	}

	return result, nil
}
