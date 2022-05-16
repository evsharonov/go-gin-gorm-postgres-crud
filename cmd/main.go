package main

import (
	"github.com/evsharonov/go-gin-gorm-crud/pkg/db"
	"github.com/evsharonov/go-gin-gorm-crud/pkg/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)

	r := gin.Default()
	r.GET("/users", h.GetAllUsers)
	r.GET("/users/:id", h.GetUserById)
	r.POST("/users/create", h.CreateUser)
	r.GET("/messages/:userid", h.GetMessages)
	r.POST("/messages/create", h.CreateMessage)

	r.Run("localhost:8080")
}
