package service

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetListService(ctx *gin.Context, req GetListServiceRequest) (result []GetListServiceResponse, err error)

	ConvertDTOToGetListResponse(services []DTOService) (resp []GetListServiceResponse)
}

type userService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &userService{
		repository,
	}
}

func (s *userService) GetListService(ctx *gin.Context, req GetListServiceRequest) (result []GetListServiceResponse, err error) {
	services, err := s.repository.GetListService(ctx, req)
	if err != nil {
		err = errors.New("failed get list service")
		return
	}

	result = s.ConvertDTOToGetListResponse(services)

	return
}

func (s *userService) ConvertDTOToGetListResponse(services []DTOService) (resp []GetListServiceResponse) {
	for _, item := range services {
		resp = append(resp, GetListServiceResponse{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			ImageUrl:    item.FileName,
		})
	}

	return
}
