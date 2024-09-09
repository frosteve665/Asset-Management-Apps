package controller

import (
	"asetku-bukan-asetmu/model"
	"asetku-bukan-asetmu/usecase"
	"asetku-bukan-asetmu/utils/common"
	"net/http"

	// "strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	router  *gin.Engine
	useCase usecase.EmployeeUseCase
}

func (e *EmployeeController) createHandler(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	employee.Id = common.GenerateUUID()
	err := e.useCase.RegisterNewEmployee(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success Create New Employee",
		"data":    employee,
	})
}

func (e *EmployeeController) listHandler(c *gin.Context) {
	employees, err := e.useCase.FindAllEmployeeList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Get List Employee",
		"data":    employees,
	})
}

// func (e *EmployeeController) listHandler(c *gin.Context) {
// 	// page, _ := strconv.Atoi(c.Query("page"))
// 	// limit, _ := strconv.Atoi(c.Query("limit"))
// name := c.Query("name")
// paginationParam := dto.PaginationParam{
// 	Page:  page,
// 	Limit: limit,
// }
// employees, paging, err := e.useCase.FindAllEmployee(paginationParam, name)
// if err != nil {
// // 	c.JSON(http.StatusInternalServerError, gin.H{
// // 		"error": err.Error(),
// // 	})
// // 	return
// // }
// status := map[string]any{
// 	"code":        200,
// 	"description": "Get All Data Successfully",
// }
// c.JSON(http.StatusOK, gin.H{
// 	"status": status,
// 	"data":   employees,
// 	"paging": paging,
// })
// }

func (e *EmployeeController) getHandler(c *gin.Context) {
	id := c.Param("id")
	employee, err := e.useCase.FindEmployeeById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success Get Employee by Id",
		"data":    employee,
	})
}

func (e *EmployeeController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := e.useCase.DeleteEmployee(id)
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

func (e *EmployeeController) updateHandler(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	employee.Id = c.Param("id")
	err := e.useCase.UpdateEmployee(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusOK,
		"message": "Success Updated Employee",
		"data":    employee,
	})
}

func NewEmployeeController(router *gin.Engine, emplUseCase usecase.EmployeeUseCase) {
	ctr := &EmployeeController{
		router:  router,
		useCase: emplUseCase,
	}

	// routerGroup := ctr.router.Group("/api/v1", middleware.AuthMiddleware())
	routerGroup := ctr.router.Group("/api/v1/employee")
	routerGroup.POST("/", ctr.createHandler)
	routerGroup.GET("/", ctr.listHandler)
	routerGroup.GET("/:id", ctr.getHandler)
	routerGroup.PUT("/:id", ctr.updateHandler)
	routerGroup.DELETE("/:id", ctr.deleteHandler)
}
