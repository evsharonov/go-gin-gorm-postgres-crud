package controller

import (
	"github.com/evsharonov/go-gin-gorm-crud/api/login/credentials"
	"github.com/evsharonov/go-gin-gorm-crud/api/login/service"
	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JwtService
}

func LoginHandler(loginService service.LoginService, jWtService service.JwtService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(context *gin.Context) string {
	credential := credentials.LoginCredentials{}
	err := context.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LogInUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Email, true)
	}
	return ""
}
