package service

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	"github.com/gin-gonic/gin"
	"go-barbershop/modules/service/model"
	"go-barbershop/utils/constant"
	"go-barbershop/utils/logger"
)

type Repository interface {
	GetListService(ctx *gin.Context, service model.DTOService) (result []model.DTOService, total int64, err error)
	InsertService(ctx *gin.Context, service model.DTOService) (err error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetListService(ctx *gin.Context, _ model.DTOService) (result []model.DTOService, total int64, err error) {
	conn := goqu.New(constant.Postgres.Dialect(), r.db)

	dataset := conn.From(constant.Service.TableName()).
		Select(
			goqu.I("id"),
			goqu.I("name"),
			goqu.I("description"),
			goqu.COALESCE(goqu.I("file_name"), "").As("file_name"),
		)

	err = dataset.ScanStructs(&result)
	if err != nil {
		logger.ErrorWithCtx(ctx, nil, err)
		return
	}

	// counter
	total, err = dataset.ClearSelect().Count()
	if err != nil {
		logger.ErrorWithCtx(ctx, nil, err)
		return
	}

	return
}

func (r *repository) InsertService(ctx *gin.Context, service model.DTOService) (err error) {
	conn := goqu.New(constant.Postgres.Dialect(), r.db)
	dataset := conn.Insert(constant.Service.TableName()).Rows(
		goqu.Record{
			"name":        service.Name,
			"description": service.Description,
		},
	)

	_, err = dataset.Executor().Exec()
	if err != nil {
		logger.ErrorWithCtx(ctx, nil, err)
		return
	}

	return
}
