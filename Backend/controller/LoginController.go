package controller

import (
	"Forum/Backend/LoCred"
	service "Forum/Backend/service"

	"github.com/gin-gonic/gin"
)

// login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService,
) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential LoCred.LoginCredentials
	credential.Email = ctx.PostForm("Email")
	credential.Password = ctx.PostForm("Password")
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Email, true)
	}
	return ""
}
