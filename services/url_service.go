package services

import (
	"example.com/m/v2/domain"
	"example.com/m/v2/internal/app/models"
	"example.com/m/v2/repositories"
	"github.com/speps/go-hashids"
)

type URLService struct{}

func NewURLService() *URLService {
	return &URLService{}
}

var geolocationRepo = repositories.NewURLRepo()

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (us *URLService) Save(urlModel models.ShortenURL) (domain.URL, error) {
	var urlEntity domain.URL

	hd := hashids.NewData()
	hd.Salt = urlModel.URL

	h, err := hashids.NewWithData(hd)

	id, _ := h.Encode([]int{1, 2, 3})

	urlEntity.ShortURL = "http://localhost:8080/" + id
	urlEntity.FullURL = urlModel.URL

	result, err := geolocationRepo.Save(urlEntity)
	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}

func (us *URLService) Get(id string) (domain.URL, error) {
	result, err := geolocationRepo.Get("http://localhost:8080/" + id)
	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}
