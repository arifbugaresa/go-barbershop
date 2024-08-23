package service

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go-barbershop/middlewares"
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
	}
}

// GetListRouter godoc
// @Tags Service
// @Summary Get List Service
// @Description	This endpoint is used for get all service
// @Accept json
// @Produce json
// @Success 200 {object} common.APIResponse{data=GetListServiceResponse} "Success"
// @Failure 500	{object} common.APIResponse "Failed"
// @Router /api/services [get]
func GetListRouter(ctx *gin.Context, srv Service) {
	var (
		req GetListServiceRequest
	)

	res, err := srv.GetListService(ctx, req)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "successfully get all data service", res)
}
