package api

import (
	"net/http"
	"strconv"

	"github.com/albar2305/enigma-laundry-apps/model"
	"github.com/albar2305/enigma-laundry-apps/model/dto"
	"github.com/albar2305/enigma-laundry-apps/usecase"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router     *gin.Engine
	customerUC usecase.CustomerUseCase
}

func (customerController *CustomerController) createHandler(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := customerController.customerUC.RegisterNewCustomer(customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, customer)
}
func (customerController *CustomerController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginationParams := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	customers, paging, err := customerController.customerUC.FindAllCustomer(paginationParams)
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
		"data":   customers,
		"paging": paging,
	})
}
func (customerController *CustomerController) getHandler(c *gin.Context) {
	id := c.Param("id")
	customer, err := customerController.customerUC.FindByIdCustomer(id)
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
		"data":   customer,
	})
}
func (customerController *CustomerController) updateHandler(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	if err := customerController.customerUC.UpdateCustomer(customer); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	status := map[string]any{
		"code":        200,
		"description": "Update Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   customer,
	})
}
func (customerController *CustomerController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := customerController.customerUC.DeleteCustomer(id)

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

func NewCustomerController(r *gin.Engine, usecase usecase.CustomerUseCase) *CustomerController {
	controller := CustomerController{
		router:     r,
		customerUC: usecase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/customers", controller.createHandler)
	rg.GET("/customers", controller.listHandler)
	rg.GET("/customers/:id", controller.getHandler)
	rg.PUT("/customers", controller.updateHandler)
	rg.DELETE("/customers/:id", controller.deleteHandler)

	return &controller

}
