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
	req, _ := http.Post("http://localhost:8080/", "application/json", bodyReader)

	// Читаем полученный ответ от сервера
	body, _ := io.ReadAll(req.Body)

	println(body)

	req.Body.Close()
}
