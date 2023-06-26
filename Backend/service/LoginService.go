package service

import (
	database "Forum/Backend/database"

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
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Création d'un nouvel utilisateur avec les données récupérées
	user := database.Users{

		Email:    email,
		Password: password,
	}

	return &loginInformation{
		emailUser:    user.Email,
		passwordUser: user.Password,
	}
	return nil // Retourner nil si les identifiants sont incorrects

}
func (info *loginInformation) LoginUser(email string, password string) bool {

	return info.emailUser == email && info.passwordUser == password
}
