package service

import (
	data "Forum/Backend/database"
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

	fmt.Println(email, password)
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
	fmt.Println("We are in info login user")
	fmt.Println("this is info email ", info.emailUser)
	if err := info.db.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		fmt.Println("User not found")
		return false
	}
	fmt.Println(user.Password)
	// Vérification du mot de passe
	return info.emailUser == email && info.passwordUser == password

}
