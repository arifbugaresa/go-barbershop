package model

type GetListServiceRequest struct {
}

type GetListServiceResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}
