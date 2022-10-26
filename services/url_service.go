package services

import (
	"encoding/json"
	"errors"
	"example.com/m/v2/domain"
	"example.com/m/v2/repositories"
	"github.com/speps/go-hashids"
	"net/url"
	"os"
)

type URLService struct{}

func NewURLService() *URLService {
	return &URLService{}
}

var address = os.Getenv("SERVER_ADDRESS")

var geolocationRepo = repositories.NewURLRepo()

func (us *URLService) Save(urlModel string) (domain.URL, error) {
	var urlEntity domain.URL

	if address == "" {
		address = "localhost:8080"
	}

	_, err := url.ParseRequestURI(urlModel)

	if err != nil {
		return domain.URL{}, errors.New("не валидный URL")
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

	result, err := geolocationRepo.Save(urlEntity)
	if err != nil {
		return domain.URL{}, err
	}

	if os.Getenv("FILE_STORAGE_PATH") != "" {
		file, err := os.OpenFile(os.Getenv("FILE_STORAGE_PATH"), os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)

		//dir, fileName, _ := strings.Cut(os.Getenv("FILE_STORAGE_PATH"), "/")
		//err := os.MkdirAll(dir, 0750)
		//if err != nil {
		//	return domain.URL{}, err
		//}
		//file, err := os.OpenFile(dir+"/"+fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
		//
		//if err != nil {
		//	return domain.URL{}, err
		//}

		data, err := json.Marshal(result)
		if err != nil {
			return domain.URL{}, err
		}

		data = append(data, '\n')

		_, err = file.Write(data)
		if err != nil {
			return domain.URL{}, err
		}
	}

	return result, nil
}

func (us *URLService) Get(id string) (domain.URL, error) {
	if address == "" {
		address = "localhost:8080"
	}

	result, err := geolocationRepo.Get("http://" + address + "/" + id)
	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}

func (us *URLService) GetByFullURL(url string) (domain.URL, error) {
	result, err := geolocationRepo.GetByFullURL(url)

	if result.FullURL == "" {
		urlModel, err := us.Save(url)
		if err != nil {
			return domain.URL{}, err
		}
		result, err = geolocationRepo.GetByFullURL(urlModel.FullURL)
		if err != nil {
			return domain.URL{}, err
		}
	}
	if err != nil {
		return domain.URL{}, err
	}

	return result, nil
}
