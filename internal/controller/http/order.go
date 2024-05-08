package http

import (
	"fmt"
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
		// All authorized users can access
		h.Use(middleware.JwtVerify(key))
		h.GET("/user/all", r.getUserAllOrders)
		h.GET("/user/current", r.getUserCurrentOrders)
		h.POST("", r.createOrderItem)
		h.PUT("/:id", r.updateOrderItem)
		h.DELETE("/:id", r.deleteOrderItem)

		// worker can access
		h.Use(middleware.WorkerVerify(key))
		h.GET("/all", r.getAll)
		h.GET("/all/current", r.getAllCurrent)

	}
}

// @Summary     Get all Orders
// @Description Get all Orders
// @ID          Order-all
// @Tags  	    Order
// @Accept      json
// @Produce     json
// @Success     200 {object} []entity.Order
// @Failure     500 {object} response
// @Router      /order/all [get]
func (r *orderRoutes) getAll(c *gin.Context) {
	order, err := r.o.GetOrder(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - menu")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, order)
}

// @Summary     Get all current order
// @Description Get all current order
// @ID          Order-all-current
// @Tags  	    Order
// @Accept      json
// @Produce     json
// @Success     200 {object} []entity.Order
// @Failure     500 {object} response
// @Router      /order/all/current [get]
func (r *orderRoutes) getAllCurrent(c *gin.Context) {
	order, err := r.o.GetCurrentOrder(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - menu")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, order)
}

// @Summary     Get all current user order
// @Description Get all current user order
// @ID          Order-all-current-user
// @Tags  	    Order
// @Accept      json
// @Produce     json
// @Success     200 {object} []entity.Order
// @Failure     500 {object} response
// @Router      /order/user/current [get]
func (r *orderRoutes) getUserCurrentOrders(c *gin.Context) {
	userId, _ := c.Get("user_id")
	order, err := r.o.GetUserCurrentOrder(c.Request.Context(), userId.(string))
	if err != nil {
		r.l.Error(err, "http - v1 - menu")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, order)
}

// @Summary     Get all  user order
// @Description Get all  user order
// @ID          Order-all--user
// @Tags  	    Order
// @Accept      json
// @Produce     json
// @Success     200 {object} []entity.Order
// @Failure     500 {object} response
// @Router      /order/user/all [get]
func (r *orderRoutes) getUserAllOrders(c *gin.Context) {
	userId, _ := c.Get("user_id")
	order, err := r.o.GetUserCurrentOrder(c.Request.Context(), userId.(string))
	if err != nil {
		r.l.Error(err, "http - v1 - menu")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, order)
}

// @Summary     Create Order
// @Description Create Order
// @ID          Order-create
// @Tags  	    Order
// @Accept      json
// @Produce     json
// @Success     200 {object} string
// @Failure     500 {object} response
// @Router      /order [post]
func (r *orderRoutes) createOrderItem(c *gin.Context) {

	var orderItemData dto.CreateOrderItemRequest
	err := c.ShouldBindJSON(&orderItemData)
	userId, _ := c.Get("user_id")
	fmt.Println(orderItemData)
	id, err := r.o.CreateOrderItem(c.Request.Context(), orderItemData, userId.(string))
	if err != nil {
		r.l.Error(err, "http - v1 - order - createOrderItem")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, id)
}

// @Summary     Update order
// @Description Update order
// @ID          Order-update
// @Tags  	    Order
// @Accept      json
// @Produce     json
// @Success     200 {object} string
// @Failure     500 {object} response
// @Router      /order/{id} [put]
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

// @Summary     Delete order
// @Description Delete order
// @ID          Order-delete
// @Tags  	    Order
// @Accept      json
// @Produce     json
// @Success     200
// @Failure     500
// @Router      /order/{id} [delete]
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
