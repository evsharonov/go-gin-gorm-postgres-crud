package handlers

import (
	"log"
	"net/http"

	"github.com/evsharonov/go-gin-gorm-crud/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetAllUsers(context *gin.Context) {

	var users []models.User

	result := h.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
		context.JSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
		return
	}

	if len(users) != 0 {
		context.JSON(http.StatusOK, users)
	} else {
		context.JSON(http.StatusNotFound, gin.H{"message": "Users not found"})
	}

	// for _, user := range mocks.Users {
	// 	fmt.Println(user)
	// }
	// context.JSON(http.StatusOK, mocks.Users)

}

func (h handler) GetUserById(context *gin.Context) {

	var user models.User

	id := context.Param("id")

	result := h.DB.First(&user, id)

	if result.Error != nil {
		log.Println(result.Error)
		context.JSON(http.StatusNotFound, gin.H{"message": result.Error.Error()})
		return
	}

	if result.RowsAffected != 0 {
		context.JSON(http.StatusOK, user)
	} else {
		context.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	}

	// for _, user := range mocks.Users {
	// 	if fmt.Sprint(user.ID) == id {
	// 		context.JSON(http.StatusOK, user)
	// 		return
	// 	}
	// }
	//context.JSON(http.StatusNotFound, gin.H{"message": "User not found"})

}

func (h handler) CreateUser(context *gin.Context) {

	var newUser models.User

	if err := context.BindJSON(&newUser); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := h.DB.Create(&newUser)

	if result.Error != nil {
		log.Println(result.Error)
		context.JSON(http.StatusBadRequest, gin.H{"message": result.Error.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Ok"})
	//mocks.Users = append(mocks.Users, newUser)

}
