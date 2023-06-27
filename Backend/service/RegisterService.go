package service

import (
	database "Forum/Backend/Database"

	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUser(c *gin.Context, db *gorm.DB) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Création d'un nouvel utilisateur avec les données récupérées
	user := database.Users{
		Username: username,
		Email:    email,
		Password: password,
	}
	fmt.Print("_________________________REGISTER_INFO_________________________")
	fmt.Println(user.Username)
	fmt.Println(user.Email)
	fmt.Println(user.Password)
	// Appel à la fonction createDB() pour créer et initialiser l'utilisateur
	//Database.CreateDB(db)
	database.AddUser(user.Username, user.Email, user.Password, db)

	// Fermeture de la connexion à la base de données

	// Réponse au client

}
