package main

import (
	"github.com/evsharonov/go-gin-gorm-crud/api/db"
	"github.com/evsharonov/go-gin-gorm-crud/api/handlers"
	"github.com/evsharonov/go-gin-gorm-crud/api/login/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	dbase := db.Init()
	h := handlers.New(dbase)

	r := gin.Default()

	r.POST("/login", middleware.JwtGetToken)

	r.GET("/users", middleware.JwtAuthorize(h.GetAllUsers))
	r.GET("/users/:id", middleware.JwtAuthorize(h.GetUserById))
	r.POST("/users/create", middleware.JwtAuthorize(h.CreateUser))
	r.GET("/messages/:userid", middleware.JwtAuthorize(h.GetMessages))
	r.POST("/messages/create", middleware.JwtAuthorize(h.CreateMessage))

	r.Run("localhost:8080")
}
