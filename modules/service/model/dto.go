package model

import "time"

type DTOService struct {
	ID          int64     `db:"id" goqu:"skipinsert"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	FileName    string    `db:"file_name"`
	ModifiedBy  string    `db:"modified_by"`
	ModifiedAt  time.Time `db:"modified_at"`
	CreatedBy   string    `db:"created_by"`
	CreatedAt   time.Time `db:"created_at"`
}
