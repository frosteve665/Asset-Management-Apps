package controller

import (
	"net/http"

	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/usecase"
	"asetku-bukan-asetmu/utils/common"

	"github.com/gin-gonic/gin"
)

type AssetCategoriesController struct {
	router  *gin.Engine
	Usecase usecase.AssetCategoriesUseCase
}

func (a *AssetCategoriesController) createHandler(c *gin.Context) {
	var assetcategories model.AssetCategories
	assetcategories.Id = common.GenerateUUID()
	if err := c.ShouldBindJSON(&assetcategories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := a.Usecase.RegisterNewAssetCategories(assetcategories)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success Create New Asset Categories",
		"data":    assetcategories,
	})
}

func (a *AssetCategoriesController) searchHandler(c *gin.Context) {
	id := c.Param("id")
	assetcategories, err := a.Usecase.FindAssetCategoriesById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Get assetvategories by Id",
		"data":    assetcategories,
	})
}

func (a *AssetCategoriesController) listHandler(ctx *gin.Context) {
	locations, err := a.Usecase.FindAllAssetCategoriesList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"status":  http.StatusOK,
		"message": "Success Show All Locations",
		"data":    locations,
	})
}

func (a *AssetCategoriesController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := a.Usecase.DeleteAssetCategories(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Delete",
	})
}

func (a *AssetCategoriesController) updateHandler(c *gin.Context) {
	var assetcategories model.AssetCategories
	if err := c.ShouldBindJSON(&assetcategories); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	assetcategories.Id = c.Param("id")
	err := a.Usecase.UpdateAssetCategories(assetcategories)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success Updated asset categories",
		"data":    assetcategories,
	})
}

// func (a *AssetCategoriesController) listHandler(c *gin.Context) {
// 	page, _ := strconv.Atoi(c.Query("page"))
// 	limit, _ := strconv.Atoi(c.Query("limit"))
// 	name := c.Query("name")
// 	paginationParam := dto.PaginationParam{
// 		Page : page,
// 		Limit : limit,
// 	}
// 	assetcategories,paging, err := a.Usecase.FindAllAssetcategories(paginationParam,name)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error" : err.Error(),
// 		})
// 		return
// 	}
// 	status := map[string]any{
// 		"code":        200,
// 		"description": "Get All Data Successfully",
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": status,
// 		"data":   assetcategories,
// 		"paging": paging,
// 	})
// }

func NewAssetCategoriesController(router *gin.Engine, assetcatagoriesUseCase usecase.AssetCategoriesUseCase) {
	ctr := &AssetCategoriesController{
		router:  router,
		Usecase: assetcatagoriesUseCase,
	}

	routerGroup := ctr.router.Group("/api/v1/asset-category")
	routerGroup.POST("/", ctr.createHandler)
	routerGroup.GET("/", ctr.listHandler)
	routerGroup.GET("/:id", ctr.searchHandler)
	routerGroup.PUT("/:id", ctr.updateHandler)
	routerGroup.DELETE("//:id", ctr.deleteHandler)
}
