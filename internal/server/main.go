package server

import (
	"context"
	"example.com/m/v2/internal/app/models"
	"example.com/m/v2/internal/config"
	pb "example.com/m/v2/internal/proto"
	"example.com/m/v2/internal/services"
	"example.com/m/v2/internal/tools"
	"github.com/pkg/errors"
	"strings"
)

// UrlServer поддерживает все необходимые методы сервера.
type UrlServer struct {
	pb.UnimplementedURLsServer
}

var urlService = services.NewURLService()

// ShortenURL реализует интерфейс добавления пользователя.
func (us *UrlServer) ShortenURL(ctx context.Context, in *pb.ShortenURLRequest) (*pb.ShortenURLResponse, error) {
	var response pb.ShortenURLResponse

	save, err := urlService.Save(in.UrlModel, in.UserId)
	if err != nil {
		response.Error = err.Error()
	}

	response.ShortUrl = save.ShortURL

	return &response, nil
}

// DeleteURLs реализует интерфейс добавления пользователя.
func (us *UrlServer) DeleteURLs(ctx context.Context, in *pb.DeleteURLsRequest) (*pb.DeleteURLsResponse, error) {
	var response pb.DeleteURLsResponse

	go urlService.Delete(tools.StringToSlice(in.Urls))

	response.Status = "in process"

	return &response, nil
}

// ReturnFullURL реализует интерфейс добавления пользователя.
func (us *UrlServer) ReturnFullURL(ctx context.Context, in *pb.ReturnFullURLRequest) (*pb.ReturnFullURLResponse, error) {
	var response pb.ReturnFullURLResponse

	url, err := urlService.GetByFullURL(in.Url)

	if err != nil {
		response.Error = err.Error()
		return &response, err
	}
	response.Url = url.FullURL

	return &response, nil
}

// GetFullURL реализует интерфейс добавления пользователя.
func (us *UrlServer) GetFullURL(ctx context.Context, in *pb.GetFullURLRequest) (*pb.GetFullURLResponse, error) {
	var response pb.GetFullURLResponse

	url, err := urlService.Get(in.UserId)

	if err != nil {
		response.Error = err.Error()
		return &response, err
	}

	if url.Deleted.Valid {
		response.Error = "урл удален"
	}

	response.UrlString = url.FullURL

	return &response, nil
}

// Ping реализует интерфейс добавления пользователя.
func (us *UrlServer) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	var response pb.PingResponse

	sqlDB, err := config.DB.DB()
	if err != nil {
		response.Err = err.Error()
		return &response, err
	}

	if err = sqlDB.Ping(); err != nil {
		err := sqlDB.Close()
		if err != nil {
			response.Err = err.Error()
			return &response, err
		}
	}

	response.StatusCode = 200

	return &response, nil
}

// GetAllUsersAndUrls реализует интерфейс добавления пользователя.
func (us *UrlServer) GetAllUsersAndUrls(ctx context.Context, in *pb.GetAllUsersAndUrlsRequest) (*pb.GetAllUsersAndUrlsResponse, error) {
	var response pb.GetAllUsersAndUrlsResponse

	listIp := strings.Split(config.Env.TrustedSubnet, ", ")

	for _, ip := range listIp {
		if ip == in.XRealIp {
			result, err := urlService.GetAllUsersAndUrls()

			if err != nil {
				response.Error = err.Error()
				return &response, err
			}

			response.CountUser = int32(result.UsersCount)
			response.CountUrls = int32(result.UrlsCount)

			return &response, nil
		}
	}

	response.Error = "ip не входит в список доверенных"

	return &response, errors.New("ip не входит в список доверенных")
}

// GetByUserID реализует интерфейс добавления пользователя.
func (us *UrlServer) GetByUserID(ctx context.Context, in *pb.GetByUserIDRequest) (*pb.GetByUserIDResponse, error) {
	var response pb.GetByUserIDResponse

	if in.BearerToken == "" {
		response.Error = "пустой токен"
		return &response, errors.New("пустой токен")
	}

	urlModel, err := urlService.GetByUserID(in.BearerToken)

	if err != nil {
		response.Error = err.Error()
		return &response, err
	}

	for _, url := range urlModel {
		response.ShortUrl = append(response.ShortUrl, url.ShortURL)
		response.FullUrl = append(response.FullUrl, url.FullURL)
	}

	return &response, nil
}

// SaveMany реализует интерфейс добавления пользователя.
func (us *UrlServer) SaveMany(ctx context.Context, in *pb.SaveManyRequest) (*pb.SaveManyResponse, error) {
	var response pb.SaveManyResponse
	var serviceModel []models.SaveBatchURLRequest

	for _, model := range in.UrlsRequest {
		serviceModel = append(serviceModel, models.SaveBatchURLRequest{
			ID:      model.Id,
			FullURL: model.FullUrl,
		})
	}

	urlModel, err := urlService.SaveMany(serviceModel)

	if err != nil {
		response.Error = err.Error()
		return &response, err
	}

	var responseModel *pb.SaveBatchURLRequest

	for _, model := range urlModel {
		responseModel.FullUrl = model.ShortURL
		responseModel.Id = model.ID
		response.UrlsRequest = append(response.UrlsRequest, responseModel)
	}

	return &response, nil
}
