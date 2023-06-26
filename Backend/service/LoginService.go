package service

import (
	database "Forum/Backend/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

type LoginService interface {
	LoginUser(email string, password string) bool
}
type loginInformation struct {
	emailUser    string
	passwordUser string
}

func StaticLoginService(c *gin.Context) LoginService {
	email := c.PostForm("Email")
	password := c.PostForm("Password")

	fmt.Println(email)
	// Création d'un nouvel utilisateur avec les données récupérées
	user := database.Users{

		Email:    email,
		Password: password,
	}
	fmt.Println(user)
	return &loginInformation{
		emailUser:    email,
		passwordUser: password,
	}

}
func (info *loginInformation) LoginUser(email string, password string) bool {

	fmt.Println(info.emailUser == email && info.passwordUser == password)
	return info.emailUser == email && info.passwordUser == password
}
