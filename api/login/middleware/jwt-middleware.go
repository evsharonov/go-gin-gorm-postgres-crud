package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/evsharonov/go-gin-gorm-crud/api/login/controller"
	"github.com/evsharonov/go-gin-gorm-crud/api/login/service"
	"github.com/gin-gonic/gin"
)

func JwtGetToken(context *gin.Context) {
	loginService := service.StaticLoginService()
	jwtService := service.JwtAuthService()
	loginController := controller.LoginHandler(loginService, jwtService)

	token := loginController.Login(context)
	if token != "" {
		context.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		context.JSON(http.StatusUnauthorized, nil)
	}
}

func JwtAuthorize(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		const BEARER_SCHEMA = "Bearer "

		authHeader := context.GetHeader("Authorization")

		if len(authHeader) == 0 {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[len(BEARER_SCHEMA):]

		token, err := service.JwtAuthService().ValidateToken(tokenString)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims)
			handler(context)
		} else {
			log.Println(err)
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
