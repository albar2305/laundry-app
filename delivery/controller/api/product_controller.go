package api

import (
	"net/http"
	"strconv"

	"github.com/albar2305/enigma-laundry-apps/model"
	"github.com/albar2305/enigma-laundry-apps/model/dto"
	"github.com/albar2305/enigma-laundry-apps/usecase"
	"github.com/albar2305/enigma-laundry-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router    *gin.Engine
	productUC usecase.ProductUseCase
}

func (p *ProductController) createHandler(c *gin.Context) {
	var product dto.ProductRequestDto
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var newProduct model.Product
	product.Id = common.GenerateID()
	newProduct.Id = product.Id
	newProduct.Name = product.Name
	newProduct.Price = product.Price
	newProduct.Uom.Id = product.UomId
	if err := p.productUC.RegisterNewProduct(newProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	status := map[string]any{
		"code":        http.StatusCreated,
		"description": "Create Data Successfully",
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": status,
		"data":   product,
	})
}
func (p *ProductController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginationParams := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	products, paging, err := p.productUC.FindAllProduct(paginationParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        http.StatusOK,
		"description": "Get All Data Successfully",
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   products,
		"paging": paging,
	})
}
func (p *ProductController) getHandler(c *gin.Context) {
	id := c.Param("id")
	product, err := p.productUC.FindByIdProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	status := map[string]any{
		"code":        http.StatusOK,
		"description": "Get By ID Data Successfully",
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   product,
	})
}
func (p *ProductController) updateHandler(c *gin.Context) {
	var product dto.ProductRequestDto
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var newProduct model.Product
	newProduct.Id = product.Id
	newProduct.Name = product.Name
	newProduct.Price = product.Price
	newProduct.Uom.Id = product.UomId
	if err := p.productUC.UpdateProduct(newProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	status := map[string]any{
		"code":        http.StatusOK,
		"description": "Update Data Successfully",
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   product,
	})
}
func (p *ProductController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := p.productUC.DeleteProduct(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	status := map[string]any{
		"code":        http.StatusNoContent,
		"description": "Delete Data Successfully",
	}
	c.JSON(http.StatusNoContent, gin.H{
		"status": status,
	})
}

func NewProductController(r *gin.Engine, usecase usecase.ProductUseCase) *ProductController {
	controller := ProductController{
		router:    r,
		productUC: usecase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/products", controller.createHandler)
	rg.GET("/products", controller.listHandler)
	rg.GET("/products/:id", controller.getHandler)
	rg.PUT("/products", controller.updateHandler)
	rg.DELETE("/products/:id", controller.deleteHandler)

	return &controller

}
