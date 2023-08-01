package api

import (
	"net/http"
	"strconv"

	"github.com/albar2305/enigma-laundry-apps/model"
	"github.com/albar2305/enigma-laundry-apps/model/dto"
	"github.com/albar2305/enigma-laundry-apps/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EmployeeController struct {
	router     *gin.Engine
	employeeUC usecase.EmployeeUseCase
}

func (e *EmployeeController) createHandler(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	employee.Id = uuid.New().String()
	if err := e.employeeUC.RegisterNewEmployee(employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, employee)
}
func (e *EmployeeController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginationParams := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	employees, paging, err := e.employeeUC.FindAllEmployee(paginationParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   employees,
		"paging": paging,
	})
}
func (e *EmployeeController) getHandler(c *gin.Context) {
	id := c.Param("id")
	employee, err := e.employeeUC.FindByIdEmployee(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	status := map[string]any{
		"code":        200,
		"description": "Get By ID Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   employee,
	})
}
func (e *EmployeeController) updateHandler(c *gin.Context) {
	var employee model.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	if err := e.employeeUC.UpdateEmployee(employee); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	status := map[string]any{
		"code":        200,
		"description": "Update Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   employee,
	})
}
func (e *EmployeeController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := e.employeeUC.DeleteEmployee(id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	status := map[string]any{
		"code":        200,
		"description": "Delete Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
	})
}

func NewEmployeeController(r *gin.Engine, usecase usecase.EmployeeUseCase) *EmployeeController {
	controller := EmployeeController{
		router:     r,
		employeeUC: usecase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/employees", controller.createHandler)
	rg.GET("/employees", controller.listHandler)
	rg.GET("/employees/:id", controller.getHandler)
	rg.PUT("/employees", controller.updateHandler)
	rg.DELETE("/employees/:id", controller.deleteHandler)

	return &controller

}
