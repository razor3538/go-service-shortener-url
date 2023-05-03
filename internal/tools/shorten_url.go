package tools

import (
	"github.com/speps/go-hashids"
)

// ShortenURL сокращает url
func ShortenURL(urlModel string) (string, error) {
	hd := hashids.NewData()
	hd.Salt = urlModel

	h, err := hashids.NewWithData(hd)

	if err != nil {
		return "", err
	}

	id, err := h.Encode([]int{1, 2, 3})

	if err != nil {
		return "", err
	}
	return id, nil
}
