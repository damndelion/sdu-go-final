package http

import (
	"github.com/damndelion/sdu-go-final/config"
	"github.com/damndelion/sdu-go-final/internal/cache"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/damndelion/sdu-go-final/docs"
	"github.com/damndelion/sdu-go-final/internal/usecase"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /api
func NewRouter(handler *gin.Engine, l *logrus.Logger, u usecase.User, a usecase.Auth, cfg *config.Config, cache cache.UserCacheInterface) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/api")
	{
		newUserRoutes(h, u, l, cfg.JWT.SecretKey, cache)
		newAuthRoutes(h, a, l)
	}
}
