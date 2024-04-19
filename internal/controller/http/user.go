package http

import (
	"encoding/json"
	"fmt"
	"github.com/damndelion/sdu-go-final/internal/cache"
	"github.com/damndelion/sdu-go-final/internal/controller/middleware"
	"github.com/damndelion/sdu-go-final/internal/entity"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/damndelion/sdu-go-final/internal/usecase"
)

type userRoutes struct {
	t     usecase.User
	l     *logrus.Logger
	cache cache.UserCacheInterface
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.User, l *logrus.Logger, key string, c cache.UserCacheInterface) {
	r := &userRoutes{t, l, c}

	h := handler.Group("/user")
	{
		h.Use(middleware.JwtVerify(key))
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
func (r *userRoutes) getAll(ctx *gin.Context) {
	userID, _ := ctx.Get("user_id")
	cacheUsers, err := r.cache.Get(ctx, userID.(string))
	fmt.Println(cacheUsers)
	if err != nil {
		r.l.Error(err, "http - v1 - user")
		errorResponse(ctx, http.StatusInternalServerError, "redis problems")
		return
	}

	if cacheUsers == "" {
		time.Sleep(1 * time.Second)

		users, err := r.t.GetUsers(ctx.Request.Context())
		if err != nil {
			r.l.Error(err, "http - v1 - user")
			errorResponse(ctx, http.StatusInternalServerError, "database problems")

			return
		}
		jsonData, err := json.Marshal(users)
		err = r.cache.Set(ctx, userID.(string), string(jsonData))
		if err != nil {
			r.l.Error(fmt.Errorf("http - v1 - user - getUsers: %w", err))
			errorResponse(ctx, http.StatusInternalServerError, "getUsersById cache error")
		}
		ctx.JSON(http.StatusOK, users)
		return
	}
	tempUsers := &[]entity.User{}
	err = json.Unmarshal([]byte(cacheUsers), tempUsers)
	if err != nil {
		r.l.Error(err, "http - v1 - users")
		errorResponse(ctx, http.StatusInternalServerError, "redis problems")
		return
	}

	ctx.JSON(http.StatusOK, tempUsers)
}
