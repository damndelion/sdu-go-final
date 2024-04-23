package http

import (
	"errors"
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"
	"github.com/damndelion/sdu-go-final/internal/controller/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"

	"github.com/damndelion/sdu-go-final/internal/usecase"
)

type menuRoutes struct {
	t usecase.Menu
	l *logrus.Logger
}

func newMenuRoutes(handler *gin.RouterGroup, t usecase.Menu, l *logrus.Logger, key string) {
	r := &menuRoutes{t, l}

	h := handler.Group("/menu")
	{
		h.GET("/all", r.getAll)

		h.Use(middleware.AdminVerify(key))
		h.POST("", r.createMenuItem)
		h.PUT("/:id", r.updateMenuItem)
		h.DELETE("/:id", r.deleteMenuItem)

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
func (r *menuRoutes) getAll(c *gin.Context) {
	menu, err := r.t.GetMenu(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - menu")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, menu)
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
func (r *menuRoutes) createMenuItem(c *gin.Context) {

	var menuItemData dto.CreateMenuItem
	err := c.ShouldBindJSON(&menuItemData)

	id, err := r.t.CreateMenuItem(c.Request.Context(), menuItemData)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			r.l.Error(err, "http - v1 - menu - createMenuItem - duplicate name")
			errorResponse(c, http.StatusConflict, "menu with that name already exists")
			return
		}
		r.l.Error(err, "http - v1 - menu - createMenuItem")
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
func (r *menuRoutes) updateMenuItem(c *gin.Context) {
	id := c.Param("id")

	var menuItemData dto.UpdateMenuItem
	err := c.ShouldBindJSON(&menuItemData)
	menuItemData.ID = id

	id, err = r.t.UpdateMenuItem(c.Request.Context(), menuItemData)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			r.l.Error(err, "http - v1 - menu - createMenuItem - duplicate name")
			errorResponse(c, http.StatusConflict, "menu with that name already exists")
			return
		}
		r.l.Error(err, "http - v1 - menu - createMenuItem")
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
func (r *menuRoutes) deleteMenuItem(c *gin.Context) {
	id := c.Param("id")

	err := r.t.DeleteMenuItem(c.Request.Context(), id)
	if err != nil {
		r.l.Error(err, "http - v1 - menu - deleteMenuItem")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, "deleted successfully")
}
