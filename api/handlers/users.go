package handlers

import (
	"log"
	"net/http"

	"github.com/evsharonov/go-gin-gorm-crud/api/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetAllUsers(context *gin.Context) {

	user := &models.User{}

	users, err := user.GetAllUsers(h.DB)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if len(*users) != 0 {
		context.JSON(http.StatusOK, *users)
	} else {
		context.JSON(http.StatusNotFound, gin.H{"message": "Users not found"})
	}

}

func (h handler) GetUserById(context *gin.Context) {

	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User ID is not valid"})
		return
	}

	user := &models.User{}

	userReceived, err := user.GetUserById(h.DB, &id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, *userReceived)

}

func (h handler) CreateUser(context *gin.Context) {

	newUser := &models.User{}

	if err := context.BindJSON(newUser); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := newUser.CreateUser(h.DB)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Ok"})

}
