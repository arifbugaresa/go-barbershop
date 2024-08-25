package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-barbershop/modules/service/model"
)

type Service interface {
	GetListService(ctx *gin.Context, req model.GetListServiceRequest) (result []model.GetListServiceResponse, total int64, err error)
	InsertService(ctx *gin.Context, req model.InsertServiceRequest) (err error)

	ConvertDTOToGetListResponse(services []model.DTOService) (resp []model.GetListServiceResponse)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
	}
}

func (s *userService) GetListService(ctx *gin.Context, req model.GetListServiceRequest) (result []model.GetListServiceResponse, total int64, err error) {
	services, total, err := s.repository.GetListService(ctx, model.DTOService{})
	if err != nil {
		err = errors.New("failed get list service")
		return
	}

	result = s.ConvertDTOToGetListResponse(services)

	return
}

func (s *userService) ConvertDTOToGetListResponse(services []model.DTOService) (resp []model.GetListServiceResponse) {
	for _, item := range services {
		resp = append(resp, model.GetListServiceResponse{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			ImageUrl:    item.FileName,
		})
	}

	return
}

func (s *userService) InsertService(ctx *gin.Context, req model.InsertServiceRequest) (err error) {
	err = s.repository.InsertService(ctx, model.DTOService{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		err = errors.New("failed insert service")
		return
	}

	return
}
