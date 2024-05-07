package http

import (
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"
	"github.com/damndelion/sdu-go-final/internal/controller/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/damndelion/sdu-go-final/internal/usecase"
)

type stockRoutes struct {
	s usecase.Stock
	l *logrus.Logger
}

func newStockRoutes(handler *gin.RouterGroup, s usecase.Stock, l *logrus.Logger, key string) {
	r := &stockRoutes{s, l}

	h := handler.Group("/stock")
	{
		// Only admin can access
		h.Use(middleware.AdminVerify(key))
		h.GET("/all", r.getAll)
		h.POST("", r.createStockItem)
		h.PUT("/:id", r.updateStockItem)
		h.DELETE("/:id", r.deleteStockItem)

	}
}

func (r *stockRoutes) getAll(c *gin.Context) {
	stock, err := r.s.GetStock(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - stock")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, stock)
}

func (r *stockRoutes) createStockItem(c *gin.Context) {

	var stockItemData dto.CreateStockItem
	err := c.ShouldBindJSON(&stockItemData)
	id, err := r.s.CreateStockItem(c.Request.Context(), stockItemData)
	if err != nil {
		r.l.Error(err, "http - v1 - stock - createStockItem")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, id)
}

func (r *stockRoutes) updateStockItem(c *gin.Context) {
	id := c.Param("id")

	var stockItemData dto.UpdateStockItem
	err := c.ShouldBindJSON(&stockItemData)
	stockItemData.ID = id
	id, err = r.s.UpdateStockItem(c.Request.Context(), stockItemData)
	if err != nil {
		r.l.Error(err, "http - v1 - stock - createStockItem")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, id)
}

func (r *stockRoutes) deleteStockItem(c *gin.Context) {
	id := c.Param("id")

	err := r.s.DeleteStockItem(c.Request.Context(), id)
	if err != nil {
		r.l.Error(err, "http - v1 - stock - deleteStockItem")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, "deleted successfully")
}
