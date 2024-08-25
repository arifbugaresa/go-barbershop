package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go-barbershop/middlewares"
	"go-barbershop/modules/service/model"
	"go-barbershop/utils/common"
)

func Initiator(router *gin.Engine, dbConnection *sql.DB) {
	var (
		repo = NewRepository(dbConnection)
		srv  = NewService(repo)
	)

	api := router.Group("/api/services")
	api.Use(middlewares.Logging())
	{
		api.GET("", func(c *gin.Context) {
			GetListRouter(c, srv)
		})

		api.POST("", func(c *gin.Context) {
			InsertService(c, srv)
		})
	}
}

// GetListRouter godoc
// @Tags Service
// @Summary Get List Service
// @Description	This endpoint is used for get all service
// @Accept json
// @Produce json
// @Success 200 {object} common.APIResponse{data=model.GetListServiceResponse} "Success"
// @Failure 500	{object} common.APIResponse "Failed"
// @Router /api/services [get]
func GetListRouter(ctx *gin.Context, srv Service) {
	var (
		req model.GetListServiceRequest
	)

	res, total, err := srv.GetListService(ctx, req)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithListData(ctx, "successfully get all data service", total, res)
}

// InsertService godoc
// @Tags Service
// @Summary Insert Service
// @Description	This endpoint is used for get all service
// @Accept json
// @Produce json
// @Param request body model.InsertServiceRequest true "Request body"
// @Failure 500	{object} common.APIResponse "Failed"
// @Router /api/services [post]
func InsertService(ctx *gin.Context, srv Service) {
	var (
		req model.InsertServiceRequest
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	err = srv.InsertService(ctx, req)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponse(ctx, "successfully insert data service")
}
