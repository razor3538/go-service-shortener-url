package services

import (
	"bytes"
	"example.com/m/v2/config"
	"example.com/m/v2/domain"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestURLService_Save(t *testing.T) {
	config.InitBD()

	type args struct {
		urlModel string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.URL
		wantErr bool
	}{
		{
			name:    "invalid test",
			args:    args{urlModel: "https//googeqwlecom"},
			wantErr: true,
		},
		{
			name:    "invalid test",
			args:    args{urlModel: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &URLService{}
			_, err := us.Save(tt.args.urlModel, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewURLService(t *testing.T) {
	tests := []struct {
		name string
		want *URLService
	}{
		{
			name: "valid test",
			want: &URLService{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewURLService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewURLService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkURLService_SaveMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		us := &URLService{}
		us.Save(uuid.New().String(), "")
	}
}

func ExampleURLService_Save() {
	// Записываем строку в io.Reader
	jsonBody := []byte("https://github.com/gin-contrib/pprof")
	bodyReader := bytes.NewReader(jsonBody)

	// Делаем запрос на сокращение урла
	req, _ := http.Post("http://localhost:8080/", "application/json", bodyReader)

	// Читаем полученный ответ от сервера
	body, _ := ioutil.ReadAll(req.Body)

	println(body)
}
