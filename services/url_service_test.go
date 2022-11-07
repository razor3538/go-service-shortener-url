package services

import (
	"example.com/m/v2/domain"
	"reflect"
	"testing"
)

func TestURLService_Save(t *testing.T) {
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
			name: "valid test",
			args: args{urlModel: "https://google.com"},
			want: domain.URL{
				FullURL: "https://google.com",
			},
			wantErr: false,
		},
		{
			name:    "invalid test",
			args:    args{urlModel: "https//google.com"},
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
			_, err := us.Save(tt.args.urlModel)
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
