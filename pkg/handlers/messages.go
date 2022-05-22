package handlers

import (
	"log"
	"net/http"

	"github.com/evsharonov/go-gin-gorm-crud/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) CreateMessage(context *gin.Context) {

	var newMessage models.Message

	if err := context.BindJSON(&newMessage); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := h.DB.Create(&newMessage)

	if result.Error != nil {
		log.Println(result.Error)
		context.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Ok"})
	// mocks.Messages = append(mocks.Messages, newMessage)
	// context.JSON(http.StatusCreated, mocks.Messages)
}

func (h handler) GetMessages(context *gin.Context) {

	userId := context.Param("userid")

	if userId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User ID is not valid"})
		return
	}

	var messages []models.Message

	result := h.DB.Where("user_to = ?", userId).Find(&messages)

	if result.Error != nil {
		log.Println(result.Error)
		context.JSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
		return
	}

	if len(messages) > 0 {
		context.JSON(http.StatusOK, messages)
	} else {
		context.JSON(http.StatusNotFound, gin.H{"message": "Messages not found"})
	}

}
