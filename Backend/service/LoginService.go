package service

import (
	data "Forum/Backend/Database"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginService interface {
	LoginUser(email string, password string) bool
}

type loginInformation struct {
	emailUser    string
	passwordUser string
	db           *gorm.DB
}

func StaticLoginService(c *gin.Context, db *gorm.DB) LoginService {
	email := c.PostForm("Email")
	password := c.PostForm("Password")

	return &loginInformation{
		emailUser:    email,
		passwordUser: password,
		db:           db,
	}
}

func (info *loginInformation) LoginUser(email string, password string) bool {
	user := data.Users{
		Email:    info.emailUser,
		Password: info.passwordUser,
	}

	if err := info.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		fmt.Println("User not found")
		return false
	}

	// Vérification du mot de passe
	return info.emailUser == email && info.passwordUser == password

}
