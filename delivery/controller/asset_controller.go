package controller

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/usecase"
	"asetku-bukan-asetmu/utils/common"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AssetController struct {
	router  *gin.Engine
	usecase usecase.AssetUsecase
}

func (a *AssetController) createHandler(ctx *gin.Context) {
	var asset model.Asset
	// Fill asset id and created_at
	asset.Id = common.GenerateUUID()
	asset.CreatedAt = time.Now()
	err := ctx.ShouldBindJSON(&asset)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	if err := a.usecase.CreateNewAsset(asset); err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"status": http.StatusInternalServerError,
			"error":  err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  http.StatusCreated,
		"message": "success register new asset",
	})
}

func (a *AssetController) listHandler(ctx *gin.Context) {
	assets, err := a.usecase.ShowAllAsset()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"message": err.Error(),
		})
		return
	}

	if len(assets) == 0 {
		ctx.JSON(http.StatusNoContent, map[string]any{
			"message": "asset is empty",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": "success show all assets",
		"data":    assets,
	})
}

func (a *AssetController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	assets, err := a.usecase.GetDetailAsset(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status": "success",
		"data":   assets,
	})
}

func (a *AssetController) placementHandler(ctx *gin.Context) {
	var assetPlacement model.AssetPlacement
	assetPlacement.UpdatedAt = time.Now()
	assetPlacement.CurrentStatus = 1
	assetPlacement.TargetStatus = 2
	err := ctx.ShouldBindJSON(&assetPlacement)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		})
		return
	}

	available, err := a.usecase.UpdateAssetLocation(assetPlacement)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	if len(available) == 0 {
		ctx.JSON(http.StatusAccepted, map[string]any{
			"status":  "failed",
			"message": "qty is beyond more current asset",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success change placement of asset",
	})
}

func NewAssetController(router *gin.Engine, assetUsecase usecase.AssetUsecase) {
	controller := &AssetController{
		router:  router,
		usecase: assetUsecase,
	}

	routerGroup := controller.router.Group("/api/v1/asset")
	routerGroup.POST("/", controller.createHandler)
	routerGroup.GET("/", controller.listHandler)
	routerGroup.GET("/detail/:id", controller.getHandler)
	routerGroup.PUT("/placement/:id", controller.placementHandler)
}
