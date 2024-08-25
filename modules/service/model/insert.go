package model

import (
	"errors"
	"go-barbershop/utils/common"
)

type InsertServiceResponse struct{}

type InsertServiceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

func (req *InsertServiceRequest) Validate() error {
	if common.IsEmptyField(req.Name) {
		return errors.New("name is required")
	}

	if common.IsEmptyField(req.Description) {
		return errors.New("description is required")
	}

	return nil
}
