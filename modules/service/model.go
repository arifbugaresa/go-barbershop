package service

type GetListServiceRequest struct {
}

type GetListServiceResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
}

type DTOService struct {
	ID          int64  `json:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	FileName    string `db:"file_name"`
	ModifiedBy  string `db:"modified_by"`
	ModifiedAt  string `db:"modified_at"`
	CreatedBy   string `db:"created_by"`
	CreatedAt   string `db:"created_at"`
}
