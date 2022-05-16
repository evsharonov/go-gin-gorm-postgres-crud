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
		log.Fatalln(err)
		return
	}

	result := h.DB.Create(&newMessage)

	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Ok"})
	// mocks.Messages = append(mocks.Messages, newMessage)
	// context.IndentedJSON(http.StatusCreated, mocks.Messages)
}

func (h handler) GetMessages(context *gin.Context) {

	userId := context.Param("userid")

	rows, err := h.DB.Model(&models.Message{}).Where("user_to = ?", userId).Rows()

	if err != nil {
		log.Fatalln(err)
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}

	defer rows.Close()

	var messages []models.Message
	var message models.Message

	for rows.Next() {
		h.DB.ScanRows(rows, &message)
		messages = append(messages, message)
	}

	// for _, message := range mocks.Messages {
	// 	if fmt.Sprint(message.UserTo) == userId {
	// 		messages = append(messages, message)
	// 	}
	// }

	if len(messages) > 0 {
		context.IndentedJSON(http.StatusOK, messages)
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Messages not found"})
	}

}
