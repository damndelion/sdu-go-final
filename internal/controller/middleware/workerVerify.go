package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func WorkerVerify(SecretKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header.Get("Authorization")

		if tokenHeader == "" || !strings.HasPrefix(tokenHeader, "Bearer ") {
			ctx.AbortWithStatus(http.StatusForbidden)

			return
		}

		tokenString := strings.TrimPrefix(tokenHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(SecretKey), nil
		})
		var claims jwt.MapClaims
		var ok bool
		if claims, ok = token.Claims.(jwt.MapClaims); ok && token.Valid {
			if exp, ok := claims["exp"].(float64); ok {
				expirationTime := time.Unix(int64(exp), 0)
				if time.Now().After(expirationTime) {
					ctx.AbortWithStatus(http.StatusUnauthorized)

					return
				}
			} else {
				ctx.AbortWithStatus(http.StatusUnauthorized)

				return
			}
		} else {
			ctx.AbortWithStatus(http.StatusUnauthorized)

			return
		}
		if err != nil || !token.Valid {
			ctx.AbortWithStatus(http.StatusForbidden)

			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})

			return
		}

		if role != "worker" || role != "admin" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})

			return
		}

		ctx.Set("user_id", claims["user_id"])

		ctx.Next()
	}
}
