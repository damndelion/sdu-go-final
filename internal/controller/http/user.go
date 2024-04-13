package http

import (
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/evrone/go-clean-template/internal/usecase"
)

type userRoutes struct {
	t usecase.User
	l *logrus.Logger
}

func newTranslationRoutes(handler *gin.RouterGroup, t usecase.User, l *logrus.Logger) {
	r := &userRoutes{t, l}

	h := handler.Group("/user")
	{
		h.GET("/all", r.getAll)

	}
}

// @Summary     Get all users
// @Description Show all users
// @ID          User
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
