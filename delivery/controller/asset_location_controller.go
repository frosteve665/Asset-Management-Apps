package controller

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/usecase"
	"asetku-bukan-asetmu/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssetLocationController struct {
	router  *gin.Engine
	usecase usecase.AssetLocationUsecase
}

func (loc *AssetLocationController) createHandler(ctx *gin.Context) {
	var location model.AssetLocation
	location.Id = common.GenerateUUID()
	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
		return
	}

	err := loc.usecase.RegisterNewLocation(location)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, map[string]any{
		"status":  http.StatusCreated,
		"message": "success add new location",
	})
}

func (loc *AssetLocationController) showHandler(ctx *gin.Context) {
	locations, err := loc.usecase.ShowAllLocation()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})
		return
	}

	if len(locations) == 0 {
		ctx.JSON(http.StatusNoContent, map[string]any{
			"message": "location is empty",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": "success show all locations",
		"data":    locations,
	})
}

func (loc *AssetLocationController) searchHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	location, err := loc.usecase.SearchLocationById(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})

		return
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": "success search location",
		"data":    location,
	})
}

func (loc *AssetLocationController) updateHandler(ctx *gin.Context) {
	var location model.AssetLocation
	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"error": err.Error(),
		})
		return
	}

	location.Id = ctx.Param("id")
	err := loc.usecase.EditExistedLocation(location)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": "success update existed location",
	})
}

func (loc *AssetLocationController) deleteHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	err := loc.usecase.DeleteSelectedLocation(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"message": "success delete location",
		"status":  http.StatusOK,
	})
}

func NewAssetLocationController(router *gin.Engine, assetLocUsecase usecase.AssetLocationUsecase) *AssetLocationController {
	controller := &AssetLocationController{
		router:  router,
		usecase: assetLocUsecase,
	}

	routerGroup := controller.router.Group("/api/v1/asset-location")
	routerGroup.POST("/", controller.createHandler)
	routerGroup.GET("/", controller.showHandler)
	routerGroup.GET("/:id", controller.searchHandler)
	routerGroup.PUT("/:id", controller.updateHandler)
	routerGroup.DELETE("/:id", controller.deleteHandler)

	return controller
}
