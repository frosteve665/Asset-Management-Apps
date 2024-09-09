package controller

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/usecase"
	"asetku-bukan-asetmu/utils/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VendorController struct {
	Router        *gin.Engine
	vendorUsecase usecase.VendorUsecase
}

func (c *VendorController) Create(ctx *gin.Context) {
	var vendor model.Vendor
	vendor.Id = common.GenerateUUID()
	err := ctx.ShouldBindJSON(&vendor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = c.vendorUsecase.Create(vendor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
	})
}

func (c *VendorController) List(ctx *gin.Context) {
	vendors, err := c.vendorUsecase.List()
	if len(vendors) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "data not found",
		})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   vendors,
	})
}

func (c *VendorController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	vendor, err := c.vendorUsecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   vendor,
	})
}

func (c *VendorController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	var vendor model.Vendor
	vendor.Id = id

	_, err := c.vendorUsecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&vendor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = c.vendorUsecase.Update(vendor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (c *VendorController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := c.vendorUsecase.Get(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	err = c.vendorUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func NewVendorController(router *gin.Engine, vendorUsecase usecase.VendorUsecase) *VendorController {
	controller := &VendorController{
		Router:        router,
		vendorUsecase: vendorUsecase,
	}

	routerGroup := controller.Router.Group("/api/v1")
	routerGroup.POST("/vendor", controller.Create)
	routerGroup.GET("/vendor", controller.List)
	routerGroup.GET("/vendor/:id", controller.Get)
	routerGroup.PUT("/vendor/:id", controller.Update)
	routerGroup.DELETE("/vendor/:id", controller.Delete)
	return controller
}
