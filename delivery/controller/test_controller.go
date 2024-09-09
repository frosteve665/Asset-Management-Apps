package controller

import (
	"asetku-bukan-asetmu/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router      *gin.Engine
	testUsecase usecase.TestUsecase
}

func (a *AuthController) testHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]any{
		"message": "Success Create Test Handler",
	})
}

func NewTestController(router *gin.Engine, testUsecase usecase.TestUsecase) {
	controller := &AuthController{
		router:      router,
		testUsecase: testUsecase,
	}

	routerGroup := router.Group("/api/v1/test")
	routerGroup.GET("/", controller.testHandler)
}
