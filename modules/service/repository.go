package service

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"go-barbershop/utils/constant"
	"go-barbershop/utils/logger"
)

type Repository interface {
	GetListService(ctx *gin.Context, user GetListServiceRequest) (result []DTOService, err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetListService(ctx *gin.Context, _ GetListServiceRequest) (result []DTOService, err error) {
	conn := goqu.New(constant.Postgres.Dialect(), r.db)

	dataset := conn.From(constant.Service.TableName()).
		Select(
			goqu.C("id"),
			goqu.C("name"),
			goqu.C("description"),
			goqu.C("file_name"),
		)

	err = dataset.ScanStructs(&result)
	if err != nil {
		logger.ErrorWithCtx(ctx, nil, err)
		return
	}

	return
}
