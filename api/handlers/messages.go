package handlers

import (
	"log"
	"net/http"

	"github.com/evsharonov/go-gin-gorm-crud/api/models"
	"github.com/gin-gonic/gin"
)

func (h handler) CreateMessage(context *gin.Context) {

	newMessage := &models.Message{}

	if err := context.BindJSON(newMessage); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := newMessage.CreateMessage(h.DB)

	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Ok"})

}

func (h handler) GetMessages(context *gin.Context) {

	userId := context.Param("userid")

	if userId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User ID is not valid"})
		return
	}

	messages := &models.Message{}

	messagesReceived, err := messages.GetMessages(h.DB, &userId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if len(*messagesReceived) > 0 {
		context.JSON(http.StatusOK, *messagesReceived)
	} else {
		context.JSON(http.StatusNotFound, gin.H{"message": "Messages not found"})
	}
}
