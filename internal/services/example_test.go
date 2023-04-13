package services

import (
	"bytes"
	"example.com/m/v2/internal/tools"
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
		tools.ErrorLog.Println(err)
	}

	// Читаем полученный ответ от сервера
	body, err := io.ReadAll(req.Body)

	if err != nil {
		tools.ErrorLog.Println(err)
	}

	tools.ErrorLog.Println(body)

	err = req.Body.Close()

	if err != nil {
		return
	}
}
