package tools

import (
	"github.com/speps/go-hashids"
)

func ShortenUrl(urlModel string) (string, error) {
	hd := hashids.NewData()
	hd.Salt = urlModel

	h, err := hashids.NewWithData(hd)

	if err != nil {
		return "", err
	}

	id, _ := h.Encode([]int{1, 2, 3})
	return id, nil
}
