package http

import (
	"fmt"
	"github.com/damndelion/sdu-go-final/internal/controller/http/dto"
	"github.com/damndelion/sdu-go-final/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type authRoutes struct {
	u usecase.Auth
	l *logrus.Logger
}

func newAuthRoutes(handler *gin.RouterGroup, u usecase.Auth, l *logrus.Logger) {
	r := &authRoutes{u, l}

	userHandler := handler.Group("/auth")
	{

		userHandler.POST("/register", r.Register)
		userHandler.POST("/login", r.Login)
	}
}

func (r *authRoutes) Register(ctx *gin.Context) {
	var registerRequest dto.RegisterRequest
	err := ctx.ShouldBindJSON(&registerRequest)

	if err != nil {
		r.l.Error(fmt.Errorf("http - auth - register: %w", err))
		errorResponse(ctx, http.StatusBadRequest, "Registration form is not correct")

		return
	}
	err = r.u.Register(ctx, registerRequest.Name, registerRequest.Email, registerRequest.Password)
	if err != nil {
		r.l.Error(fmt.Errorf("http  - auth - register: %w", err))
		errorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Error: %v", err))

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user successfully registered"})
}

func (r *authRoutes) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		r.l.Error(fmt.Errorf("http - auth - login: %w", err))
		errorResponse(ctx, http.StatusBadRequest, "Login form error")

		return
	}

	token, err := r.u.Login(ctx, loginRequest.Email, loginRequest.Password)
	if err != nil {
		r.l.Error(fmt.Errorf("http - auth - login: %w", err))
		errorResponse(ctx, http.StatusInternalServerError, "Login error")

		return
	}

	ctx.JSON(http.StatusOK, token)
}
