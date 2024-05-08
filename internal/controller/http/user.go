package http

import (
	"github.com/damndelion/sdu-go-final/internal/controller/middleware"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/damndelion/sdu-go-final/internal/usecase"
)

type userRoutes struct {
	t usecase.User
	l *logrus.Logger
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.User, l *logrus.Logger, key string) {
	r := &userRoutes{t, l}

	h := handler.Group("/user")
	{

		h.Use(middleware.JwtVerify(key))
		h.GET("/profile", r.getAll)

		h.Use(middleware.AdminVerify(key))
		h.GET("/all", r.getAll)

	}
}

// @Summary     Get all users
// @Description Show all users
// @ID          User-all
// @Tags  	    User
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.User
// @Failure     500 {object} response
// @Router      /user/all [get]
func (r *userRoutes) getAll(c *gin.Context) {
	users, err := r.t.GetUsers(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - user")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary     Get user profile
// @Description Show user's personal profile information
// @ID          User-by-id
// @Tags  	    User
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.User
// @Failure     500 {object} response
// @Router      /user/{id} [get]
func (r *userRoutes) getById(c *gin.Context) {
	userId, _ := c.Get("user_id")
	users, err := r.t.GetUserByID(c.Request.Context(), userId.(string))
	if err != nil {
		r.l.Error(err, "http - v1 - user")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, users)
}
