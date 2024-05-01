package http

import (
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"
	"github.com/damndelion/sdu-go-final/internal/controller/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/damndelion/sdu-go-final/internal/usecase"
)

type orderMenuRoutes struct {
	om usecase.OrderMenu
	l  *logrus.Logger
}

func newOrderMenuRoutes(handler *gin.RouterGroup, om usecase.OrderMenu, l *logrus.Logger, key string) {
	r := &orderMenuRoutes{om, l}

	h := handler.Group("/order-menu")
	{
		h.GET("/all", r.getAll)

		h.Use(middleware.AdminVerify(key))
		h.PUT("/:id", r.updateOrderMenuItem)
		h.DELETE("/:id", r.deleteOrderMenuItem)

	}
}

func (r *orderMenuRoutes) getAll(c *gin.Context) {
	orderMenu, err := r.om.GetOrderMenu(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - menu")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, orderMenu)
}

func (r *orderMenuRoutes) updateOrderMenuItem(c *gin.Context) {
	id := c.Param("id")

	var orderMenuItemData dto.UpdateOrderMenuItem
	err := c.ShouldBindJSON(&orderMenuItemData)
	orderMenuItemData.ID = id
	id, err = r.om.UpdateOrderMenuItem(c.Request.Context(), orderMenuItemData)
	if err != nil {
		r.l.Error(err, "http - v1 - orderMenu - updateOrderMenuItem")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, id)
}

func (r *orderMenuRoutes) deleteOrderMenuItem(c *gin.Context) {
	id := c.Param("id")

	err := r.om.DeleteOrderMenuItem(c.Request.Context(), id)
	if err != nil {
		r.l.Error(err, "http - v1 - orderMenu - deleteOrderMenuItem")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, "deleted successfully")
}
