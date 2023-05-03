package services

import (
	"reflect"
	"testing"

	"example.com/m/v2/internal/config"
	"example.com/m/v2/internal/domain"
	"github.com/google/uuid"
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

		_, err := us.Save(uuid.New().String(), "")
		if err != nil {
			return
		}
	}
}
