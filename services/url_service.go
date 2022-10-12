package services

import (
	"example.com/m/v2/domain"
	"example.com/m/v2/internal/app/models"
	"example.com/m/v2/repositories"
	"math/rand"
)

type URLService struct{}

func NewURLService() *URLService {
	return &URLService{}
}

var geolocationRepo = repositories.NewURLRepo()

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (us *URLService) Save(urlModel models.ShortenURL) (domain.URL, error) {
	var urlEntity domain.URL

	//_, err := url.ParseRequestURI(urlModel.URL)
	//if err != nil {
	//	return domain.URL{}, err
	//}
	//
	//u, err := url.Parse(urlModel.URL)
	//if err != nil || u.Host == "" {
	//	return domain.URL{}, err
	//}

	shortURL := make([]byte, 5)

	for i := range shortURL {
		shortURL[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	urlEntity.ShortURL = "http://" + "localhost/" + string(shortURL)
	urlEntity.FullURL = urlModel.URL

	result, err := geolocationRepo.Save(urlEntity)
	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}

func (us *URLService) Get(id string) (domain.URL, error) {
	result, err := geolocationRepo.Get(id)
	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}
