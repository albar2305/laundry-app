package api

import (
	"github.com/albar2305/enigma-laundry-apps/model"
	"github.com/albar2305/enigma-laundry-apps/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UomController struct {
	uomUC  usecase.UomUseCase
	router *gin.Engine
}

func (u *UomController) createHandler(c *gin.Context) {
	var uom model.Uom
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	uom.Id = uuid.New().String()
	if err := u.uomUC.RegisterNewUom(uom); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	status := map[string]any{
		"code":        200,
		"description": "Create Data Successfully",
	}
	c.JSON(201, gin.H{
		"status": status,
		"data":   uom,
	})

}
func (u *UomController) listHandler(c *gin.Context) {
	uoms, err := u.uomUC.FindAllUom()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   uoms,
	})
}
func (u *UomController) getHandler(c *gin.Context) {
	id := c.Param("id")
	uom, err := u.uomUC.FindByIdUom(id)
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
		"data":   uom,
	})
}

func (u *UomController) updateHandler(c *gin.Context) {
	var uom model.Uom
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}

	if err := u.uomUC.UpdateUom(uom); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}

	status := map[string]any{
		"code":        200,
		"description": "Update Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   uom,
	})
}

func (u *UomController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	err := u.uomUC.DeleteUom(id)

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

func NewUomController(usecase usecase.UomUseCase, r *gin.Engine) *UomController {
	controller := UomController{
		router: r,
		uomUC:  usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/uoms", controller.createHandler)
	rg.GET("/uoms", controller.listHandler)
	rg.GET("/uoms/:id", controller.getHandler)
	rg.PUT("/uoms", controller.updateHandler)
	rg.DELETE("/uoms/:id", controller.deleteHandler)

	return &controller
}
