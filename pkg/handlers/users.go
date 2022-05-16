package handlers

import (
	"log"
	"net/http"

	"github.com/evsharonov/go-gin-gorm-crud/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetAllUsers(context *gin.Context) {

	rows, err := h.DB.Model(&models.User{}).Rows()

	if err != nil {
		log.Fatalln(err)
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}

	defer rows.Close()

	var users []models.User
	var user models.User

	for rows.Next() {
		h.DB.ScanRows(rows, &user)
		users = append(users, user)
	}

	if len(users) != 0 {
		context.IndentedJSON(http.StatusOK, users)
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Users not found"})
	}

	// for _, user := range mocks.Users {
	// 	fmt.Println(user)
	// }
	// context.IndentedJSON(http.StatusOK, mocks.Users)

}

func (h handler) GetUserById(context *gin.Context) {

	var user models.User

	id := context.Param("id")

	result := h.DB.First(&user, id)

	if result.RowsAffected != 0 {
		context.IndentedJSON(http.StatusOK, user)
	} else {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": result.Error})
	}

	// for _, user := range mocks.Users {
	// 	if fmt.Sprint(user.ID) == id {
	// 		context.IndentedJSON(http.StatusOK, user)
	// 		return
	// 	}
	// }
	//context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})

}

func (h handler) CreateUser(context *gin.Context) {

	var newUser models.User

	if err := context.BindJSON(&newUser); err != nil {
		log.Fatalln(err)
	}

	result := h.DB.Create(&newUser)

	if result.Error != nil {
		log.Fatalln(result.Error)
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Ok"})
	//mocks.Users = append(mocks.Users, newUser)

}
