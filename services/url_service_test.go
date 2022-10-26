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

func TestURLService_Get(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name         string
		args         args
		want         domain.URL
		wantErr      bool
		wantLocation string
	}{
		{
			name: "valid test",
			args: args{id: "y1IAte"},
			want: domain.URL{
				Base:     domain.Base{ID: 50},
				FullURL:  "https://google.com",
				ShortURL: "http://localhost:8080/y1IAte",
			},
			wantErr:      false,
			wantLocation: "https://google.com",
		},
		{
			name:    "invalid test",
			args:    args{id: "adsgwq1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &URLService{}
			got, err := us.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
