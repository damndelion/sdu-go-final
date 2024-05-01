package http

import (
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"
	"github.com/damndelion/sdu-go-final/internal/controller/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/damndelion/sdu-go-final/internal/usecase"
)

type orderRoutes struct {
	o usecase.Order
	l *logrus.Logger
}

func newOrderRoutes(handler *gin.RouterGroup, o usecase.Order, l *logrus.Logger, key string) {
	r := &orderRoutes{o, l}

	h := handler.Group("/order")
	{
		h.GET("/all", r.getAll)

		h.Use(middleware.AdminVerify(key))
		h.POST("", r.createOrderItem)
		h.PUT("/:id", r.updateOrderItem)
		h.DELETE("/:id", r.deleteOrderItem)

	}
}

// @Summary     Get menu
// @Description Show menu
// @ID          Menu-all
// @Tags  	    Menu
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.Menu
// @Failure     500 {object} response
// @Router      /menu/all [get]
func (r *orderRoutes) getAll(c *gin.Context) {
	order, err := r.o.GetOrder(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - menu")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, order)
}

// @Summary     Create menu item
// @Description Create menu item
// @ID          Menu-create
// @Tags  	    Menu
// @Accept      json
// @Produce     json
// @Success     200 {object} string
// @Failure     500 {object} response
// @Router      /menu [post]
func (r *orderRoutes) createOrderItem(c *gin.Context) {

	var orderItemData dto.CreateOrderItemRequest
	err := c.ShouldBindJSON(&orderItemData)
	userId, _ := c.Get("user_id")
	id, err := r.o.CreateOrderItem(c.Request.Context(), orderItemData, userId.(string))
	if err != nil {
		r.l.Error(err, "http - v1 - order - createOrderItem")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, id)
}

// @Summary     Update menu item
// @Description Update menu item
// @ID          Menu-update
// @Tags  	    Menu
// @Accept      json
// @Produce     json
// @Success     200 {object} string
// @Failure     500 {object} response
// @Router      /menu [put]
func (r *orderRoutes) updateOrderItem(c *gin.Context) {
	id := c.Param("id")

	var orderItemData dto.UpdateOrderItem
	err := c.ShouldBindJSON(&orderItemData)
	orderItemData.ID = id
	id, err = r.o.UpdateOrderItem(c.Request.Context(), orderItemData)
	if err != nil {
		r.l.Error(err, "http - v1 - order - createOrderItem")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, id)
}

// @Summary     Delete menu item
// @Description Delete menu item
// @ID          Menu-delete
// @Tags  	    Menu
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500
// @Router      /menu [delete]
func (r *orderRoutes) deleteOrderItem(c *gin.Context) {
	id := c.Param("id")

	err := r.o.DeleteOrderItem(c.Request.Context(), id)
	if err != nil {
		r.l.Error(err, "http - v1 - order - deleteOrderItem")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, "deleted successfully")
}
