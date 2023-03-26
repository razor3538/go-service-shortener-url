package services

import (
	"bytes"
	"io"
	"net/http"
)

func ExampleURLService_Save() {
	// Записываем строку в io.Reader
	jsonBody := []byte("https://github.com/gin-contrib/pprof")
	bodyReader := bytes.NewReader(jsonBody)

	// Делаем запрос на сокращение урла
	req, err := http.Post("http://localhost:8080/", "application/json", bodyReader)

	if err != nil {
		println(err)
	}

	// Читаем полученный ответ от сервера
	body, err := io.ReadAll(req.Body)

	if err != nil {
		println(err)
	}

	println(body)

	err = req.Body.Close()
	
	if err != nil {
		return
	}
}
